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

// FakeFederatedClusterRoleBindings implements FederatedClusterRoleBindingInterface
type FakeFederatedClusterRoleBindings struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedclusterrolebindingsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedclusterrolebindings"}

var federatedclusterrolebindingsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedClusterRoleBinding"}

// Get takes name of the federatedClusterRoleBinding, and returns the corresponding federatedClusterRoleBinding object, and an error if there is any.
func (c *FakeFederatedClusterRoleBindings) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedclusterrolebindingsResource, c.ns, name), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// List takes label and field selectors, and returns the list of FederatedClusterRoleBindings that match those selectors.
func (c *FakeFederatedClusterRoleBindings) List(opts v1.ListOptions) (result *v1beta1.FederatedClusterRoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedclusterrolebindingsResource, federatedclusterrolebindingsKind, c.ns, opts), &v1beta1.FederatedClusterRoleBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedClusterRoleBindingList{}
	for _, item := range obj.(*v1beta1.FederatedClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedClusterRoleBindings.
func (c *FakeFederatedClusterRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedclusterrolebindingsResource, c.ns, opts))

}

// Create takes the representation of a federatedClusterRoleBinding and creates it.  Returns the server's representation of the federatedClusterRoleBinding, and an error, if there is any.
func (c *FakeFederatedClusterRoleBindings) Create(federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedclusterrolebindingsResource, c.ns, federatedClusterRoleBinding), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// Update takes the representation of a federatedClusterRoleBinding and updates it. Returns the server's representation of the federatedClusterRoleBinding, and an error, if there is any.
func (c *FakeFederatedClusterRoleBindings) Update(federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedclusterrolebindingsResource, c.ns, federatedClusterRoleBinding), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedClusterRoleBindings) UpdateStatus(federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding) (*v1beta1.FederatedClusterRoleBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedclusterrolebindingsResource, "status", c.ns, federatedClusterRoleBinding), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// Delete takes name of the federatedClusterRoleBinding and deletes it. Returns an error if one occurs.
func (c *FakeFederatedClusterRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedclusterrolebindingsResource, c.ns, name), &v1beta1.FederatedClusterRoleBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedClusterRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedclusterrolebindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedClusterRoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched federatedClusterRoleBinding.
func (c *FakeFederatedClusterRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedclusterrolebindingsResource, c.ns, name, data, subresources...), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}
