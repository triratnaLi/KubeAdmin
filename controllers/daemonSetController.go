package controllers

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type DaemonSetController struct {
	beego.Controller
}

type DSNameSpaceData struct {
	Name              string
	Labels            map[string]string
	CreationTimestamp metav1.Time
}

type DaemonSetData struct {
	Name         string
	Labels       map[string]string
	Annotations  map[string]string
	Function     string
	DeployEnv    string
	NodeGroup    string
	ServiceGroup string
	//GitName        string
	Replicas       string
	ContainerPorts []int32
	Status         string
	NameSpace      string
}

type DSViewData struct {
	SelectedNamespace   string
	Namespaces          []DSNameSpaceData
	DaemonSets          []DaemonSetData
	DaemonSetLabelsFunc func(string) map[string]string
}

func (this *DaemonSetController) Get() {

	SelectedNamespace := this.GetString("namespace")
	if SelectedNamespace == "" {
		SelectedNamespace = "All"
	}

	fmt.Println(SelectedNamespace)

	config, err := clientcmd.BuildConfigFromFlags("", "conf/k8s.conf")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		this.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	var dsNameSpaceData []DSNameSpaceData
	var daemonsets []DaemonSetData

	var totalNumberReady int32
	var totalNumberAvailable int32
	var totalNumberUnavailable int32
	for _, ns := range namespaces.Items {

		data := DSNameSpaceData{
			Name:              ns.ObjectMeta.Name,
			Labels:            ns.ObjectMeta.Labels,
			CreationTimestamp: ns.ObjectMeta.CreationTimestamp,
		}

		dsNameSpaceData = append(dsNameSpaceData, data)

		if SelectedNamespace == "All" || ns.ObjectMeta.Name == SelectedNamespace {
			daemons, err := clientset.AppsV1().DaemonSets(ns.ObjectMeta.Name).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				this.Data["json"] = map[string]interface{}{"error": err.Error()}
				return
			}

			for _, daemon := range daemons.Items {

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
				daemonsetData := DaemonSetData{
					Name:           daemon.Name,
					Labels:         daemon.Labels,
					ContainerPorts: containerPorts,
					Replicas:       replicas,
					Annotations:    daemon.ObjectMeta.Annotations,
					Function:       daemon.Spec.Template.ObjectMeta.Annotations["Function"],
					DeployEnv:      daemon.Spec.Template.ObjectMeta.Annotations["deployEnv"],
					NodeGroup:      daemon.Spec.Template.ObjectMeta.Annotations["nodeGroup"],
					ServiceGroup:   daemon.Spec.Template.ObjectMeta.Annotations["serviceGroup"],
					NameSpace:      daemon.Namespace,
				}
				daemonsets = append(daemonsets, daemonsetData)
			}

		}
	}

	dsViewData := DSViewData{
		SelectedNamespace: SelectedNamespace,
		Namespaces:        dsNameSpaceData,
		DaemonSets:        daemonsets,
	}
	TotalDaemonSet := dsViewData.DaemonSets

	currentPath := this.Ctx.Request.URL.Path
	this.Data["Page"] = currentPath

	this.Data["DSViewData"] = dsViewData
	this.Data["User"] = this.GetSession("user")
	this.Data["TotalDaemonSet"] = len(TotalDaemonSet)
	this.Data["TotalNumberReady"] = totalNumberReady
	this.Data["TotalNumberAvailable"] = totalNumberAvailable
	this.Data["TotalNumberUnavailable"] = totalNumberUnavailable

	this.TplName = "kubernetes/daemonsetQuery.html"

}
