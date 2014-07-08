package mailer

import (
	"bytes"
	"html/template"
	"os"
	"time"

	"github.com/mostafah/mandrill"
	"github.com/solojavier/hazlo/persistence"
)

func SendWeeklyReport() {
	mandrill.Key = os.Getenv("MANDRILL_KEY")
	msg := mandrill.NewMessageTo(os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_NAME"))
	msg.HTML = reportBody()
	msg.Subject = "Reporte Semanal"
	msg.FromEmail = "no-reply@hazlo.herokuapp.com"
	msg.FromName = "Hazlo"

	_, err := msg.Send(false)

	if err != nil {
		panic(err)
	}
}

func reportBody() string {
	var doc bytes.Buffer
	year := time.Now().Year()
	_, week := time.Now().ISOWeek()
	reports := persistence.QueryReports(year, week)

	t, _ := template.New("t").ParseFiles("templates/report_email.tmpl")
	err := t.ExecuteTemplate(&doc, "report", reports)

	if err != nil {
		panic(err)
	}

	return doc.String()
}
