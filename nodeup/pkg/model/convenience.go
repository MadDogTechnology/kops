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

package model

import (
	"os"

	"github.com/golang/glog"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/upup/pkg/fi"
)

// s is a helper that builds a *string from a string value
func s(v string) *string {
	return fi.String(v)
}

// i64 is a helper that builds a *int64 from an int64 value
func i64(v int64) *int64 {
	return fi.Int64(v)
}

func getProxyEnvVars(proxy *kops.EgressProxySpec) []v1.EnvVar {

	if proxy == nil {
		return nil
	}

	httpProxy := os.Getenv("http_proxy")
	noProxy := os.Getenv("NO_PROXY")

	if httpProxy == "" || noProxy == "" {
		glog.Warning("http_proxy or NO_PROXY environment variable is empty")
		glog.Warning("http_proxy=%q", httpProxy)
		glog.Warning("NO_PROXY=%q", noProxy)
	}
	return []v1.EnvVar{
		{Name: "http_proxy", Value: httpProxy},
		{Name: "https_proxy", Value: httpProxy},
		{Name: "NO_PROXY", Value: noProxy},
		{Name: "no_proxy", Value: noProxy},
	}

}
