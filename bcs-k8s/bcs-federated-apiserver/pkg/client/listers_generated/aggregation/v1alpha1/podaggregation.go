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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-federated-apiserver/pkg/apis/aggregation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodAggregationLister helps list PodAggregations.
type PodAggregationLister interface {
	// List lists all PodAggregations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PodAggregation, err error)
	// PodAggregations returns an object that can list and get PodAggregations.
	PodAggregations(namespace string) PodAggregationNamespaceLister
	PodAggregationListerExpansion
}

// podAggregationLister implements the PodAggregationLister interface.
type podAggregationLister struct {
	indexer cache.Indexer
}

// NewPodAggregationLister returns a new PodAggregationLister.
func NewPodAggregationLister(indexer cache.Indexer) PodAggregationLister {
	return &podAggregationLister{indexer: indexer}
}

// List lists all PodAggregations in the indexer.
func (s *podAggregationLister) List(selector labels.Selector) (ret []*v1alpha1.PodAggregation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodAggregation))
	})
	return ret, err
}

// PodAggregations returns an object that can list and get PodAggregations.
func (s *podAggregationLister) PodAggregations(namespace string) PodAggregationNamespaceLister {
	return podAggregationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodAggregationNamespaceLister helps list and get PodAggregations.
type PodAggregationNamespaceLister interface {
	// List lists all PodAggregations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PodAggregation, err error)
	// Get retrieves the PodAggregation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PodAggregation, error)
	PodAggregationNamespaceListerExpansion
}

// podAggregationNamespaceLister implements the PodAggregationNamespaceLister
// interface.
type podAggregationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodAggregations in the indexer for a given namespace.
func (s podAggregationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PodAggregation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodAggregation))
	})
	return ret, err
}

// Get retrieves the PodAggregation from the indexer for a given namespace and name.
func (s podAggregationNamespaceLister) Get(name string) (*v1alpha1.PodAggregation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podaggregation"), name)
	}
	return obj.(*v1alpha1.PodAggregation), nil
}