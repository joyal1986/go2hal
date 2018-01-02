package database

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/CardFrontendDevopsTeam/GoMongo"
)

/*
User Json object
 */
type User struct {
	EmployeeNumber string `json:"employeeNumber"`
	CallOutName    string `json:"calloutName"`
	JIRAName       string `json:"jiraName"`
}

/*
AddUser alows for a new user to be added to the database
 */
func AddUser(employeeNumber, CalloutName, JiraName string) {
	c := database.Mongo.C("Users")
	u := User{CallOutName: CalloutName, EmployeeNumber: employeeNumber, JIRAName: JiraName}
	c.Insert(u)
}

/*
FindUserByCalloutName Return a user whos details matches the callout
 */
func FindUserByCalloutName(name string) User {
	var r User
	c := database.Mongo.C("Users")
	c.Find(bson.M{"calloutname": name}).One(&r)
	return r
}