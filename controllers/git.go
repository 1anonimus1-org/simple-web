package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// ServerController operations for Server
type GitController struct {
	beego.Controller
}

type ServerStatus struct {
	AppName string `json:"app_name"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (c *GitController) GitWebHook() {
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &inputArray); err == nil {
	// } else {
	// 	c.Data["json"] = Response{
	// 		Success: 0,
	// 		Error:   "2 = " + err.Error(),
	// 	}
	// 	status = "Gagal"
	// 	keterangan = "2 = " + err.Error()
	// }
	fmt.Println("\n===== LOG NOTIF ME =====")
	fmt.Println("Request Message :\n", string(c.Ctx.Input.RequestBody))
}
