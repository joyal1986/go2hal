package database

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/CardFrontendDevopsTeam/GoMongo"
)

type auditEntry struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Datatype string
	Payload string
}

//SaveAudit adds a audit record to the mongo DB
func SaveAudit(dataType,payload string){
	r := auditEntry{Datatype:dataType,Payload:payload}
	c := database.Mongo.C("audit")
	c.Insert(r)
}
