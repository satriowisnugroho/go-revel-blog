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

func (c App) Words() revel.Result {
	a := c.Params.Route.Get("a")
	b := c.Params.Route.Get("b")

	res := make(map[string]interface{})
	res["code"] = 200
	res["status"] = true
	res["message"] = "Success"
	res["data"] = a + " " + b

	return c.RenderJSON(res)
}
