// Generated code
// run `make generate` to update

// Code generated by client-gen. DO NOT EDIT.

package client

import (
	reporterv1 "github.com/atlassian/voyager/pkg/reporter/client/typed/reporter/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ReporterV1() reporterv1.ReporterV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Reporter() reporterv1.ReporterV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	reporterV1 *reporterv1.ReporterV1Client
}

// ReporterV1 retrieves the ReporterV1Client
func (c *Clientset) ReporterV1() reporterv1.ReporterV1Interface {
	return c.reporterV1
}

// Deprecated: Reporter retrieves the default version of ReporterClient.
// Please explicitly pick a version.
func (c *Clientset) Reporter() reporterv1.ReporterV1Interface {
	return c.reporterV1
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
	cs.reporterV1, err = reporterv1.NewForConfig(&configShallowCopy)
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
	cs.reporterV1 = reporterv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.reporterV1 = reporterv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
