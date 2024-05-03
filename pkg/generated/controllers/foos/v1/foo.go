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

// Code generated by wrangler_test. DO NOT EDIT.

package v1

import (
	"github.com/rancher/wrangler/v3/pkg/generic"
	v1 "github.com/tomleb/wrangler_test/pkg/apis/foos/v1"
)

// FooController interface for managing Foo resources.
type FooController interface {
	generic.ControllerInterface[*v1.Foo, *v1.FooList]
}

// FooClient interface for managing Foo resources in Kubernetes.
type FooClient interface {
	generic.ClientInterface[*v1.Foo, *v1.FooList]
}

// FooCache interface for retrieving Foo resources in memory.
type FooCache interface {
	generic.CacheInterface[*v1.Foo]
}
