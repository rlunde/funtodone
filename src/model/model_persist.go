package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SaveTask(t *Task, c *mgo.Collection) error {
	err := c.Insert(t)
	return err
}

func FindTaskByID(c *mgo.Collection, id bson.ObjectId) (*Task, error) {
	result := Task{}
	err := c.Find(bson.M{"_id": id}).One(&result)
	return &result, err
}
