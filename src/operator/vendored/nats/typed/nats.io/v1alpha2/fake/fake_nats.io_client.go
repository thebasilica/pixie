// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha2 "px.dev/pixie/src/operator/vendored/nats/typed/nats.io/v1alpha2"
)

type FakeNatsV1alpha2 struct {
	*testing.Fake
}

func (c *FakeNatsV1alpha2) NatsClusters(namespace string) v1alpha2.NatsClusterInterface {
	return &FakeNatsClusters{c, namespace}
}

func (c *FakeNatsV1alpha2) NatsServiceRoles(namespace string) v1alpha2.NatsServiceRoleInterface {
	return &FakeNatsServiceRoles{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNatsV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}