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

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/apis/mesh/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppSvcLister helps list AppSvcs.
type AppSvcLister interface {
	// List lists all AppSvcs in the indexer.
	List(selector labels.Selector) (ret []*v1.AppSvc, err error)
	// AppSvcs returns an object that can list and get AppSvcs.
	AppSvcs(namespace string) AppSvcNamespaceLister
	AppSvcListerExpansion
}

// appSvcLister implements the AppSvcLister interface.
type appSvcLister struct {
	indexer cache.Indexer
}

// NewAppSvcLister returns a new AppSvcLister.
func NewAppSvcLister(indexer cache.Indexer) AppSvcLister {
	return &appSvcLister{indexer: indexer}
}

// List lists all AppSvcs in the indexer.
func (s *appSvcLister) List(selector labels.Selector) (ret []*v1.AppSvc, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AppSvc))
	})
	return ret, err
}

// AppSvcs returns an object that can list and get AppSvcs.
func (s *appSvcLister) AppSvcs(namespace string) AppSvcNamespaceLister {
	return appSvcNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppSvcNamespaceLister helps list and get AppSvcs.
type AppSvcNamespaceLister interface {
	// List lists all AppSvcs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AppSvc, err error)
	// Get retrieves the AppSvc from the indexer for a given namespace and name.
	Get(name string) (*v1.AppSvc, error)
	AppSvcNamespaceListerExpansion
}

// appSvcNamespaceLister implements the AppSvcNamespaceLister
// interface.
type appSvcNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppSvcs in the indexer for a given namespace.
func (s appSvcNamespaceLister) List(selector labels.Selector) (ret []*v1.AppSvc, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AppSvc))
	})
	return ret, err
}

// Get retrieves the AppSvc from the indexer for a given namespace and name.
func (s appSvcNamespaceLister) Get(name string) (*v1.AppSvc, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("appsvc"), name)
	}
	return obj.(*v1.AppSvc), nil
}
