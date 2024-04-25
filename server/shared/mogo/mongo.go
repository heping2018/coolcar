package mgo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const IdField = "_id"
const UpdateAtFiledName = "updatedat"

func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat"`
}

type ObjId struct {
	ID primitive.ObjectID `bson:"_id"`
}
