package main

import (
	controllergen "github.com/rancher/wrangler/v2/pkg/controller-gen"
	"github.com/rancher/wrangler/v2/pkg/controller-gen/args"
	appsv1 "k8s.io/api/apps/v1"
)

func main() {
	controllergen.Run(args.Options{
		ImportPackage: "github.com/tomleb/wrangler_test/pkg/generated",
		OutputPackage: "github.com/tomleb/wrangler_test/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			appsv1.GroupName: {
				Types: []interface{}{
					appsv1.Deployment{},
					appsv1.DaemonSet{},
					appsv1.StatefulSet{},
				},
				// Known to not work before
				// GenerateTypes: true,
				// GenerateClients:   true,
				// GenerateInformers: true,
				// GenerateListers:   true,
			},
			"cluster.cattle.io": {
				PackageName: "cluster.cattle.io",
				Types: []interface{}{
					// All structs with an embedded ObjectMeta field will be picked up
					"./pkg/apis/cluster.cattle.io/v3",
				},
				GenerateTypes:     true,
				GenerateClients:   true,
				GenerateInformers: true,
				GenerateListers:   true,
			},
		},
	})
}
