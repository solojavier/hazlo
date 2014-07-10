package persistence

import (
	"os"
	"strconv"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func CreateReport(user string, goal int, progress int, measurement string) (id string) {
	date := time.Now()
	_, week := date.ISOWeek()
	fulfillment := (float64(progress) / float64(goal)) * 100

	report := Report{user, date, week, date.Year(), goal, progress, measurement, fulfillment}

	createReportRecord(report)

	return strconv.Itoa(report.Year) + "/" + strconv.Itoa(report.Week) + "/" + report.User
}

func QueryReports(year int, week int) []Report {
	result := []Report{}
	s, c := getReportCollection()

	c.Find(bson.M{"year": year, "week": week}).Sort("user").All(&result)
	s.Close()

	return result
}

func createReportRecord(report Report) {
	selector := bson.M{"year": report.Year, "week": report.Week, "user": report.User}
	s, c := getReportCollection()

	_, err := c.Upsert(selector, &report)
	s.Close()

	if err != nil {
		panic(err)
	}
}

func getReportCollection() (s *mgo.Session, c *mgo.Collection) {
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		panic(err)
	}

	return session, session.DB("heroku_app25841211").C("report")
}
