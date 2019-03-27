/*
Copyright 2018 Openstorage.org

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterDomainUpdates implements ClusterDomainUpdateInterface
type FakeClusterDomainUpdates struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var clusterdomainupdatesResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "clusterdomainupdates"}

var clusterdomainupdatesKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "ClusterDomainUpdate"}

// Get takes name of the clusterDomainUpdate, and returns the corresponding clusterDomainUpdate object, and an error if there is any.
func (c *FakeClusterDomainUpdates) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterDomainUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterdomainupdatesResource, c.ns, name), &v1alpha1.ClusterDomainUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainUpdate), err
}

// List takes label and field selectors, and returns the list of ClusterDomainUpdates that match those selectors.
func (c *FakeClusterDomainUpdates) List(opts v1.ListOptions) (result *v1alpha1.ClusterDomainUpdateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterdomainupdatesResource, clusterdomainupdatesKind, c.ns, opts), &v1alpha1.ClusterDomainUpdateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterDomainUpdateList{ListMeta: obj.(*v1alpha1.ClusterDomainUpdateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterDomainUpdateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterDomainUpdates.
func (c *FakeClusterDomainUpdates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterdomainupdatesResource, c.ns, opts))

}

// Create takes the representation of a clusterDomainUpdate and creates it.  Returns the server's representation of the clusterDomainUpdate, and an error, if there is any.
func (c *FakeClusterDomainUpdates) Create(clusterDomainUpdate *v1alpha1.ClusterDomainUpdate) (result *v1alpha1.ClusterDomainUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterdomainupdatesResource, c.ns, clusterDomainUpdate), &v1alpha1.ClusterDomainUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainUpdate), err
}

// Update takes the representation of a clusterDomainUpdate and updates it. Returns the server's representation of the clusterDomainUpdate, and an error, if there is any.
func (c *FakeClusterDomainUpdates) Update(clusterDomainUpdate *v1alpha1.ClusterDomainUpdate) (result *v1alpha1.ClusterDomainUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterdomainupdatesResource, c.ns, clusterDomainUpdate), &v1alpha1.ClusterDomainUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainUpdate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterDomainUpdates) UpdateStatus(clusterDomainUpdate *v1alpha1.ClusterDomainUpdate) (*v1alpha1.ClusterDomainUpdate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterdomainupdatesResource, "status", c.ns, clusterDomainUpdate), &v1alpha1.ClusterDomainUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainUpdate), err
}

// Delete takes name of the clusterDomainUpdate and deletes it. Returns an error if one occurs.
func (c *FakeClusterDomainUpdates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterdomainupdatesResource, c.ns, name), &v1alpha1.ClusterDomainUpdate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterDomainUpdates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterdomainupdatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterDomainUpdateList{})
	return err
}

// Patch applies the patch and returns the patched clusterDomainUpdate.
func (c *FakeClusterDomainUpdates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterDomainUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterdomainupdatesResource, c.ns, name, data, subresources...), &v1alpha1.ClusterDomainUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDomainUpdate), err
}