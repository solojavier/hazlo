package persistence

import (
	"os"
	"strconv"
	"time"

	"github.com/solojavier/hazlo/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func getStepCollection() (s *mgo.Session, c *mgo.Collection) {
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		panic(err)
	}

	return session, session.DB("heroku_app25841211").C("step")
}

func CreateStep(user string, goal int, progress int) (id string) {
	date := time.Now()
	_, week := date.ISOWeek()

	step := models.Step{user, date, week, date.Year(), goal, progress}
	selector := bson.M{"week": step.Week, "user": step.User}

	s, c := getStepCollection()
	defer s.Close()

	_, err := c.Upsert(selector, &step)
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(step.Year) + "/" + strconv.Itoa(step.Week) + "/" + step.User
}

func QuerySteps(year int, week int) []models.Step {
	s, c := getStepCollection()
	result := []models.Step{}

	defer s.Close()
	c.Find(bson.M{"year": year, "week": week}).Sort("user").All(&result)

	return result
}
