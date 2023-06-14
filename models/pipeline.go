package models

import "github.com/beego/beego/v2/client/orm"

type Pipeline struct {
	Id          int
	ProjectId   string
	ProjectName string
	ImageName   string
	AppName     string
	PackagePath string
	Package     string
}

func init() {

	orm.RegisterModel(new(Pipeline))
}

func CreatePipeline(pipeline *Pipeline) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(pipeline)
	return id, err
}
