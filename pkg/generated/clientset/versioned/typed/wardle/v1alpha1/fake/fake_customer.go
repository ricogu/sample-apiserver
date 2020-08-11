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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1"
)

// FakeCustomers implements CustomerInterface
type FakeCustomers struct {
	Fake *FakeWardleV1alpha1
	ns   string
}

var customersResource = schema.GroupVersionResource{Group: "wardle.example.com", Version: "v1alpha1", Resource: "customers"}

var customersKind = schema.GroupVersionKind{Group: "wardle.example.com", Version: "v1alpha1", Kind: "Customer"}

// Get takes name of the customer, and returns the corresponding customer object, and an error if there is any.
func (c *FakeCustomers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Customer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(customersResource, c.ns, name), &v1alpha1.Customer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Customer), err
}

// List takes label and field selectors, and returns the list of Customers that match those selectors.
func (c *FakeCustomers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(customersResource, customersKind, c.ns, opts), &v1alpha1.CustomerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CustomerList{ListMeta: obj.(*v1alpha1.CustomerList).ListMeta}
	for _, item := range obj.(*v1alpha1.CustomerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customers.
func (c *FakeCustomers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(customersResource, c.ns, opts))

}

// Create takes the representation of a customer and creates it.  Returns the server's representation of the customer, and an error, if there is any.
func (c *FakeCustomers) Create(ctx context.Context, customer *v1alpha1.Customer, opts v1.CreateOptions) (result *v1alpha1.Customer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(customersResource, c.ns, customer), &v1alpha1.Customer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Customer), err
}

// Update takes the representation of a customer and updates it. Returns the server's representation of the customer, and an error, if there is any.
func (c *FakeCustomers) Update(ctx context.Context, customer *v1alpha1.Customer, opts v1.UpdateOptions) (result *v1alpha1.Customer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(customersResource, c.ns, customer), &v1alpha1.Customer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Customer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomers) UpdateStatus(ctx context.Context, customer *v1alpha1.Customer, opts v1.UpdateOptions) (*v1alpha1.Customer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(customersResource, "status", c.ns, customer), &v1alpha1.Customer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Customer), err
}

// Delete takes name of the customer and deletes it. Returns an error if one occurs.
func (c *FakeCustomers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(customersResource, c.ns, name), &v1alpha1.Customer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(customersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CustomerList{})
	return err
}

// Patch applies the patch and returns the patched customer.
func (c *FakeCustomers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Customer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(customersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Customer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Customer), err
}
