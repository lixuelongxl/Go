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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedDeployments implements FederatedDeploymentInterface
type FakeFederatedDeployments struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federateddeploymentsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federateddeployments"}

var federateddeploymentsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedDeployment"}

// Get takes name of the federatedDeployment, and returns the corresponding federatedDeployment object, and an error if there is any.
func (c *FakeFederatedDeployments) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federateddeploymentsResource, c.ns, name), &v1beta1.FederatedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDeployment), err
}

// List takes label and field selectors, and returns the list of FederatedDeployments that match those selectors.
func (c *FakeFederatedDeployments) List(opts v1.ListOptions) (result *v1beta1.FederatedDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federateddeploymentsResource, federateddeploymentsKind, c.ns, opts), &v1beta1.FederatedDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedDeploymentList{}
	for _, item := range obj.(*v1beta1.FederatedDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedDeployments.
func (c *FakeFederatedDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federateddeploymentsResource, c.ns, opts))

}

// Create takes the representation of a federatedDeployment and creates it.  Returns the server's representation of the federatedDeployment, and an error, if there is any.
func (c *FakeFederatedDeployments) Create(federatedDeployment *v1beta1.FederatedDeployment) (result *v1beta1.FederatedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federateddeploymentsResource, c.ns, federatedDeployment), &v1beta1.FederatedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDeployment), err
}

// Update takes the representation of a federatedDeployment and updates it. Returns the server's representation of the federatedDeployment, and an error, if there is any.
func (c *FakeFederatedDeployments) Update(federatedDeployment *v1beta1.FederatedDeployment) (result *v1beta1.FederatedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federateddeploymentsResource, c.ns, federatedDeployment), &v1beta1.FederatedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedDeployments) UpdateStatus(federatedDeployment *v1beta1.FederatedDeployment) (*v1beta1.FederatedDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federateddeploymentsResource, "status", c.ns, federatedDeployment), &v1beta1.FederatedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDeployment), err
}

// Delete takes name of the federatedDeployment and deletes it. Returns an error if one occurs.
func (c *FakeFederatedDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federateddeploymentsResource, c.ns, name), &v1beta1.FederatedDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federateddeploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched federatedDeployment.
func (c *FakeFederatedDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federateddeploymentsResource, c.ns, name, data, subresources...), &v1beta1.FederatedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedDeployment), err
}
