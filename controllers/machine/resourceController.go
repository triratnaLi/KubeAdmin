package machine

import (
	"Ayile/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type ResourceController struct {
	beego.Controller
}

func (this *ResourceController) Query() {
	o := orm.NewOrm()
	machines := []models.Machine{}
	_, err := o.QueryTable("machine").All(&machines)
	if err != nil {
		this.Abort("500")
	}

	currentPath := this.Ctx.Request.URL.Path

	totalCpu := 0
	for _, machine := range machines {
		totalCpu += machine.CPU
	}

	totalMem := 0
	for _, machine := range machines {
		totalMem += machine.Memory
	}

	totalSysDisk := 0
	for _, machine := range machines {
		totalSysDisk += machine.SystemDisk
	}

	totalDataDisk := 0
	for _, machine := range machines {
		totalDataDisk += machine.DataDisk
	}

	this.Data["Machines"] = machines
	this.Data["TotalCpu"] = totalCpu
	this.Data["TotalMem"] = totalMem
	this.Data["TotalSysDisk"] = totalSysDisk
	this.Data["TotalDataDisk"] = totalDataDisk
	this.Data["User"] = this.GetSession("user")
	this.Data["Page"] = currentPath

	this.TplName = "resource/query.html"
}

func (this *ResourceController) Add() {
	machine := models.Machine{}
	if err := this.ParseForm(&machine); err != nil {
		this.Abort("500")
	}
	o := orm.NewOrm()
	_, err := o.Insert(&machine)
	if err != nil {
		this.Abort("500")
	}
	this.Redirect("/resource/query", 302)
}

func (this *ResourceController) Edit() {
	o := orm.NewOrm()
	id, _ := this.GetInt("id")
	machine := models.Machine{Id: id}
	if err := o.Read(&machine); err != nil {
		this.Abort("500")
	}
	this.Data["Machine"] = machine
	this.TplName = "resource/resource.html"
}

func (this *ResourceController) Update() {
	machine := models.Machine{}
	if err := this.ParseForm(&machine); err != nil {
		this.Abort("500")
	}
	o := orm.NewOrm()
	_, err := o.Update(&machine)
	if err != nil {
		this.Abort("500")
	}
	this.Redirect("resource/resource.html", 302)
}

func (this *ResourceController) Delete() {
	o := orm.NewOrm()
	id, _ := this.GetInt("id")
	machine := models.Machine{Id: id}
	if _, err := o.Delete(&machine); err != nil {
		this.Abort("500")
	}
	this.Redirect("resource/resource.html", 302)
}
