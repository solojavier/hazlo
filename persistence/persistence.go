package persistence

import (
	"os"
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

	step := models.Step{bson.NewObjectId(), user, date, week, date.Year(), goal, progress}

	s, c := getStepCollection()
	defer s.Close()

	err := c.Insert(&step)
	if err != nil {
		panic(err)
	}

	return step.Id.Hex()
}

func LastStep(user string) models.Step {
	s, c := getStepCollection()
	_, week := time.Now().ISOWeek()
	result := models.Step{}

	defer s.Close()
	c.Find(bson.M{"user": user, "week": week - 1}).Sort("-date").One(&result)

	return result
}
