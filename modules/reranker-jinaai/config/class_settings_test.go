//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weaviate/weaviate/entities/moduletools"
	"github.com/weaviate/weaviate/entities/schema"
)

func Test_classSettings_Validate(t *testing.T) {
	tests := []struct {
		name      string
		cfg       moduletools.ClassConfig
		wantModel string
		wantErr   error
	}{
		{
			name: "default settings",
			cfg: fakeClassConfig{
				classConfig: map[string]interface{}{},
			},
			wantModel: "jina-reranker-v2-base-multilingual",
		},
		{
			name: "custom settings",
			cfg: fakeClassConfig{
				classConfig: map[string]interface{}{
					"model": "jina-reranker-v1-base-en",
				},
			},
			wantModel: "jina-reranker-v1-base-en",
		},
		{
			name: "unsupported model error",
			cfg: fakeClassConfig{
				classConfig: map[string]interface{}{
					"model": "jina-reranker-v1-base-pl",
				},
			},
			wantErr: fmt.Errorf("wrong Jinaai model name, available model names are: [jina-reranker-v2-base-multilingual jina-reranker-v1-base-en jina-reranker-v1-turbo-en jina-reranker-v1-tiny-en jina-colbert-v1-en]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := NewClassSettings(tt.cfg)
			if tt.wantErr != nil {
				assert.EqualError(t, ic.Validate(nil), tt.wantErr.Error())
			} else {
				assert.Equal(t, tt.wantModel, ic.Model())
			}
		})
	}
}

type fakeClassConfig struct {
	classConfig map[string]interface{}
}

func (f fakeClassConfig) Class() map[string]interface{} {
	return f.classConfig
}

func (f fakeClassConfig) Tenant() string {
	return ""
}

func (f fakeClassConfig) ClassByModuleName(moduleName string) map[string]interface{} {
	return f.classConfig
}

func (f fakeClassConfig) Property(propName string) map[string]interface{} {
	return nil
}

func (f fakeClassConfig) TargetVector() string {
	return ""
}

func (f fakeClassConfig) PropertiesDataTypes() map[string]schema.DataType {
	return nil
}
