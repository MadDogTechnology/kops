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
	"testing"
	"k8s.io/kops/pkg/apis/kops"
)

func TestPopulateClusterSpec_Proxy(t *testing.T) {
	c := buildMinimalCluster()

	c.Spec.EgressProxy = &kops.EgressProxySpec{
		ProxyExcludes: "google.com",
		HTTPProxy: kops.HTTPProxySpec{
			Host: "52.205.179.249",
			Port: 3128,
		},
	}

	c.Spec.EgressProxy = assignProxy(c)
	c.Spec.NonMasqueradeCIDR = "100.64.0.1/10"

	if c.Spec.EgressProxy.ProxyExcludes != "google.com,127.0.0.1,localhost,169.254.169.254,testcluster.test.com,100.64.0.1" {
		t.Fatalf("Incorrect proxy excludes set: %v", c.Spec.EgressProxy.ProxyExcludes)
	}

	c.Spec.EgressProxy = &kops.EgressProxySpec{
		HTTPProxy: kops.HTTPProxySpec{
			Host: "52.205.179.249",
			Port: 3128,
		},
	}

	c.Spec.EgressProxy = assignProxy(c)

	if c.Spec.EgressProxy.ProxyExcludes != "127.0.0.1,localhost,169.254.169.254,testcluster.test.com,100.64.0.1" {
		t.Fatalf("Incorrect proxy excludes set: %v", c.Spec.EgressProxy.ProxyExcludes)
	}

}
