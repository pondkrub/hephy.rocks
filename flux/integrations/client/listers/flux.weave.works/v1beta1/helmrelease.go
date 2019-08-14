/*
Copyright 2018 Weaveworks Ltd.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/weaveworks/flux/integrations/apis/flux.weave.works/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HelmReleaseLister helps list HelmReleases.
type HelmReleaseLister interface {
	// List lists all HelmReleases in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.HelmRelease, err error)
	// HelmReleases returns an object that can list and get HelmReleases.
	HelmReleases(namespace string) HelmReleaseNamespaceLister
	HelmReleaseListerExpansion
}

// helmReleaseLister implements the HelmReleaseLister interface.
type helmReleaseLister struct {
	indexer cache.Indexer
}

// NewHelmReleaseLister returns a new HelmReleaseLister.
func NewHelmReleaseLister(indexer cache.Indexer) HelmReleaseLister {
	return &helmReleaseLister{indexer: indexer}
}

// List lists all HelmReleases in the indexer.
func (s *helmReleaseLister) List(selector labels.Selector) (ret []*v1beta1.HelmRelease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.HelmRelease))
	})
	return ret, err
}

// HelmReleases returns an object that can list and get HelmReleases.
func (s *helmReleaseLister) HelmReleases(namespace string) HelmReleaseNamespaceLister {
	return helmReleaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HelmReleaseNamespaceLister helps list and get HelmReleases.
type HelmReleaseNamespaceLister interface {
	// List lists all HelmReleases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.HelmRelease, err error)
	// Get retrieves the HelmRelease from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.HelmRelease, error)
	HelmReleaseNamespaceListerExpansion
}

// helmReleaseNamespaceLister implements the HelmReleaseNamespaceLister
// interface.
type helmReleaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HelmReleases in the indexer for a given namespace.
func (s helmReleaseNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.HelmRelease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.HelmRelease))
	})
	return ret, err
}

// Get retrieves the HelmRelease from the indexer for a given namespace and name.
func (s helmReleaseNamespaceLister) Get(name string) (*v1beta1.HelmRelease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("helmrelease"), name)
	}
	return obj.(*v1beta1.HelmRelease), nil
}
