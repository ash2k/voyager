// Generated code
// run `make generate` to update

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	creatorv1 "github.com/atlassian/voyager/pkg/apis/creator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServices implements ServiceInterface
type FakeServices struct {
	Fake *FakeCreatorV1
}

var servicesResource = schema.GroupVersionResource{Group: "creator", Version: "v1", Resource: "services"}

var servicesKind = schema.GroupVersionKind{Group: "creator", Version: "v1", Kind: "Service"}

// Get takes name of the service, and returns the corresponding service object, and an error if there is any.
func (c *FakeServices) Get(name string, options v1.GetOptions) (result *creatorv1.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(servicesResource, name), &creatorv1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*creatorv1.Service), err
}

// List takes label and field selectors, and returns the list of Services that match those selectors.
func (c *FakeServices) List(opts v1.ListOptions) (result *creatorv1.ServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(servicesResource, servicesKind, opts), &creatorv1.ServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &creatorv1.ServiceList{ListMeta: obj.(*creatorv1.ServiceList).ListMeta}
	for _, item := range obj.(*creatorv1.ServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested services.
func (c *FakeServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(servicesResource, opts))
}

// Create takes the representation of a service and creates it.  Returns the server's representation of the service, and an error, if there is any.
func (c *FakeServices) Create(service *creatorv1.Service) (result *creatorv1.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(servicesResource, service), &creatorv1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*creatorv1.Service), err
}

// Update takes the representation of a service and updates it. Returns the server's representation of the service, and an error, if there is any.
func (c *FakeServices) Update(service *creatorv1.Service) (result *creatorv1.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(servicesResource, service), &creatorv1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*creatorv1.Service), err
}

// Delete takes name of the service and deletes it. Returns an error if one occurs.
func (c *FakeServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(servicesResource, name), &creatorv1.Service{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(servicesResource, listOptions)

	_, err := c.Fake.Invokes(action, &creatorv1.ServiceList{})
	return err
}

// Patch applies the patch and returns the patched service.
func (c *FakeServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *creatorv1.Service, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicesResource, name, pt, data, subresources...), &creatorv1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*creatorv1.Service), err
}
