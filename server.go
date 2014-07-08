package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	"github.com/solojavier/hazlo/mailer"
	"github.com/solojavier/hazlo/models"
	"github.com/solojavier/hazlo/persistence"
)

type updateForm struct {
	Goal        int    `form:"goal"`
	Progress    int    `form:"progress"`
	User        string `form:"user"`
	Measurement string `form:"measurement"`
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/reports/:user", func(params martini.Params, r render.Render) {
		r.HTML(200, "report_form", params["user"])
	})

	m.Get("/reports/:year/:week", func(params martini.Params, r render.Render) {
		year := ptoi(params["year"], r)
		week := ptoi(params["week"], r)
		reports := persistence.QueryReports(year, week)

		r.HTML(200, "reports", reports)
	})

	m.Post("/reports", binding.Bind(updateForm{}), func(form updateForm, params martini.Params, res http.ResponseWriter) int {
		date := time.Now()
		_, week := date.ISOWeek()
		report := models.Report{form.User, date, week, date.Year(), form.Goal, form.Progress, form.Measurement}
		id := persistence.CreateReport(report)

		res.Header().Set("Location", "reports/"+id)

		return 201
	})

	m.Get("/emails/weekly", func(params martini.Params, r render.Render) string {
		mailer.SendWeeklyReport()

		return "OK"
	})

	m.Run()
}

func ptoi(param string, r render.Render) (param_value int) {
	param_value, err := strconv.Atoi(param)

	if err != nil {
		r.Error(422)
	}

	return param_value
}
