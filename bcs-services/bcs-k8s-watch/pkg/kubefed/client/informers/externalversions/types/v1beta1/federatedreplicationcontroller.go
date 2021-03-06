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

package v1beta1

import (
	types_v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	versioned "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/listers/types/v1beta1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FederatedReplicationControllerInformer provides access to a shared informer and lister for
// FederatedReplicationControllers.
type FederatedReplicationControllerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FederatedReplicationControllerLister
}

type federatedReplicationControllerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedReplicationControllerInformer constructs a new informer for FederatedReplicationController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedReplicationControllerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedReplicationControllerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedReplicationControllerInformer constructs a new informer for FederatedReplicationController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedReplicationControllerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedReplicationControllers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedReplicationControllers(namespace).Watch(options)
			},
		},
		&types_v1beta1.FederatedReplicationController{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedReplicationControllerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedReplicationControllerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedReplicationControllerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&types_v1beta1.FederatedReplicationController{}, f.defaultInformer)
}

func (f *federatedReplicationControllerInformer) Lister() v1beta1.FederatedReplicationControllerLister {
	return v1beta1.NewFederatedReplicationControllerLister(f.Informer().GetIndexer())
}
