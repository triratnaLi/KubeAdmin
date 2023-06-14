package gitlab

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"

	"Ayile/models"
)

type PipelineController struct {
	beego.Controller
}

func (c *PipelineController) ShowForm() {
	c.TplName = "gitlab/pipeline_form.tpl"
}

func (c *PipelineController) CreatePipeline() {
	pipeline := models.Pipeline{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pipeline)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	id, err := models.CreatePipeline(&pipeline)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = id
	c.ServeJSON()
}
