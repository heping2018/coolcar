package dao

import (
	"context"
	mongotesting "coolcar/shared/mogo/testing"

	"testing"

	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI = ""

func TestResoulveAccountId(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf(" %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	m.newObjId = func() primitive.ObjectID {
		ObjectId, _ := primitive.ObjectIDFromHex("6624bd8c0f743a30531effb1")
		return ObjectId
	}
	cases := []struct {
		name   string
		openId string
		want   string
	}{
		{
			name:   "existing_user",
			openId: "openid_1",
			want:   "123123",
		},
		{
			name:   "another_existing_user",
			openId: "openid_2",
			want:   "123123",
		},
		{
			name:   "new_user",
			want:   "6624bd8c0f743a30531effb1",
			openId: "openid_3",
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openId)
			if err != nil {
				t.Errorf("faidl resolve %v", err)
			}
			if id != cc.want {
				t.Errorf("faidl resolve ")
			}

		})
	}
	// id, err := m.ResolveAccountID(c, m.newObjId().Hex())
	// if err != nil {
	// 	t.Fatalf(" %v", err)
	// } else {
	// 	want := "6624bd8c0f743a30531effb1"
	// 	if id != want {
	// 		t.Errorf("resolve account id: want: %q ,got: %q", want, id)
	// 	}
	// }

}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
