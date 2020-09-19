package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// ServerController operations for Server
type ServerController struct {
	beego.Controller
}

type ServerStatus struct {
	AppName string `json:"app_name"`
	Status  string `json:"status"`
	Version string `json:"version"`
	IP      string `json:"ip"`
	Port    int    `json:"port"`
}

func (c *ServerController) Status() {

	ip := c.Ctx.Input.IP()
	port := c.Ctx.Input.Port()

	var serverStatus = ServerStatus{
		AppName: beego.AppConfig.String("appname"),
		Status:  "Active",
		Version: "1.0",
		IP:      ip,
		Port:    port,
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = serverStatus

	c.ServeJSON()
	fmt.Println("IP   : ", ip)
	fmt.Println("PORT : ", port)
	return
}
