package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/solojavier/it/persistence"
)

type UpdateForm struct {
	Goal   int `form:"goal"`
	Status int `form:"status"`
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/users/:id/report", func(params martini.Params, r render.Render) {
		r.HTML(200, "report_form", params["id"])
	})

	m.Get("/users/:id", func(params martini.Params, r render.Render) {
		step := persistence.LastStep(params["id"])
		r.HTML(200, "user", step)
	})

	m.Post("/users/:id/step", binding.Bind(UpdateForm{}), func(form UpdateForm, params martini.Params, res http.ResponseWriter) int {
		id := persistence.CreateStep(params["id"], form.Goal, form.Status)

		res.Header().Set("Location", "step/"+id)

		return 201
	})

	m.Run()
}
