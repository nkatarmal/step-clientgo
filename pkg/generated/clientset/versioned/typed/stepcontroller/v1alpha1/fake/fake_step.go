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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/nkatarmal/step-clientgo/pkg/apis/stepcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSteps implements StepInterface
type FakeSteps struct {
	Fake *FakeStepcontrollerV1alpha1
	ns   string
}

var stepsResource = schema.GroupVersionResource{Group: "stepcontroller.k8s.io", Version: "v1alpha1", Resource: "steps"}

var stepsKind = schema.GroupVersionKind{Group: "stepcontroller.k8s.io", Version: "v1alpha1", Kind: "Step"}

// Get takes name of the step, and returns the corresponding step object, and an error if there is any.
func (c *FakeSteps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Step, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(stepsResource, c.ns, name), &v1alpha1.Step{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Step), err
}

// List takes label and field selectors, and returns the list of Steps that match those selectors.
func (c *FakeSteps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StepList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(stepsResource, stepsKind, c.ns, opts), &v1alpha1.StepList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StepList{ListMeta: obj.(*v1alpha1.StepList).ListMeta}
	for _, item := range obj.(*v1alpha1.StepList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested steps.
func (c *FakeSteps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(stepsResource, c.ns, opts))

}

// Create takes the representation of a step and creates it.  Returns the server's representation of the step, and an error, if there is any.
func (c *FakeSteps) Create(ctx context.Context, step *v1alpha1.Step, opts v1.CreateOptions) (result *v1alpha1.Step, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(stepsResource, c.ns, step), &v1alpha1.Step{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Step), err
}

// Update takes the representation of a step and updates it. Returns the server's representation of the step, and an error, if there is any.
func (c *FakeSteps) Update(ctx context.Context, step *v1alpha1.Step, opts v1.UpdateOptions) (result *v1alpha1.Step, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(stepsResource, c.ns, step), &v1alpha1.Step{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Step), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSteps) UpdateStatus(ctx context.Context, step *v1alpha1.Step, opts v1.UpdateOptions) (*v1alpha1.Step, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(stepsResource, "status", c.ns, step), &v1alpha1.Step{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Step), err
}

// Delete takes name of the step and deletes it. Returns an error if one occurs.
func (c *FakeSteps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(stepsResource, c.ns, name), &v1alpha1.Step{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSteps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(stepsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StepList{})
	return err
}

// Patch applies the patch and returns the patched step.
func (c *FakeSteps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Step, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(stepsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Step{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Step), err
}
