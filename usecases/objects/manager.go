package objects

import (
	"context"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/entities/additional"
	"github.com/weaviate/weaviate/entities/filters"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/modulecapabilities"
	"github.com/weaviate/weaviate/entities/schema/crossref"
	"github.com/weaviate/weaviate/entities/search"
	"github.com/weaviate/weaviate/usecases/config"
)

type objectsMetrics interface {
	BatchInc()
	BatchDec()
	BatchRefInc()
	BatchRefDec()
	BatchDeleteInc()
	BatchDeleteDec()
	AddObjectInc()
	AddObjectDec()
	UpdateObjectInc()
	UpdateObjectDec()
	MergeObjectInc()
	MergeObjectDec()
	DeleteObjectInc()
	DeleteObjectDec()
	GetObjectInc()
	GetObjectDec()
	HeadObjectInc()
	HeadObjectDec()
	AddReferenceInc()
	AddReferenceDec()
	UpdateReferenceInc()
	UpdateReferenceDec()
	DeleteReferenceInc()
	DeleteReferenceDec()
	AddUsageDimensions(className, queryType, operation string, dims int)
}

type timeSource interface {
	Now() int64
}

type locks interface {
	LockConnector() (func() error, error)
	LockSchema() (func() error, error)
}

type authorizer interface {
	Authorize(principal *models.Principal, verb, resource string) error
}

type VectorRepo interface {
	PutObject(ctx context.Context, concept *models.Object, vector []float32, vectors models.Vectors,
		repl *additional.ReplicationProperties) error
	DeleteObject(ctx context.Context, className string, id strfmt.UUID,
		repl *additional.ReplicationProperties, tenant string) error
	Object(ctx context.Context, class string, id strfmt.UUID, props search.SelectProperties,
		additional additional.Properties, repl *additional.ReplicationProperties,
		tenant string) (*search.Result, error)
	Exists(ctx context.Context, class string, id strfmt.UUID,
		repl *additional.ReplicationProperties, tenant string) (bool, error)
	ObjectByID(ctx context.Context, id strfmt.UUID, props search.SelectProperties,
		additional additional.Properties, tenant string) (*search.Result, error)
	ObjectSearch(ctx context.Context, offset, limit int, filters *filters.LocalFilter,
		sort []filters.Sort, additional additional.Properties, tenant string) (search.Results, error)
	AddReference(ctx context.Context, source *crossref.RefSource,
		target *crossref.Ref, repl *additional.ReplicationProperties, tenant string) error
	Merge(ctx context.Context, merge MergeDocument, repl *additional.ReplicationProperties, tenant string) error
	Query(context.Context, *QueryInput) (search.Results, *Error)
}

type ModulesProvider interface {
	GetObjectAdditionalExtend(ctx context.Context, in *search.Result,
		moduleParams map[string]interface{}) (*search.Result, error)
	ListObjectsAdditionalExtend(ctx context.Context, in search.Results,
		moduleParams map[string]interface{}) (search.Results, error)
	UsingRef2Vec(className string) bool
	UpdateVector(ctx context.Context, object *models.Object, class *models.Class, repo modulecapabilities.FindObjectFn,
		logger logrus.FieldLogger) error
	VectorizerName(className string) (string, error)
}

type Manager struct {
	config            *config.WeaviateConfig
	locks             locks
	schemaManager     schemaManager
	logger            logrus.FieldLogger
	authorizer        authorizer
	vectorRepo        VectorRepo
	timeSource        timeSource
	modulesProvider   ModulesProvider
	autoSchemaManager *autoSchemaManager
	metrics           objectsMetrics
}

func NewManager(locks locks, schemaManager schemaManager,
	config *config.WeaviateConfig, logger logrus.FieldLogger,
	authorizer authorizer, vectorRepo VectorRepo,
	modulesProvider ModulesProvider, metrics objectsMetrics,
) *Manager {
	return &Manager{
		config:            config,
		locks:             locks,
		schemaManager:     schemaManager,
		logger:            logger,
		authorizer:        authorizer,
		vectorRepo:        vectorRepo,
		timeSource:        defaultTimeSource{},
		modulesProvider:   modulesProvider,
		autoSchemaManager: newAutoSchemaManager(schemaManager, vectorRepo, config, logger),
		metrics:           metrics,
	}
}

func generateUUID() (strfmt.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("could not generate uuid v4: %v", err)
	}

	return strfmt.UUID(id.String()), nil
}

type defaultTimeSource struct{}

func (ts defaultTimeSource) Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (m *Manager) UpdateObject(ctx context.Context, principal *models.Principal,
	class string, id strfmt.UUID, updates *models.Object,
	repl *additional.ReplicationProperties,
) (*models.Object, error) {
	path := fmt.Sprintf("objects/%s/%s", class, id)
	if class == "" {
		path = fmt.Sprintf("objects/%s", id)
	}
	err := m.authorizer.Authorize(principal, "update", path)
	if err != nil {
		return nil, err
	}

	m.metrics.UpdateObjectInc()
	defer m.metrics.UpdateObjectDec()

	unlock, err := m.locks.LockSchema()
	if err != nil {
		return nil, NewErrInternal("could not acquire lock: %v", err)
	}
	defer unlock()

	return m.updateObjectToConnectorAndSchema(ctx, principal, class, id, updates, repl)
}

func (m *Manager) updateObjectToConnectorAndSchema(ctx context.Context,
	principal *models.Principal, className string, id strfmt.UUID, updates *models.Object,
	repl *additional.ReplicationProperties,
) (*models.Object, error) {
	if id != updates.ID {
		return nil, NewErrInvalidUserInput("invalid update: field 'id' is immutable")
	}

	obj, err := m.getObjectFromRepo(ctx, className, id, additional.Properties{}, repl, updates.Tenant)
	if err != nil {
		return nil, err
	}

	err = m.autoSchemaManager.autoSchema(ctx, principal, updates, false)
	if err != nil {
		return nil, NewErrInvalidUserInput("invalid object: %v", err)
	}

	m.logger.
		WithField("object", "kinds_update_requested").
		WithField("original", obj).
		WithField("updated", updates).
		WithField("id", id).
		Debug("received update kind request")

	prevObj := obj.Object()
	err = m.validateObjectAndNormalizeNames(
		ctx, principal, repl, updates, prevObj)
	if err != nil {
		return nil, NewErrInvalidUserInput("invalid object: %v", err)
	}

	updates.CreationTimeUnix = obj.Created
	updates.LastUpdateTimeUnix = m.timeSource.Now()

	class, err := m.schemaManager.GetClass(ctx, principal, className)
	if err != nil {
		return nil, err
	}

	err = m.modulesProvider.UpdateVector(ctx, updates, class, m.findObject, m.logger)
	if err != nil {
		return nil, NewErrInternal("update object: %v", err)
	}

	err = m.vectorRepo.PutObject(ctx, updates, updates.Vector, updates.Vectors, repl)
	if err != nil {
		return nil, fmt.Errorf("put object: %w", err)
	}

	return updates, nil
}
