/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-general-pod-autoscaler/pkg/apis/autoscaling/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GeneralPodAutoscalerLister helps list GeneralPodAutoscalers.
type GeneralPodAutoscalerLister interface {
	// List lists all GeneralPodAutoscalers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GeneralPodAutoscaler, err error)
	// GeneralPodAutoscalers returns an object that can list and get GeneralPodAutoscalers.
	GeneralPodAutoscalers(namespace string) GeneralPodAutoscalerNamespaceLister
	GeneralPodAutoscalerListerExpansion
}

// generalPodAutoscalerLister implements the GeneralPodAutoscalerLister interface.
type generalPodAutoscalerLister struct {
	indexer cache.Indexer
}

// NewGeneralPodAutoscalerLister returns a new GeneralPodAutoscalerLister.
func NewGeneralPodAutoscalerLister(indexer cache.Indexer) GeneralPodAutoscalerLister {
	return &generalPodAutoscalerLister{indexer: indexer}
}

// List lists all GeneralPodAutoscalers in the indexer.
func (s *generalPodAutoscalerLister) List(selector labels.Selector) (ret []*v1alpha1.GeneralPodAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GeneralPodAutoscaler))
	})
	return ret, err
}

// GeneralPodAutoscalers returns an object that can list and get GeneralPodAutoscalers.
func (s *generalPodAutoscalerLister) GeneralPodAutoscalers(namespace string) GeneralPodAutoscalerNamespaceLister {
	return generalPodAutoscalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GeneralPodAutoscalerNamespaceLister helps list and get GeneralPodAutoscalers.
type GeneralPodAutoscalerNamespaceLister interface {
	// List lists all GeneralPodAutoscalers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GeneralPodAutoscaler, err error)
	// Get retrieves the GeneralPodAutoscaler from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GeneralPodAutoscaler, error)
	GeneralPodAutoscalerNamespaceListerExpansion
}

// generalPodAutoscalerNamespaceLister implements the GeneralPodAutoscalerNamespaceLister
// interface.
type generalPodAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GeneralPodAutoscalers in the indexer for a given namespace.
func (s generalPodAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GeneralPodAutoscaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GeneralPodAutoscaler))
	})
	return ret, err
}

// Get retrieves the GeneralPodAutoscaler from the indexer for a given namespace and name.
func (s generalPodAutoscalerNamespaceLister) Get(name string) (*v1alpha1.GeneralPodAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("generalpodautoscaler"), name)
	}
	return obj.(*v1alpha1.GeneralPodAutoscaler), nil
}