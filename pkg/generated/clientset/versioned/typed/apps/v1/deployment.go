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

// Code generated by wranglertest. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "github.com/tomleb/wrangler_test/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeploymentsGetter has a method to return a DeploymentInterface.
// A group's client should implement this interface.
type DeploymentsGetter interface {
	Deployments(namespace string) DeploymentInterface
}

// DeploymentInterface has methods to work with Deployment resources.
type DeploymentInterface interface {
	Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (*v1.Deployment, error)
	Update(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (*v1.Deployment, error)
	UpdateStatus(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (*v1.Deployment, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Deployment, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DeploymentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Deployment, err error)
	GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, deploymentName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)

	DeploymentExpansion
}

// deployments implements DeploymentInterface
type deployments struct {
	client rest.Interface
	ns     string
}

// newDeployments returns a Deployments
func newDeployments(c *AppsV1Client, namespace string) *deployments {
	return &deployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *deployments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deployments) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *deployments) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Update(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deployment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deployments) UpdateStatus(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deployment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *deployments) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deployments) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deployment.
func (c *deployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetScale takes name of the deployment, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *deployments) GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *deployments) UpdateScale(ctx context.Context, deploymentName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}
