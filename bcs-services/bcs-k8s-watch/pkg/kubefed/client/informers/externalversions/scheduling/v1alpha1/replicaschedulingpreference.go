/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	scheduling_v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/scheduling/v1alpha1"
	versioned "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/listers/scheduling/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReplicaSchedulingPreferenceInformer provides access to a shared informer and lister for
// ReplicaSchedulingPreferences.
type ReplicaSchedulingPreferenceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReplicaSchedulingPreferenceLister
}

type replicaSchedulingPreferenceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReplicaSchedulingPreferenceInformer constructs a new informer for ReplicaSchedulingPreference type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReplicaSchedulingPreferenceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReplicaSchedulingPreferenceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReplicaSchedulingPreferenceInformer constructs a new informer for ReplicaSchedulingPreference type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReplicaSchedulingPreferenceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1alpha1().ReplicaSchedulingPreferences(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1alpha1().ReplicaSchedulingPreferences(namespace).Watch(options)
			},
		},
		&scheduling_v1alpha1.ReplicaSchedulingPreference{},
		resyncPeriod,
		indexers,
	)
}

func (f *replicaSchedulingPreferenceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReplicaSchedulingPreferenceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *replicaSchedulingPreferenceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&scheduling_v1alpha1.ReplicaSchedulingPreference{}, f.defaultInformer)
}

func (f *replicaSchedulingPreferenceInformer) Lister() v1alpha1.ReplicaSchedulingPreferenceLister {
	return v1alpha1.NewReplicaSchedulingPreferenceLister(f.Informer().GetIndexer())
}
