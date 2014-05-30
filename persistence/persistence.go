package persistence

import (
	"os"
	"time"

	"github.com/solojavier/make/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func getStepCollection() (s *mgo.Session, c *mgo.Collection) {
	session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if err != nil {
		panic(err)
	}

	return session, session.DB("make").C("step")
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
	defer s.Close()

	result := models.Step{}
	err := c.Find(bson.M{"user": user}).Sort("-date").One(&result)

	if err != nil {
		panic(err)
	}

	return result
}
