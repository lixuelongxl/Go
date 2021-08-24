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

package v2

import (
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-mesos/kubebkbcsv2/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AdmissionWebhookConfigurations returns a AdmissionWebhookConfigurationInformer.
	AdmissionWebhookConfigurations() AdmissionWebhookConfigurationInformer
	// Agents returns a AgentInformer.
	Agents() AgentInformer
	// AgentSchedInfos returns a AgentSchedInfoInformer.
	AgentSchedInfos() AgentSchedInfoInformer
	// Applications returns a ApplicationInformer.
	Applications() ApplicationInformer
	// BcsClusterAgentSettings returns a BcsClusterAgentSettingInformer.
	BcsClusterAgentSettings() BcsClusterAgentSettingInformer
	// BcsCommandInfos returns a BcsCommandInfoInformer.
	BcsCommandInfos() BcsCommandInfoInformer
	// BcsConfigMaps returns a BcsConfigMapInformer.
	BcsConfigMaps() BcsConfigMapInformer
	// BcsDaemonsets returns a BcsDaemonsetInformer.
	BcsDaemonsets() BcsDaemonsetInformer
	// BcsEndpoints returns a BcsEndpointInformer.
	BcsEndpoints() BcsEndpointInformer
	// BcsSecrets returns a BcsSecretInformer.
	BcsSecrets() BcsSecretInformer
	// BcsServices returns a BcsServiceInformer.
	BcsServices() BcsServiceInformer
	// BcsTransactions returns a BcsTransactionInformer.
	BcsTransactions() BcsTransactionInformer
	// Crds returns a CrdInformer.
	Crds() CrdInformer
	// Crrs returns a CrrInformer.
	Crrs() CrrInformer
	// Deployments returns a DeploymentInformer.
	Deployments() DeploymentInformer
	// Frameworks returns a FrameworkInformer.
	Frameworks() FrameworkInformer
	// Tasks returns a TaskInformer.
	Tasks() TaskInformer
	// TaskGroups returns a TaskGroupInformer.
	TaskGroups() TaskGroupInformer
	// Versions returns a VersionInformer.
	Versions() VersionInformer
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

// AdmissionWebhookConfigurations returns a AdmissionWebhookConfigurationInformer.
func (v *version) AdmissionWebhookConfigurations() AdmissionWebhookConfigurationInformer {
	return &admissionWebhookConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Agents returns a AgentInformer.
func (v *version) Agents() AgentInformer {
	return &agentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AgentSchedInfos returns a AgentSchedInfoInformer.
func (v *version) AgentSchedInfos() AgentSchedInfoInformer {
	return &agentSchedInfoInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Applications returns a ApplicationInformer.
func (v *version) Applications() ApplicationInformer {
	return &applicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsClusterAgentSettings returns a BcsClusterAgentSettingInformer.
func (v *version) BcsClusterAgentSettings() BcsClusterAgentSettingInformer {
	return &bcsClusterAgentSettingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsCommandInfos returns a BcsCommandInfoInformer.
func (v *version) BcsCommandInfos() BcsCommandInfoInformer {
	return &bcsCommandInfoInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsConfigMaps returns a BcsConfigMapInformer.
func (v *version) BcsConfigMaps() BcsConfigMapInformer {
	return &bcsConfigMapInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsDaemonsets returns a BcsDaemonsetInformer.
func (v *version) BcsDaemonsets() BcsDaemonsetInformer {
	return &bcsDaemonsetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsEndpoints returns a BcsEndpointInformer.
func (v *version) BcsEndpoints() BcsEndpointInformer {
	return &bcsEndpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsSecrets returns a BcsSecretInformer.
func (v *version) BcsSecrets() BcsSecretInformer {
	return &bcsSecretInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsServices returns a BcsServiceInformer.
func (v *version) BcsServices() BcsServiceInformer {
	return &bcsServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BcsTransactions returns a BcsTransactionInformer.
func (v *version) BcsTransactions() BcsTransactionInformer {
	return &bcsTransactionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Crds returns a CrdInformer.
func (v *version) Crds() CrdInformer {
	return &crdInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Crrs returns a CrrInformer.
func (v *version) Crrs() CrrInformer {
	return &crrInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Deployments returns a DeploymentInformer.
func (v *version) Deployments() DeploymentInformer {
	return &deploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Frameworks returns a FrameworkInformer.
func (v *version) Frameworks() FrameworkInformer {
	return &frameworkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Tasks returns a TaskInformer.
func (v *version) Tasks() TaskInformer {
	return &taskInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TaskGroups returns a TaskGroupInformer.
func (v *version) TaskGroups() TaskGroupInformer {
	return &taskGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Versions returns a VersionInformer.
func (v *version) Versions() VersionInformer {
	return &versionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}