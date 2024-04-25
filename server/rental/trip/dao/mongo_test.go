package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mongotesting "coolcar/shared/mogo/testing"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf(" %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: "account1",
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "startPoint",
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName:  "endPoint",
			FeeCent:  10000,
			KmDriven: 35,
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 115,
			},
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Error(err)
	}
	t.Error(tr)

	got, err := m.GetTrip(c, id.TripId(tr.CarId), "account1")
	if err != nil {
		t.Errorf("cannot get trip : %v", err)
	}
	t.Error(got)
}

var mongoURI = ""

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
