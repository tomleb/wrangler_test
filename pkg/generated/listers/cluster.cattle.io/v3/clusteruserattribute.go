/*
Copyright The Kubernetes Authors.

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

// Code generated by wrangler_test. DO NOT EDIT.

package v3

import (
	v3 "github.com/tomleb/wrangler_test/pkg/apis/cluster.cattle.io/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterUserAttributeLister helps list ClusterUserAttributes.
// All objects returned here must be treated as read-only.
type ClusterUserAttributeLister interface {
	// List lists all ClusterUserAttributes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.ClusterUserAttribute, err error)
	// ClusterUserAttributes returns an object that can list and get ClusterUserAttributes.
	ClusterUserAttributes(namespace string) ClusterUserAttributeNamespaceLister
	ClusterUserAttributeListerExpansion
}

// clusterUserAttributeLister implements the ClusterUserAttributeLister interface.
type clusterUserAttributeLister struct {
	indexer cache.Indexer
}

// NewClusterUserAttributeLister returns a new ClusterUserAttributeLister.
func NewClusterUserAttributeLister(indexer cache.Indexer) ClusterUserAttributeLister {
	return &clusterUserAttributeLister{indexer: indexer}
}

// List lists all ClusterUserAttributes in the indexer.
func (s *clusterUserAttributeLister) List(selector labels.Selector) (ret []*v3.ClusterUserAttribute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ClusterUserAttribute))
	})
	return ret, err
}

// ClusterUserAttributes returns an object that can list and get ClusterUserAttributes.
func (s *clusterUserAttributeLister) ClusterUserAttributes(namespace string) ClusterUserAttributeNamespaceLister {
	return clusterUserAttributeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterUserAttributeNamespaceLister helps list and get ClusterUserAttributes.
// All objects returned here must be treated as read-only.
type ClusterUserAttributeNamespaceLister interface {
	// List lists all ClusterUserAttributes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.ClusterUserAttribute, err error)
	// Get retrieves the ClusterUserAttribute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.ClusterUserAttribute, error)
	ClusterUserAttributeNamespaceListerExpansion
}

// clusterUserAttributeNamespaceLister implements the ClusterUserAttributeNamespaceLister
// interface.
type clusterUserAttributeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterUserAttributes in the indexer for a given namespace.
func (s clusterUserAttributeNamespaceLister) List(selector labels.Selector) (ret []*v3.ClusterUserAttribute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ClusterUserAttribute))
	})
	return ret, err
}

// Get retrieves the ClusterUserAttribute from the indexer for a given namespace and name.
func (s clusterUserAttributeNamespaceLister) Get(name string) (*v3.ClusterUserAttribute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("clusteruserattribute"), name)
	}
	return obj.(*v3.ClusterUserAttribute), nil
}
