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

// FakeClusterAuthTokens implements ClusterAuthTokenInterface
type FakeClusterAuthTokens struct {
	Fake *FakeClusterV3
	ns   string
}

var clusterauthtokensResource = v3.SchemeGroupVersion.WithResource("clusterauthtokens")

var clusterauthtokensKind = v3.SchemeGroupVersion.WithKind("ClusterAuthToken")

// Get takes name of the clusterAuthToken, and returns the corresponding clusterAuthToken object, and an error if there is any.
func (c *FakeClusterAuthTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterAuthToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterauthtokensResource, c.ns, name), &v3.ClusterAuthToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAuthToken), err
}

// List takes label and field selectors, and returns the list of ClusterAuthTokens that match those selectors.
func (c *FakeClusterAuthTokens) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterAuthTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterauthtokensResource, clusterauthtokensKind, c.ns, opts), &v3.ClusterAuthTokenList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterAuthTokenList{ListMeta: obj.(*v3.ClusterAuthTokenList).ListMeta}
	for _, item := range obj.(*v3.ClusterAuthTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterAuthTokens.
func (c *FakeClusterAuthTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterauthtokensResource, c.ns, opts))

}

// Create takes the representation of a clusterAuthToken and creates it.  Returns the server's representation of the clusterAuthToken, and an error, if there is any.
func (c *FakeClusterAuthTokens) Create(ctx context.Context, clusterAuthToken *v3.ClusterAuthToken, opts v1.CreateOptions) (result *v3.ClusterAuthToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterauthtokensResource, c.ns, clusterAuthToken), &v3.ClusterAuthToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAuthToken), err
}

// Update takes the representation of a clusterAuthToken and updates it. Returns the server's representation of the clusterAuthToken, and an error, if there is any.
func (c *FakeClusterAuthTokens) Update(ctx context.Context, clusterAuthToken *v3.ClusterAuthToken, opts v1.UpdateOptions) (result *v3.ClusterAuthToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterauthtokensResource, c.ns, clusterAuthToken), &v3.ClusterAuthToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAuthToken), err
}

// Delete takes name of the clusterAuthToken and deletes it. Returns an error if one occurs.
func (c *FakeClusterAuthTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusterauthtokensResource, c.ns, name, opts), &v3.ClusterAuthToken{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterAuthTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterauthtokensResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterAuthTokenList{})
	return err
}

// Patch applies the patch and returns the patched clusterAuthToken.
func (c *FakeClusterAuthTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterAuthToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterauthtokensResource, c.ns, name, pt, data, subresources...), &v3.ClusterAuthToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAuthToken), err
}
