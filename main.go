package main

import (
	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"
)

func main() {
	controllergen.Run(args.Options{
		ImportPackage: "github.com/tomleb/wrangler_test/pkg/generated",
		OutputPackage: "github.com/tomleb/wrangler_test/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"foos": {
				PackageName: "foos",
				Types: []interface{}{
					// All structs with an embedded ObjectMeta field will be picked up
					"./pkg/apis/foos/v1",
				},
				GenerateTypes:   true,
				GenerateClients: true,
				// GenerateInformers: true,
				// GenerateListers:   true,
			},
		},
	})
}
