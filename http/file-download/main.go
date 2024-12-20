package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/vladimirvivien/gexe"
)

// The following code download the kubectl client command-line binary
func main() {
	if err := os.Mkdir("kube", 0745); err != nil {
		fmt.Printf("kube dir: %s\n", err)
		os.Exit(1)
	}

	exe := gexe.SetVar("KUBE_RELEASE_URL", "https://dl.k8s.io/release").
		SetVar("OS", runtime.GOOS).
		SetVar("ARCH", runtime.GOARCH).
		SetVar("kubebin", "./kube/kubectl")

	// the following is equivalent to:
	// curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
	get := exe.Get(
		"${KUBE_RELEASE_URL}/",
		exe.Get("${KUBE_RELEASE_URL}/stable.txt").String(),
		"/bin/${OS}/${ARCH}/kubectl",
	)

	if err := exe.FileWrite("$kubebin").WithMode(0766).From(get.Body()).Err(); err != nil {
		fmt.Printf("kubectl install: %s", err)
	}
	fmt.Printf("kubectl saved: %s\n\n", exe.Eval("$kubebin"))

	// print version
	fmt.Println(gexe.Run("$kubebin version --client=true"), "\n")
}
