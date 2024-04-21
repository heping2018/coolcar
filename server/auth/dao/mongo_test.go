package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestResoulveAccountId(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf(" %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Fatalf(" %v", err)
	} else {
		want := "6624bd8c0f743a30531effb1"
		if id != want {
			t.Errorf("resolve account id: want: %q ,got: %q", want, id)
		}
	}

}
