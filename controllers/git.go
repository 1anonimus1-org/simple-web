package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// GitController operations for Server
type GitController struct {
	beego.Controller
}

type Response struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c *GitController) GitWebHook() {
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &inputArray); err == nil {
	c.Data["json"] = Response{
		Success: 1,
		Message: "Sukses",
	}
	// } else {
	// c.Data["json"] = Response{
	// 	Success: 0,
	// 	Message:   "2 = " + err.Error(),
	// }
	// 	status = "Gagal"
	// 	keterangan = "2 = " + err.Error()
	// }
	c.ServeJSON()

	fmt.Println("\n===== LOG NOTIF ME =====")
	fmt.Println("Request Message :\n", string(c.Ctx.Input.RequestBody))
}
