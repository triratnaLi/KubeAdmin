package kubernetes

import (
	"fmt"

	"Ayile/services" // 导入你的services包
	beego "github.com/beego/beego/v2/server/web"
)

type ServicesController struct {
	beego.Controller
}

func (c *ServicesController) QuerySvc() {
	SelectedNamespace := c.GetString("namespace")
	if SelectedNamespace == "" {
		SelectedNamespace = "All"
	}

	fmt.Println(SelectedNamespace)

	serviceViewData, err := services.QueryServices(SelectedNamespace)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
		return
	}

	c.TplName = "kubernetes/ServiceQuery.html"
	c.Data["ServiceViewData"] = serviceViewData

	currentPath := c.Ctx.Request.URL.Path
	c.Data["User"] = c.GetSession("user")
	c.Data["Page"] = currentPath

	fmt.Println(currentPath)
	fmt.Println(serviceViewData.Namespaces[0].Name)
}
