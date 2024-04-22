package dao

import (
	"context"
	mgo "coolcar/shared/mogo"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIdField = "open_id"

type Mongo struct {
	col      *mongo.Collection
	newObjId func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("account"),
		newObjId: primitive.NewObjectID,
	}
}

func (m *Mongo) ResolveAccountID(c context.Context, openId string) (string, error) {
	// m.col.InsertOne(c, bson.M{
	// 	mgo.IdField: m.newObjId(),
	// 	openIdField: openId,
	// })
	insertedId := m.newObjId()
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIdField: openId,
	}, mgo.SetOnInsert(bson.M{
		mgo.IdField: insertedId,
		openIdField: openId,
	}), options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate, %v", err)
	}
	var row mgo.ObjId
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate, %v", err)
	}
	return row.ID.Hex(), nil
}
