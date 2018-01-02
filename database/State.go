package database

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/CardFrontendDevopsTeam/GoMongo"
)

/*
State is the current user state when flowing through a multi step process.
 */
type State struct {
	Userid int
	State string
	Field []string
}

/*
SetState sets the users current state
 */
func SetState(user int, state string, field []string) error{
	s := State{Userid:user,State:state,Field:field}
	c := database.Mongo.C("userstate")
	c.RemoveAll(bson.M{"userid":user})
	return c.Insert(&s)
}

/*
GetState returns the users current state
 */
func GetState(user int) State{
	c := database.Mongo.C("userstate")
	var s State
	c.Find(bson.M{"userid":user}).One(&s)
	return s
}


