package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mogo"
	"coolcar/shared/mogo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripFieId      = "trip"
	accountIdField = tripFieId + ".accountid"
)

type Mongo struct {
	col      *mongo.Collection
	newObjId func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("trip"),
		newObjId: primitive.NewObjectID,
	}
}

type TripRecord struct {
	mgo.ObjId          `bson: "inline"`
	mgo.UpdatedAtField `bson: "inline"`
	Trip               *rentalpb.Trip `bson: "trip"`
}




func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*rentalpb.Trip, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = m.newObjId()
	r.UpdatedAt = mgo.UpdatedAt()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r.Trip, nil
}

func (m *Mongo) GetTrip(c context.Context, id id.TripId, accountId id.AccountId) (*TripRecord, error) {
	objId, err := objid.FromId(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	res := m.col.FindOne(c, bson.M{
		mgo.IdField:    objId,
		accountIdField: accountId,
	})

	if err := res.Err(); err != nil {
		return nil, err
	}
	var tr TripRecord
	err = res.Decode(&tr)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &tr, nil
}
