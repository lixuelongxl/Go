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
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/scheduling/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicaSchedulingPreferences implements ReplicaSchedulingPreferenceInterface
type FakeReplicaSchedulingPreferences struct {
	Fake *FakeSchedulingV1alpha1
	ns   string
}

var replicaschedulingpreferencesResource = schema.GroupVersionResource{Group: "scheduling.kubefed.io", Version: "v1alpha1", Resource: "replicaschedulingpreferences"}

var replicaschedulingpreferencesKind = schema.GroupVersionKind{Group: "scheduling.kubefed.io", Version: "v1alpha1", Kind: "ReplicaSchedulingPreference"}

// Get takes name of the replicaSchedulingPreference, and returns the corresponding replicaSchedulingPreference object, and an error if there is any.
func (c *FakeReplicaSchedulingPreferences) Get(name string, options v1.GetOptions) (result *v1alpha1.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicaschedulingpreferencesResource, c.ns, name), &v1alpha1.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicaSchedulingPreference), err
}

// List takes label and field selectors, and returns the list of ReplicaSchedulingPreferences that match those selectors.
func (c *FakeReplicaSchedulingPreferences) List(opts v1.ListOptions) (result *v1alpha1.ReplicaSchedulingPreferenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicaschedulingpreferencesResource, replicaschedulingpreferencesKind, c.ns, opts), &v1alpha1.ReplicaSchedulingPreferenceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReplicaSchedulingPreferenceList{}
	for _, item := range obj.(*v1alpha1.ReplicaSchedulingPreferenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicaSchedulingPreferences.
func (c *FakeReplicaSchedulingPreferences) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicaschedulingpreferencesResource, c.ns, opts))

}

// Create takes the representation of a replicaSchedulingPreference and creates it.  Returns the server's representation of the replicaSchedulingPreference, and an error, if there is any.
func (c *FakeReplicaSchedulingPreferences) Create(replicaSchedulingPreference *v1alpha1.ReplicaSchedulingPreference) (result *v1alpha1.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicaschedulingpreferencesResource, c.ns, replicaSchedulingPreference), &v1alpha1.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicaSchedulingPreference), err
}

// Update takes the representation of a replicaSchedulingPreference and updates it. Returns the server's representation of the replicaSchedulingPreference, and an error, if there is any.
func (c *FakeReplicaSchedulingPreferences) Update(replicaSchedulingPreference *v1alpha1.ReplicaSchedulingPreference) (result *v1alpha1.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicaschedulingpreferencesResource, c.ns, replicaSchedulingPreference), &v1alpha1.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicaSchedulingPreference), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicaSchedulingPreferences) UpdateStatus(replicaSchedulingPreference *v1alpha1.ReplicaSchedulingPreference) (*v1alpha1.ReplicaSchedulingPreference, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicaschedulingpreferencesResource, "status", c.ns, replicaSchedulingPreference), &v1alpha1.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicaSchedulingPreference), err
}

// Delete takes name of the replicaSchedulingPreference and deletes it. Returns an error if one occurs.
func (c *FakeReplicaSchedulingPreferences) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicaschedulingpreferencesResource, c.ns, name), &v1alpha1.ReplicaSchedulingPreference{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicaSchedulingPreferences) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicaschedulingpreferencesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReplicaSchedulingPreferenceList{})
	return err
}

// Patch applies the patch and returns the patched replicaSchedulingPreference.
func (c *FakeReplicaSchedulingPreferences) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReplicaSchedulingPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicaschedulingpreferencesResource, c.ns, name, data, subresources...), &v1alpha1.ReplicaSchedulingPreference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicaSchedulingPreference), err
}
