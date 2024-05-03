package examples

import (
	"context"
	"os"

	"github.com/rancher/wrangler/v3/pkg/kubeconfig"
	foosv1 "github.com/tomleb/wrangler_test/pkg/apis/foos/v1"
	"github.com/tomleb/wrangler_test/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ExampleFoo() {
	ctx := context.Background()
	clientConfig, _ := kubeconfig.GetNonInteractiveClientConfig(os.Getenv("KUBECONFIG")).ClientConfig()
	clientset, _ := versioned.NewForConfig(clientConfig)
	foo := &foosv1.Foo{
		Spec: foosv1.FooSpec{
			Bar: "toto",
		},
	}

	clientset.FoosV1().Foos("").Create(ctx, foo, metav1.CreateOptions{})

	foo.Spec.Bar = "baz"
	clientset.FoosV1().Foos("").Update(ctx, foo, metav1.UpdateOptions{})

	clientset.FoosV1().Foos("").Delete(ctx, foo.GetName(), metav1.DeleteOptions{})
}
