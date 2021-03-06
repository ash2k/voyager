// Generated code
// run `make generate` to update

// Code generated by client-gen. DO NOT EDIT.

package client

import (
	trebuchetv1 "github.com/atlassian/voyager/pkg/trebuchet/client/typed/trebuchet/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	TrebuchetV1() trebuchetv1.TrebuchetV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Trebuchet() trebuchetv1.TrebuchetV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	trebuchetV1 *trebuchetv1.TrebuchetV1Client
}

// TrebuchetV1 retrieves the TrebuchetV1Client
func (c *Clientset) TrebuchetV1() trebuchetv1.TrebuchetV1Interface {
	return c.trebuchetV1
}

// Deprecated: Trebuchet retrieves the default version of TrebuchetClient.
// Please explicitly pick a version.
func (c *Clientset) Trebuchet() trebuchetv1.TrebuchetV1Interface {
	return c.trebuchetV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.trebuchetV1, err = trebuchetv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.trebuchetV1 = trebuchetv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.trebuchetV1 = trebuchetv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
