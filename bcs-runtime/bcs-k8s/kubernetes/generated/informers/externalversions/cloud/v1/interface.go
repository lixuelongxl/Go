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

package v1

import (
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CloudIPs returns a CloudIPInformer.
	CloudIPs() CloudIPInformer
	// CloudIPQuotas returns a CloudIPQuotaInformer.
	CloudIPQuotas() CloudIPQuotaInformer
	// CloudSubnets returns a CloudSubnetInformer.
	CloudSubnets() CloudSubnetInformer
	// NodeNetworks returns a NodeNetworkInformer.
	NodeNetworks() NodeNetworkInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CloudIPs returns a CloudIPInformer.
func (v *version) CloudIPs() CloudIPInformer {
	return &cloudIPInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudIPQuotas returns a CloudIPQuotaInformer.
func (v *version) CloudIPQuotas() CloudIPQuotaInformer {
	return &cloudIPQuotaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudSubnets returns a CloudSubnetInformer.
func (v *version) CloudSubnets() CloudSubnetInformer {
	return &cloudSubnetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NodeNetworks returns a NodeNetworkInformer.
func (v *version) NodeNetworks() NodeNetworkInformer {
	return &nodeNetworkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
