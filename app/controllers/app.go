package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello() revel.Result {
	data := make(map[string]interface{})
	data["code"] = 200
	data["status"] = true
	data["message"] = "Hello World!"
	data["data"] = make(map[string]interface{})

	return c.RenderJSON(data)
}
