package k8s_tools

import (
	"context"
	"fmt"

	"github.com/atotto/clipboard"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8s struct {
	clientset  *kubernetes.Clientset
	kubeconfig string
	namespace  string
}

func NewK8s(namespace, kubeconfig string) (*K8s, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	return &K8s{
		clientset:  clientset,
		kubeconfig: kubeconfig,
		namespace:  namespace,
	}, nil
}

func (k *K8s) GetSecret(key, secretName string) error {
	// Get secret name and key
	secret, err := k.clientset.CoreV1().Secrets(k.namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if key != "" {
		if value, ok := secret.Data[key]; ok {
			err := clipboard.WriteAll(string(string(value)))
			if err != nil {
				return err
			}
			fmt.Println("Secret copied to clipboard")
			return nil
		} else {
			return fmt.Errorf("key %s not found in secret %s", key, secretName)
		}
	} else {
		if len(secret.Data) == 1 {
			for _, value := range secret.Data {
				err = clipboard.WriteAll(string(value))
				if err != nil {
					return err
				}
				fmt.Println("Secret copied to clipboard")
				break
			}
		} else {
			fmt.Printf("The secret %s contains multiple keys.\n", secretName)
			for key, value := range secret.Data {
				fmt.Printf("  - %s: %s\n", key, string(value))
			}
			return nil
		}
	}

	return nil
}
