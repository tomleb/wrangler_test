/*
Copyright The Kubernetes Authors.

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

// Code generated by wrangler_test. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/tomleb/wrangler_test/pkg/apis/cluster.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterUserAttributes implements ClusterUserAttributeInterface
type FakeClusterUserAttributes struct {
	Fake *FakeClusterV3
	ns   string
}

var clusteruserattributesResource = v3.SchemeGroupVersion.WithResource("clusteruserattributes")

var clusteruserattributesKind = v3.SchemeGroupVersion.WithKind("ClusterUserAttribute")

// Get takes name of the clusterUserAttribute, and returns the corresponding clusterUserAttribute object, and an error if there is any.
func (c *FakeClusterUserAttributes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterUserAttribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusteruserattributesResource, c.ns, name), &v3.ClusterUserAttribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterUserAttribute), err
}

// List takes label and field selectors, and returns the list of ClusterUserAttributes that match those selectors.
func (c *FakeClusterUserAttributes) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterUserAttributeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusteruserattributesResource, clusteruserattributesKind, c.ns, opts), &v3.ClusterUserAttributeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterUserAttributeList{ListMeta: obj.(*v3.ClusterUserAttributeList).ListMeta}
	for _, item := range obj.(*v3.ClusterUserAttributeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterUserAttributes.
func (c *FakeClusterUserAttributes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusteruserattributesResource, c.ns, opts))

}

// Create takes the representation of a clusterUserAttribute and creates it.  Returns the server's representation of the clusterUserAttribute, and an error, if there is any.
func (c *FakeClusterUserAttributes) Create(ctx context.Context, clusterUserAttribute *v3.ClusterUserAttribute, opts v1.CreateOptions) (result *v3.ClusterUserAttribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusteruserattributesResource, c.ns, clusterUserAttribute), &v3.ClusterUserAttribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterUserAttribute), err
}

// Update takes the representation of a clusterUserAttribute and updates it. Returns the server's representation of the clusterUserAttribute, and an error, if there is any.
func (c *FakeClusterUserAttributes) Update(ctx context.Context, clusterUserAttribute *v3.ClusterUserAttribute, opts v1.UpdateOptions) (result *v3.ClusterUserAttribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusteruserattributesResource, c.ns, clusterUserAttribute), &v3.ClusterUserAttribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterUserAttribute), err
}

// Delete takes name of the clusterUserAttribute and deletes it. Returns an error if one occurs.
func (c *FakeClusterUserAttributes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusteruserattributesResource, c.ns, name, opts), &v3.ClusterUserAttribute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterUserAttributes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusteruserattributesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterUserAttributeList{})
	return err
}

// Patch applies the patch and returns the patched clusterUserAttribute.
func (c *FakeClusterUserAttributes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterUserAttribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusteruserattributesResource, c.ns, name, pt, data, subresources...), &v3.ClusterUserAttribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterUserAttribute), err
}
