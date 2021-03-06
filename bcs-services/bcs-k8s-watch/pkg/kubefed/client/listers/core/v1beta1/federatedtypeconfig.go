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
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/core/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedTypeConfigLister helps list FederatedTypeConfigs.
type FederatedTypeConfigLister interface {
	// List lists all FederatedTypeConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedTypeConfig, err error)
	// FederatedTypeConfigs returns an object that can list and get FederatedTypeConfigs.
	FederatedTypeConfigs(namespace string) FederatedTypeConfigNamespaceLister
	FederatedTypeConfigListerExpansion
}

// federatedTypeConfigLister implements the FederatedTypeConfigLister interface.
type federatedTypeConfigLister struct {
	indexer cache.Indexer
}

// NewFederatedTypeConfigLister returns a new FederatedTypeConfigLister.
func NewFederatedTypeConfigLister(indexer cache.Indexer) FederatedTypeConfigLister {
	return &federatedTypeConfigLister{indexer: indexer}
}

// List lists all FederatedTypeConfigs in the indexer.
func (s *federatedTypeConfigLister) List(selector labels.Selector) (ret []*v1beta1.FederatedTypeConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedTypeConfig))
	})
	return ret, err
}

// FederatedTypeConfigs returns an object that can list and get FederatedTypeConfigs.
func (s *federatedTypeConfigLister) FederatedTypeConfigs(namespace string) FederatedTypeConfigNamespaceLister {
	return federatedTypeConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedTypeConfigNamespaceLister helps list and get FederatedTypeConfigs.
type FederatedTypeConfigNamespaceLister interface {
	// List lists all FederatedTypeConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedTypeConfig, err error)
	// Get retrieves the FederatedTypeConfig from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedTypeConfig, error)
	FederatedTypeConfigNamespaceListerExpansion
}

// federatedTypeConfigNamespaceLister implements the FederatedTypeConfigNamespaceLister
// interface.
type federatedTypeConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedTypeConfigs in the indexer for a given namespace.
func (s federatedTypeConfigNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedTypeConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedTypeConfig))
	})
	return ret, err
}

// Get retrieves the FederatedTypeConfig from the indexer for a given namespace and name.
func (s federatedTypeConfigNamespaceLister) Get(name string) (*v1beta1.FederatedTypeConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedtypeconfig"), name)
	}
	return obj.(*v1beta1.FederatedTypeConfig), nil
}
