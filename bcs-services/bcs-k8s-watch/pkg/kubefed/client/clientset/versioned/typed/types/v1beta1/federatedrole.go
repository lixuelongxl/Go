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

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	scheme "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedRolesGetter has a method to return a FederatedRoleInterface.
// A group's client should implement this interface.
type FederatedRolesGetter interface {
	FederatedRoles(namespace string) FederatedRoleInterface
}

// FederatedRoleInterface has methods to work with FederatedRole resources.
type FederatedRoleInterface interface {
	Create(*v1beta1.FederatedRole) (*v1beta1.FederatedRole, error)
	Update(*v1beta1.FederatedRole) (*v1beta1.FederatedRole, error)
	UpdateStatus(*v1beta1.FederatedRole) (*v1beta1.FederatedRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedRole, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedRole, err error)
	FederatedRoleExpansion
}

// federatedRoles implements FederatedRoleInterface
type federatedRoles struct {
	client rest.Interface
	ns     string
}

// newFederatedRoles returns a FederatedRoles
func newFederatedRoles(c *TypesV1beta1Client, namespace string) *federatedRoles {
	return &federatedRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedRole, and returns the corresponding federatedRole object, and an error if there is any.
func (c *federatedRoles) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedRole, err error) {
	result = &v1beta1.FederatedRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedRoles that match those selectors.
func (c *federatedRoles) List(opts v1.ListOptions) (result *v1beta1.FederatedRoleList, err error) {
	result = &v1beta1.FederatedRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedRoles.
func (c *federatedRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedRole and creates it.  Returns the server's representation of the federatedRole, and an error, if there is any.
func (c *federatedRoles) Create(federatedRole *v1beta1.FederatedRole) (result *v1beta1.FederatedRole, err error) {
	result = &v1beta1.FederatedRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedroles").
		Body(federatedRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedRole and updates it. Returns the server's representation of the federatedRole, and an error, if there is any.
func (c *federatedRoles) Update(federatedRole *v1beta1.FederatedRole) (result *v1beta1.FederatedRole, err error) {
	result = &v1beta1.FederatedRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedroles").
		Name(federatedRole.Name).
		Body(federatedRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedRoles) UpdateStatus(federatedRole *v1beta1.FederatedRole) (result *v1beta1.FederatedRole, err error) {
	result = &v1beta1.FederatedRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedroles").
		Name(federatedRole.Name).
		SubResource("status").
		Body(federatedRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedRole and deletes it. Returns an error if one occurs.
func (c *federatedRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedRole.
func (c *federatedRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedRole, err error) {
	result = &v1beta1.FederatedRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
