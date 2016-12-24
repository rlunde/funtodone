package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SaveTask - save a task to a mongodb collection
func SaveTask(t *Task, c *mgo.Collection) error {
	err := c.Insert(t)
	return err
}

//FindTaskByID - read a task from mongodb by its ID
func FindTaskByID(c *mgo.Collection, id bson.ObjectId) (*Task, error) {
	result := Task{}
	err := c.Find(bson.M{"_id": id}).One(&result)
	return &result, err
}
