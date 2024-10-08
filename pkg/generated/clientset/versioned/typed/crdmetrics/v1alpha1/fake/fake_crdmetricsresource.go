/*
Copyright 2024 The Kubernetes crdmetrics Authors.

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
	"context"

	v1alpha1 "github.com/rexagod/crdmetrics/pkg/apis/crdmetrics/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCRDMetricsResources implements CRDMetricsResourceInterface
type FakeCRDMetricsResources struct {
	Fake *FakeCrdmetricsV1alpha1
	ns   string
}

var crdmetricsresourcesResource = v1alpha1.SchemeGroupVersion.WithResource("crdmetricsresources")

var crdmetricsresourcesKind = v1alpha1.SchemeGroupVersion.WithKind("CRDMetricsResource")

// Get takes name of the cRDMetricsResource, and returns the corresponding cRDMetricsResource object, and an error if there is any.
func (c *FakeCRDMetricsResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CRDMetricsResource, err error) {
	emptyResult := &v1alpha1.CRDMetricsResource{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(crdmetricsresourcesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CRDMetricsResource), err
}

// List takes label and field selectors, and returns the list of CRDMetricsResources that match those selectors.
func (c *FakeCRDMetricsResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CRDMetricsResourceList, err error) {
	emptyResult := &v1alpha1.CRDMetricsResourceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(crdmetricsresourcesResource, crdmetricsresourcesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CRDMetricsResourceList{ListMeta: obj.(*v1alpha1.CRDMetricsResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CRDMetricsResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cRDMetricsResources.
func (c *FakeCRDMetricsResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(crdmetricsresourcesResource, c.ns, opts))

}

// Create takes the representation of a cRDMetricsResource and creates it.  Returns the server's representation of the cRDMetricsResource, and an error, if there is any.
func (c *FakeCRDMetricsResources) Create(ctx context.Context, cRDMetricsResource *v1alpha1.CRDMetricsResource, opts v1.CreateOptions) (result *v1alpha1.CRDMetricsResource, err error) {
	emptyResult := &v1alpha1.CRDMetricsResource{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(crdmetricsresourcesResource, c.ns, cRDMetricsResource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CRDMetricsResource), err
}

// Update takes the representation of a cRDMetricsResource and updates it. Returns the server's representation of the cRDMetricsResource, and an error, if there is any.
func (c *FakeCRDMetricsResources) Update(ctx context.Context, cRDMetricsResource *v1alpha1.CRDMetricsResource, opts v1.UpdateOptions) (result *v1alpha1.CRDMetricsResource, err error) {
	emptyResult := &v1alpha1.CRDMetricsResource{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(crdmetricsresourcesResource, c.ns, cRDMetricsResource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CRDMetricsResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCRDMetricsResources) UpdateStatus(ctx context.Context, cRDMetricsResource *v1alpha1.CRDMetricsResource, opts v1.UpdateOptions) (result *v1alpha1.CRDMetricsResource, err error) {
	emptyResult := &v1alpha1.CRDMetricsResource{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(crdmetricsresourcesResource, "status", c.ns, cRDMetricsResource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CRDMetricsResource), err
}

// Delete takes name of the cRDMetricsResource and deletes it. Returns an error if one occurs.
func (c *FakeCRDMetricsResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(crdmetricsresourcesResource, c.ns, name, opts), &v1alpha1.CRDMetricsResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCRDMetricsResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(crdmetricsresourcesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CRDMetricsResourceList{})
	return err
}

// Patch applies the patch and returns the patched cRDMetricsResource.
func (c *FakeCRDMetricsResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CRDMetricsResource, err error) {
	emptyResult := &v1alpha1.CRDMetricsResource{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(crdmetricsresourcesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CRDMetricsResource), err
}
