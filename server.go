package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/solojavier/make/persistence"
)

type UpdateForm struct {
	Goal   int `form:"goal"`
	Status int `form:"status"`
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/user/:id/report", func(params martini.Params, r render.Render) {
		r.HTML(200, "report_form", params["id"])
	})

	m.Get("/user/:id", func(params martini.Params, r render.Render) {
		step := persistence.LastStep(params["id"])
		r.HTML(200, "user", step)
	})

	m.Post("/user/:id/step", binding.Bind(UpdateForm{}), func(form UpdateForm, params martini.Params, res http.ResponseWriter) int {
		id := persistence.CreateStep(params["id"], form.Goal, form.Status)

		res.Header().Set("Location", "step/"+id)

		return 201
	})

	m.Run()
}
