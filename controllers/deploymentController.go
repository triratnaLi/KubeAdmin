package controllers

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type DeploymentController struct {
	beego.Controller
}

type NameSpaceData struct {
	Name              string
	Labels            map[string]string
	CreationTimestamp metav1.Time
}

type DeploymentData struct {
	Name           string
	Labels         map[string]string
	Annotations    map[string]string
	Function       string
	DeployEnv      string
	NodeGroup      string
	ServiceGroup   string
	GitName        string
	Replicas       string
	ContainerPorts []int32
	Status         string
}

type ViewData struct {
	SelectedNamespace    string
	Namespaces           []NameSpaceData
	Deployments          []DeploymentData
	DeploymentLabelsFunc func(string) map[string]string
}

func (c *DeploymentController) Get() {
	selectedNamespace := c.GetString("namespace")

	config, err := clientcmd.BuildConfigFromFlags("", "conf/k8s.conf")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	var namespaceData []NameSpaceData
	var deployments []DeploymentData

	for _, ns := range namespaces.Items {

		data := NameSpaceData{
			Name:              ns.ObjectMeta.Name,
			Labels:            ns.ObjectMeta.Labels,
			CreationTimestamp: ns.ObjectMeta.CreationTimestamp,
		}
		namespaceData = append(namespaceData, data)

		if selectedNamespace != "" && ns.ObjectMeta.Name == selectedNamespace {
			deploys, err := clientset.AppsV1().Deployments(selectedNamespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				c.Data["json"] = map[string]interface{}{"error": err.Error()}
				return
			}

			for _, deploy := range deploys.Items {

				var status string
				for _, condition := range deploy.Status.Conditions {
					if condition.Type == "Available" && condition.Status == "True" {
						status = "Running"
						break
					} else if condition.Status != "True" {
						status = "Not Running"
					}
				}

				var containerPorts []int32
				for _, container := range deploy.Spec.Template.Spec.Containers {
					for _, port := range container.Ports {
						containerPorts = append(containerPorts, port.ContainerPort)
					}
				}
				replicas := fmt.Sprintf("%d/%d", deploy.Status.AvailableReplicas, *deploy.Spec.Replicas)

				deploymentData := DeploymentData{
					Name:           deploy.Name,
					Labels:         deploy.ObjectMeta.Labels,
					Annotations:    deploy.ObjectMeta.Annotations,
					Function:       deploy.Spec.Template.ObjectMeta.Annotations["Function"],
					DeployEnv:      deploy.Spec.Template.ObjectMeta.Annotations["deployEnv"],
					NodeGroup:      deploy.Spec.Template.ObjectMeta.Annotations["nodeGroup"],
					ServiceGroup:   deploy.Spec.Template.ObjectMeta.Annotations["serviceGroup"],
					GitName:        deploy.Spec.Template.ObjectMeta.Annotations["GitName"],
					Replicas:       replicas,
					ContainerPorts: containerPorts,
					Status:         status,
				}
				deployments = append(deployments, deploymentData)

			}
		}
	}

	viewData := ViewData{
		SelectedNamespace:    selectedNamespace,
		Namespaces:           namespaceData,
		Deployments:          deployments,
		DeploymentLabelsFunc: GetDeploymentLabels,
	}

	currentPath := c.Ctx.Request.URL.Path
	c.Data["ViewData"] = viewData
	c.TplName = "kubernetes/DeploymentQuery.html"
	c.Data["Page"] = currentPath
}

func GetDeploymentLabels(deploymentName string) map[string]string {
	// 在这里根据 deploymentName 获取相应的标签
	// 返回一个 map[string]string 类型的标签数据

	return map[string]string{}
}
