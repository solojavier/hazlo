package main

import (
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/solojavier/hazlo/persistence"
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

	m.Get("/steps/:year/:week", func(params martini.Params, r render.Render) {
		year, err := strconv.Atoi(params["year"])

		if err != nil {
			r.Error(422)
		}

		week, err := strconv.Atoi(params["week"])

		if err != nil {
			r.Error(422)
		}

		steps := persistence.QuerySteps(year, week)

		r.HTML(200, "steps", steps)
	})

	m.Post("/users/:id/step", binding.Bind(UpdateForm{}), func(form UpdateForm, params martini.Params, res http.ResponseWriter) int {
		id := persistence.CreateStep(params["id"], form.Goal, form.Status)

		res.Header().Set("Location", "steps/"+id)

		return 201
	})

	m.Run()
}
