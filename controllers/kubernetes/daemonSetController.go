package kubernetes

import (
	"Ayile/models"
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSetController struct {
	beego.Controller
}

type DaemonSetData struct {
	Name           string
	Labels         map[string]string
	Annotations    map[string]string
	Function       string
	DeployEnv      string
	NodeGroup      string
	ServiceGroup   string
	Replicas       string
	ContainerPorts []int32
	Namespace      string
}

type DaemonSetViewData struct {
	SelectedNamespace string
	Namespaces        []models.NamespaceData
	DaemonSets        []DaemonSetData
}

func (c *DaemonSetController) QueryDaemonSets() {
	selectedNamespace := c.GetString("namespace")
	if selectedNamespace == "" {
		selectedNamespace = "All"
	}

	client, err := models.NewKubernetesClient()
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	namespaces, err := client.ListNamespaces()
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	var daemonSets []DaemonSetData
	var totalNumberReady int32
	var totalNumberAvailable int32
	var totalNumberUnavailable int32

	for _, ns := range namespaces {
		if selectedNamespace == "All" || ns.Name == selectedNamespace {
			daemonSetList, err := client.GetClientset().AppsV1().DaemonSets(ns.Name).List(context.TODO(), metav1.ListOptions{})

			if err != nil {
				c.Data["json"] = map[string]interface{}{"error": err.Error()}
				return
			}

			for _, daemon := range daemonSetList.Items {

				totalNumberReady = totalNumberReady + daemon.Status.NumberReady
				totalNumberAvailable = totalNumberAvailable + daemon.Status.NumberAvailable
				totalNumberUnavailable = totalNumberUnavailable + daemon.Status.NumberUnavailable
				var containerPorts []int32
				for _, container := range daemon.Spec.Template.Spec.Containers {
					for _, port := range container.Ports {
						containerPorts = append(containerPorts, port.ContainerPort)
					}
				}

				replicas := fmt.Sprintf("%d/%d", daemon.Status.NumberReady, daemon.Status.NumberAvailable)

				daemonSet := DaemonSetData{
					Name:           daemon.Name,
					Labels:         daemon.Labels,
					ContainerPorts: containerPorts,
					Replicas:       replicas,
					Annotations:    daemon.ObjectMeta.Annotations,
					Function:       daemon.Spec.Template.ObjectMeta.Annotations["Function"],
					DeployEnv:      daemon.Spec.Template.ObjectMeta.Annotations["deployEnv"],
					NodeGroup:      daemon.Spec.Template.ObjectMeta.Annotations["nodeGroup"],
					ServiceGroup:   daemon.Spec.Template.ObjectMeta.Annotations["serviceGroup"],
					Namespace:      daemon.Namespace,
				}
				daemonSets = append(daemonSets, daemonSet)
			}
		}
	}

	daemonSetViewData := DaemonSetViewData{
		SelectedNamespace: selectedNamespace,
		Namespaces:        namespaces,
		DaemonSets:        daemonSets,
	}

	TotalDaemonSet := daemonSetViewData.DaemonSets
	c.Data["TotalDaemonSet"] = len(TotalDaemonSet)
	c.Data["DaemonSetViewData"] = daemonSetViewData
	c.Data["User"] = c.GetSession("user")
	c.TplName = "kubernetes/daemonsetQuery.html"
	currentPath := c.Ctx.Request.URL.Path
	c.Data["Page"] = currentPath

}
