package kubernetes

import (
	"Ayile/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
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
	Status         string
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
	for _, ns := range namespaces {
		if selectedNamespace == "All" || ns.Name == selectedNamespace {
			// TODO: Retrieve daemon sets for the namespace and populate daemonSets variable
			// Example: daemonSets = append(daemonSets, daemonsetData)
		}
	}

	daemonSetViewData := DaemonSetViewData{
		SelectedNamespace: selectedNamespace,
		Namespaces:        namespaces,
		DaemonSets:        daemonSets,
	}

	c.Data["DaemonSetViewData"] = daemonSetViewData
	c.Data["User"] = c.GetSession("user")
	c.TplName = "kubernetes/daemonsetQuery.html"
	fmt.Println(daemonSetViewData.Namespaces[1].Name)
}
