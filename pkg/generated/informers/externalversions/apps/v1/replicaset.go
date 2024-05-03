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

// Code generated by wranglertest. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	versioned "github.com/tomleb/wrangler_test/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/tomleb/wrangler_test/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/tomleb/wrangler_test/pkg/generated/listers/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReplicaSetInformer provides access to a shared informer and lister for
// ReplicaSets.
type ReplicaSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ReplicaSetLister
}

type replicaSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReplicaSetInformer constructs a new informer for ReplicaSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReplicaSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReplicaSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReplicaSetInformer constructs a new informer for ReplicaSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReplicaSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().ReplicaSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().ReplicaSets(namespace).Watch(context.TODO(), options)
			},
		},
		&appsv1.ReplicaSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *replicaSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReplicaSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *replicaSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1.ReplicaSet{}, f.defaultInformer)
}

func (f *replicaSetInformer) Lister() v1.ReplicaSetLister {
	return v1.NewReplicaSetLister(f.Informer().GetIndexer())
}
