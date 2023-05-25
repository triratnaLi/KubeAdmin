package models

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient struct {
	clientset *kubernetes.Clientset
}

type NamespaceData struct {
	Name              string
	Labels            map[string]string
	CreationTimestamp metav1.Time
}

func NewKubernetesClient() (*KubernetesClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "conf/k8s.conf")
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubernetesClient{
		clientset: clientset,
	}, nil
}

func (kc *KubernetesClient) ListNamespaces() ([]NamespaceData, error) {
	namespaces, err := kc.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var namespaceData []NamespaceData
	for _, ns := range namespaces.Items {
		data := NamespaceData{
			Name:              ns.ObjectMeta.Name,
			Labels:            ns.ObjectMeta.Labels,
			CreationTimestamp: ns.ObjectMeta.CreationTimestamp,
		}
		namespaceData = append(namespaceData, data)
	}

	return namespaceData, nil
}

func (kc *KubernetesClient) GetClientset() *kubernetes.Clientset {
	return kc.clientset
}
