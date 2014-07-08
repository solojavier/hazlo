package persistence

import (
	"os"
	"strconv"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func CreateReport(report Report) (id string) {
	selector := bson.M{"year": report.Year, "week": report.Week, "user": report.User}
	s, c := getReportCollection()
	defer s.Close()

	_, err := c.Upsert(selector, &report)
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(report.Year) + "/" + strconv.Itoa(report.Week) + "/" + report.User
}

func QueryReports(year int, week int) []Report {
	result := []Report{}
	s, c := getReportCollection()

	defer s.Close()
	c.Find(bson.M{"year": year, "week": week}).Sort("user").All(&result)

	return result
}

func getReportCollection() (s *mgo.Session, c *mgo.Collection) {
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		panic(err)
	}

	return session, session.DB("heroku_app25841211").C("report")
}
