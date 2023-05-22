package models

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sService struct {
}

// deployment
type k8sDeployment struct {
}

func (this *k8sService) GetClient() (*kubernetes.Clientset, error) {
	path := beego.AppConfig.DefaultString("k8s::path", "conf/k8s.conf")
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)

}

type deploymentService struct {
	k8sService
}

func (this *deploymentService) Query() []appsV1.Deployment {
	clientset, err := this.GetClient()
	if err != nil {
		return []appsV1.Deployment{}
	}

	deploymentList, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return []appsV1.Deployment{}
	}
	return deploymentList.Items
}

func (this *deploymentService) Delete(name string, namespace string) {
	clientset, err := this.GetClient()
	if err != nil {
		return
	}
	clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

//namespace

type NameSpace struct {
	k8sService
}

func (this *NameSpace) Query() []coreV1.Namespace {
	clientset, err := this.GetClient()

	if err != nil {
		return []coreV1.Namespace{}
	}
	nameSpaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return []coreV1.Namespace{}
	}
	return nameSpaceList.Items
}

var DeploymentService = new(deploymentService)
var NamespaceService = new(NameSpace)
