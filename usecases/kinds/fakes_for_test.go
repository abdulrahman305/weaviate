//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package kinds

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/entities/search"
	"github.com/semi-technologies/weaviate/usecases/network/common/peers"
	"github.com/stretchr/testify/mock"
)

type fakeRepo struct {
	mock.Mock
	GetThingResponse     *models.Thing
	GetActionResponse    *models.Action
	UpdateThingParameter *models.Thing
}

func (f *fakeRepo) ClassExists(ctx context.Context, id strfmt.UUID) (bool, error) {
	args := f.Called(id)
	return args.Bool(0), args.Error(1)
}

func (f *fakeRepo) AddAction(ctx context.Context, class *models.Action, id strfmt.UUID) error {
	args := f.Called(class, id)
	return args.Error(0)
}

func (f *fakeRepo) AddThing(ctx context.Context, class *models.Thing, id strfmt.UUID) error {
	args := f.Called(class, id)
	return args.Error(0)
}

func (f *fakeRepo) GetThing(ctx context.Context, id strfmt.UUID, thing *models.Thing) error {
	*thing = *f.GetThingResponse
	return nil
}

func (f *fakeRepo) GetAction(ctx context.Context, id strfmt.UUID, action *models.Action) error {
	*action = *f.GetActionResponse
	return nil
}

func (f *fakeRepo) ListThings(ctx context.Context, limit int, thingsResponse *models.ThingsListResponse) error {
	panic("not implemented")
}

func (f *fakeRepo) ListActions(ctx context.Context, limit int, actionsResponse *models.ActionsListResponse) error {
	panic("not implemented")
}

func (f *fakeRepo) UpdateAction(ctx context.Context, class *models.Action, id strfmt.UUID) error {
	panic("not implemented")
}

func (f *fakeRepo) UpdateThing(ctx context.Context, class *models.Thing, id strfmt.UUID) error {
	f.UpdateThingParameter = class
	return nil
}

func (f *fakeRepo) DeleteThing(ctx context.Context, thing *models.Thing, uuid strfmt.UUID) error {
	args := f.Called(thing, uuid)
	return args.Error(0)
}

func (f *fakeRepo) DeleteAction(ctx context.Context, thing *models.Action, uuid strfmt.UUID) error {
	args := f.Called(thing, uuid)
	return args.Error(0)
}

func (f *fakeRepo) AddThingsBatch(ctx context.Context, things BatchThings) error {
	args := f.Called(things)
	return args.Error(0)
}

func (f *fakeRepo) AddActionsBatch(ctx context.Context, actions BatchActions) error {
	args := f.Called(actions)
	return args.Error(0)
}

func (f *fakeRepo) AddBatchReferences(ctx context.Context, references BatchReferences) error {
	panic("not implemented")
}

type fakeSchemaManager struct {
	CalledWith struct {
		kind      kind.Kind
		fromClass string
		property  string
		toClass   string
	}
	GetSchemaResponse schema.Schema
}

func (f *fakeSchemaManager) UpdatePropertyAddDataType(ctx context.Context, principal *models.Principal,
	k kind.Kind, fromClass, property, toClass string) error {
	f.CalledWith = struct {
		kind      kind.Kind
		fromClass string
		property  string
		toClass   string
	}{
		kind:      k,
		fromClass: fromClass,
		property:  property,
		toClass:   toClass,
	}
	return nil
}

func (f *fakeSchemaManager) GetSchema(principal *models.Principal) (schema.Schema, error) {
	return f.GetSchemaResponse, nil
}

type fakeLocks struct{}

func (f *fakeLocks) LockConnector() (func() error, error) {
	return func() error { return nil }, nil
}

func (f *fakeLocks) LockSchema() (func() error, error) {
	return func() error { return nil }, nil
}

type fakeVectorizer struct{}

func (f *fakeVectorizer) Thing(ctx context.Context, thing *models.Thing) ([]float32, error) {
	return []float32{0, 1, 2}, nil
}

func (f *fakeVectorizer) Action(ctx context.Context, thing *models.Action) ([]float32, error) {
	return []float32{0, 1, 2}, nil
}

func (f *fakeVectorizer) Corpi(ctx context.Context, corpi []string) ([]float32, error) {
	panic("not implemented")
}

type fakeNetwork struct {
	peerURI string
}

func (f *fakeNetwork) ListPeers() (peers.Peers, error) {
	myPeers := peers.Peers{
		peers.Peer{
			Name: "BestWeaviate",
			URI:  strfmt.URI(f.peerURI),
			Schema: schema.Schema{
				Things: &models.Schema{
					Classes: []*models.Class{
						&models.Class{
							Class: "BestThing",
						},
					},
				},
			},
		},
	}

	return myPeers, nil
}

type fakeAuthorizer struct{}

func (f *fakeAuthorizer) Authorize(principal *models.Principal, verb, resource string) error {
	return nil
}

type fakeC11y struct{}

func (f *fakeC11y) IsWordPresent(ctx context.Context, word string) (bool, error) {
	panic("not implemented")
}
func (f *fakeC11y) SafeGetSimilarWordsWithCertainty(ctx context.Context, word string, certainty float32) ([]string, error) {
	panic("not implemented")
}

type fakeVectorRepo struct {
	mock.Mock
}

func (f *fakeVectorRepo) ThingByID(ctx context.Context,
	id strfmt.UUID) (*search.Result, error) {
	return nil, nil
}

func (f *fakeVectorRepo) ActionByID(ctx context.Context,
	id strfmt.UUID) (*search.Result, error) {
	return nil, nil
}

func (f *fakeVectorRepo) ThingSearch(ctx context.Context, limit int,
	filters *filters.LocalFilter) (search.Results, error) {
	return nil, nil
}

func (f *fakeVectorRepo) ActionSearch(ctx context.Context, limit int,
	filters *filters.LocalFilter) (search.Results, error) {
	return nil, nil
}

func (f *fakeVectorRepo) PutThing(ctx context.Context,
	concept *models.Thing, vector []float32) error {
	return nil
}
func (f *fakeVectorRepo) PutAction(ctx context.Context,
	concept *models.Action, vector []float32) error {
	return nil
}

func (r *fakeVectorRepo) BatchPutThings(ctx context.Context, batch BatchThings) (BatchThings, error) {
	return nil, nil
}

func (r *fakeVectorRepo) BatchPutActions(ctx context.Context, batch BatchActions) (BatchActions, error) {
	return nil, nil
}

func (f *fakeVectorRepo) DeleteAction(ctx context.Context,
	className string, id strfmt.UUID) error {
	args := f.Called(className, id)
	return args.Error(0)
}

func (f *fakeVectorRepo) DeleteThing(ctx context.Context,
	className string, id strfmt.UUID) error {
	args := f.Called(className, id)
	return args.Error(0)
}