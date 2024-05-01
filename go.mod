module wranglertest

go 1.22.1

replace github.com/rancher/wrangler/v2 => github.com/tomleb/rancher-wrangler/v2 v2.2.0-rc1.0.20240501122407-54d673bf211a

require (
	github.com/rancher/norman v0.0.0-20240417185323-cf0f9cc85249
	github.com/rancher/rancher/pkg/apis v0.0.0-20240501135617-dbe409f2e7a2
	github.com/rancher/wrangler/v2 v2.2.0-rc6
	k8s.io/api v0.30.0
	k8s.io/apimachinery v0.30.0
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/code-generator v0.30.0 // indirect
	k8s.io/gengo v0.0.0-20240310015720-9cff6334dab4 // indirect
	k8s.io/gengo/v2 v2.0.0-20240228010128-51d4e06bde70 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)
