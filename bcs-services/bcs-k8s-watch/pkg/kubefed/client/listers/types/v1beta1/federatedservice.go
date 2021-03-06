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

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedServiceLister helps list FederatedServices.
type FederatedServiceLister interface {
	// List lists all FederatedServices in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedService, err error)
	// FederatedServices returns an object that can list and get FederatedServices.
	FederatedServices(namespace string) FederatedServiceNamespaceLister
	FederatedServiceListerExpansion
}

// federatedServiceLister implements the FederatedServiceLister interface.
type federatedServiceLister struct {
	indexer cache.Indexer
}

// NewFederatedServiceLister returns a new FederatedServiceLister.
func NewFederatedServiceLister(indexer cache.Indexer) FederatedServiceLister {
	return &federatedServiceLister{indexer: indexer}
}

// List lists all FederatedServices in the indexer.
func (s *federatedServiceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedService))
	})
	return ret, err
}

// FederatedServices returns an object that can list and get FederatedServices.
func (s *federatedServiceLister) FederatedServices(namespace string) FederatedServiceNamespaceLister {
	return federatedServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedServiceNamespaceLister helps list and get FederatedServices.
type FederatedServiceNamespaceLister interface {
	// List lists all FederatedServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedService, err error)
	// Get retrieves the FederatedService from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedService, error)
	FederatedServiceNamespaceListerExpansion
}

// federatedServiceNamespaceLister implements the FederatedServiceNamespaceLister
// interface.
type federatedServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedServices in the indexer for a given namespace.
func (s federatedServiceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedService))
	})
	return ret, err
}

// Get retrieves the FederatedService from the indexer for a given namespace and name.
func (s federatedServiceNamespaceLister) Get(name string) (*v1beta1.FederatedService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedservice"), name)
	}
	return obj.(*v1beta1.FederatedService), nil
}
