/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocks

import (
	"net/http"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"k8s.io/client-go/kubernetes/scheme"
)

// Mock Types for Reconciler tests:
type Manager struct {
	client client.Client
	cache  cache.Cache
	config rest.Config
	Scheme *runtime.Scheme
}

var _ manager.Manager = &Manager{}

func NewManager(c FakeClient) Manager {
	return Manager{
		client: c,
		cache:  FakeCache{},
	}
}

func (Manager) Add(manager.Runnable) error {
	panic("implement me")
}

func (Manager) SetFields(interface{}) error {
	panic("implement me")
}

func (Manager) Start(<-chan struct{}) error {
	panic("implement me")
}

func (m Manager) GetConfig() *rest.Config {
	return &m.config
}

func (m Manager) GetScheme() *runtime.Scheme {
	return m.Scheme
}

func (m Manager) GetClient() client.Client {
	if m.client == nil {
		m.client = NewClient(scheme.Scheme)
	}
	return m.client
}

func (Manager) GetFieldIndexer() client.FieldIndexer {
	panic("implement me")
}

func (m Manager) GetCache() cache.Cache {
	if m.cache == nil {
		m.cache = FakeCache{}
	}
	return m.cache
}

func (Manager) GetRecorder(name string) record.EventRecorder {
	panic("implement me")
}

func (Manager) GetRESTMapper() meta.RESTMapper {
	panic("implement me")
}

func (Manager) GetAPIReader() client.Reader {
	panic("implement me")
}

func (Manager) GetEventRecorderFor(name string) record.EventRecorder {
	panic("implement me")
}

func (Manager) GetWebhookServer() *webhook.Server {
	panic("implement me")
}

func (Manager) AddHealthzCheck(name string, check healthz.Checker) error {
	panic("implement me")
}

func (Manager) AddReadyzCheck(name string, check healthz.Checker) error {
	panic("implement me")
}

func (Manager) AddMetricsExtraHandler(path string, handler http.Handler) error {
	panic("implement me")
}

func (Manager) Elected() <-chan struct{} {
	panic("implement me")
}
