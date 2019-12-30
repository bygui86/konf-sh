package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func main() {

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	config := clientcmd.GetConfigFromFileOrDie(*kubeconfig)
	fmt.Println(*config)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return "~/"
}
