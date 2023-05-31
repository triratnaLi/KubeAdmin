package kubernetes

import (
	"Ayile/models"
	"context"
	beego "github.com/beego/beego/v2/server/web"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ServicesController struct {
	beego.Controller
}

type ServiceData struct {
	Name                  string
	Status                string
	Port                  []int32
	TargetPort            []int32
	NodePort              []int32
	Namespace             string
	Type                  v1.ServiceType
	SessionAffinity       string
	InternalTrafficPolicy string
	Selector              string
}

type ServiceViewData struct {
	SelectedNamespace string
	Namespaces        []models.NamespaceData
	Services          []ServiceData
}

func (c *ServicesController) QueryServices() {
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

	var services []ServiceData

	for _, ns := range namespaces {
		if selectedNamespace == "All" || ns.Name == selectedNamespace {
			svcList, err := client.GetClientset().CoreV1().Services(ns.Name).List(context.TODO(), metav1.ListOptions{})

			if err != nil {
				c.Data["json"] = map[string]interface{}{"error": err.Error()}
				return
			}

			for _, svc := range svcList.Items {
				var port []int32
				var targetport []int32
				var nodeport []int32
				for _, ports := range svc.Spec.Ports {

					port = append(port, ports.Port)
					targetport = append(targetport, int32(ports.TargetPort.IntValue()))
					nodeport = append(nodeport, ports.NodePort)
				}

				selector := labels.Set(svc.Spec.Selector).String()
				serviceData := ServiceData{
					Name:                  svc.Name,
					Namespace:             svc.Namespace,
					Type:                  svc.Spec.Type,
					SessionAffinity:       string(svc.Spec.SessionAffinity),
					InternalTrafficPolicy: string(*svc.Spec.InternalTrafficPolicy),
					Selector:              selector,
					Port:                  port,
					NodePort:              nodeport,
					TargetPort:            targetport,
				}
				services = append(services, serviceData)
			}
		}
	}

	serviceViewData := ServiceViewData{
		SelectedNamespace: selectedNamespace,
		Namespaces:        namespaces,
		Services:          services,
	}

	c.Data["ServiceViewData"] = serviceViewData
	c.Data["User"] = c.GetSession("user")
	c.TplName = "kubernetes/ServiceQuery.html"
	currentPath := c.Ctx.Request.URL.Path
	c.Data["Page"] = currentPath

}
