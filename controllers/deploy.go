package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"os"
	"os/exec"
)

type deployController struct {
	beego.Controller
}

func (this *deployController) Get() {
	this.TplName = "deploy/deploy.html"
}

func (this *deployController) Post() {
	file, header, err := this.GetFile("file")

	if err != nil {
		this.Ctx.WriteString("文件上传失败" + err.Error())
		return
	}

	defer file.Close()

	//保存上传的文件到服务器临时目录
	dstPath := "tmp" + header.Filename
	dst, err := os.Create(dstPath)

	if err != nil {
		this.Ctx.WriteString("文件保存失败：" + err.Error())
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		this.Ctx.WriteString("文件保存失败" + err.Error())
		return
	}

	// 构建 Docker 镜像
	imageName := "your_image_name:tag"
	buildCmd := exec.Command("docker", "build", "-t", imageName, ".")
	buildCmd.Dir = "/path/to/your/app" // 替换为实际的应用目录
	buildOutput, err := buildCmd.CombinedOutput()
	if err != nil {
		this.Ctx.WriteString("Docker 镜像构建失败：" + err.Error())
		return
	}

	// 部署到 Kubernetes
	deployCmd := exec.Command("kubectl", "apply", "-f", "/path/to/your/deployment.yaml") // 替换为实际的 deployment.yaml 文件路径
	deployOutput, err := deployCmd.CombinedOutput()
	if err != nil {
		this.Ctx.WriteString("Kubernetes 部署失败：" + err.Error())
		return
	}

	this.Ctx.WriteString(fmt.Sprintf("文件上传成功！Docker 镜像构建成功！Kubernetes 部署成功！\n\nDocker 镜像构建输出：%s\n\nKubernetes 部署输出：%s", buildOutput, deployOutput))
}
