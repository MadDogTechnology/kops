/*
Copyright 2016 The Kubernetes Authors.

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

package cloudup

import (
	"k8s.io/kops/pkg/apis/kops"
	"testing"
)

func TestPopulateClusterSpec_Proxy(t *testing.T) {
	c := buildMinimalCluster()

	c.Spec.EgressProxy = &kops.EgressProxySpec{
		ProxyExcludes: "google.com",
		HTTPProxy: kops.HTTPProxy{
			Host: "52.205.179.249",
			Port: 3128,
		},
	}

	c.Spec.NonMasqueradeCIDR = "100.64.0.1/10"
	var err error
	c.Spec.EgressProxy, err = assignProxy(c)
	if err != nil {
		t.Fatalf("unable to assign proxy, %v", err)
	}

	if c.Spec.EgressProxy.ProxyExcludes != "google.com,127.0.0.1,localhost,testcluster.test.com,100.64.0.2,100.64.0.1/10,169.254.169.254" {
		t.Fatalf("Incorrect proxy excludes set: %v", c.Spec.EgressProxy.ProxyExcludes)
	}

	c.Spec.EgressProxy = &kops.EgressProxySpec{
		HTTPProxy: kops.HTTPProxy{
			Host: "52.205.179.249",
			Port: 3128,
		},
	}

	c.Spec.NonMasqueradeCIDR = "100.64.0.0/10"
	c.Spec.EgressProxy.ProxyExcludes = ""

	c.Spec.EgressProxy, err = assignProxy(c)
	if err != nil {
		t.Fatalf("unable to assign proxy, %v", err)
	}

	if c.Spec.EgressProxy.ProxyExcludes != "127.0.0.1,localhost,testcluster.test.com,100.64.0.1,100.64.0.0/10,169.254.169.254" {
		t.Fatalf("Incorrect proxy excludes set: %v", c.Spec.EgressProxy.ProxyExcludes)
	}

	c.Spec.NonMasqueradeCIDR = "172.16.0.5/12"
	c.Spec.CloudProvider = "gce"
	c.Spec.EgressProxy.ProxyExcludes = ""
	c.Spec.EgressProxy, err = assignProxy(c)
	if err != nil {
		t.Fatalf("unable to assign proxy, %v", err)
	}

	if c.Spec.EgressProxy.ProxyExcludes != "127.0.0.1,localhost,testcluster.test.com,172.16.0.6,172.16.0.5/12" {
		t.Fatalf("Incorrect proxy excludes set: %v", c.Spec.EgressProxy.ProxyExcludes)
	}

}
