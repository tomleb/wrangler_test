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
	"context"
	time "time"

	clustercattleiov3 "github.com/tomleb/wrangler_test/pkg/apis/cluster.cattle.io/v3"
	versioned "github.com/tomleb/wrangler_test/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/tomleb/wrangler_test/pkg/generated/informers/externalversions/internalinterfaces"
	v3 "github.com/tomleb/wrangler_test/pkg/generated/listers/cluster.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterUserAttributeInformer provides access to a shared informer and lister for
// ClusterUserAttributes.
type ClusterUserAttributeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v3.ClusterUserAttributeLister
}

type clusterUserAttributeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterUserAttributeInformer constructs a new informer for ClusterUserAttribute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterUserAttributeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterUserAttributeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterUserAttributeInformer constructs a new informer for ClusterUserAttribute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterUserAttributeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV3().ClusterUserAttributes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV3().ClusterUserAttributes(namespace).Watch(context.TODO(), options)
			},
		},
		&clustercattleiov3.ClusterUserAttribute{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterUserAttributeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterUserAttributeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterUserAttributeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clustercattleiov3.ClusterUserAttribute{}, f.defaultInformer)
}

func (f *clusterUserAttributeInformer) Lister() v3.ClusterUserAttributeLister {
	return v3.NewClusterUserAttributeLister(f.Informer().GetIndexer())
}
