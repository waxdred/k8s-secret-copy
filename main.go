package main

import (
	"fmt"

	config "github.com/waxdred/k8s-secret-copy/Config"
	k8s_tools "github.com/waxdred/k8s-secret-copy/k8s_tool"
)

func main() {
	cfg := config.NewConfig()
	cfg.AutoComplete()
	k, err := k8s_tools.NewK8s(cfg.Namespace, cfg.Kubeconfig)
	if err != nil {
		fmt.Println("Error creating K8s client: ", err)
		return
	}
	err = k.GetSecret(cfg.Key, cfg.SecretName)
	if err != nil {
		fmt.Println("Error getting secret: ", err)
		return
	}

}
