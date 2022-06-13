/*


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

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/VictoriaMetrics/operator/api/client/versioned/scheme"
	v1beta1 "github.com/VictoriaMetrics/operator/api/core/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VMAgentsGetter has a method to return a VMAgentInterface.
// A group's client should implement this interface.
type VMAgentsGetter interface {
	VMAgents(namespace string) VMAgentInterface
}

// VMAgentInterface has methods to work with VMAgent resources.
type VMAgentInterface interface {
	Create(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.CreateOptions) (*v1beta1.VMAgent, error)
	Update(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.UpdateOptions) (*v1beta1.VMAgent, error)
	UpdateStatus(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.UpdateOptions) (*v1beta1.VMAgent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VMAgent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VMAgentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMAgent, err error)
	VMAgentExpansion
}

// vMAgents implements VMAgentInterface
type vMAgents struct {
	client rest.Interface
	ns     string
}

// newVMAgents returns a VMAgents
func newVMAgents(c *CoreV1beta1Client, namespace string) *vMAgents {
	return &vMAgents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vMAgent, and returns the corresponding vMAgent object, and an error if there is any.
func (c *vMAgents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMAgent, err error) {
	result = &v1beta1.VMAgent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmagents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VMAgents that match those selectors.
func (c *vMAgents) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMAgentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VMAgentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vMAgents.
func (c *vMAgents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vmagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vMAgent and creates it.  Returns the server's representation of the vMAgent, and an error, if there is any.
func (c *vMAgents) Create(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.CreateOptions) (result *v1beta1.VMAgent, err error) {
	result = &v1beta1.VMAgent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vmagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMAgent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vMAgent and updates it. Returns the server's representation of the vMAgent, and an error, if there is any.
func (c *vMAgents) Update(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.UpdateOptions) (result *v1beta1.VMAgent, err error) {
	result = &v1beta1.VMAgent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmagents").
		Name(vMAgent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMAgent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vMAgents) UpdateStatus(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.UpdateOptions) (result *v1beta1.VMAgent, err error) {
	result = &v1beta1.VMAgent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmagents").
		Name(vMAgent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMAgent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vMAgent and deletes it. Returns an error if one occurs.
func (c *vMAgents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmagents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vMAgents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmagents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vMAgent.
func (c *vMAgents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMAgent, err error) {
	result = &v1beta1.VMAgent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vmagents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
