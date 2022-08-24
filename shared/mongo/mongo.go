package mgutil

import (
	"coolcar/shared/mongo/objid"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	IDFieldName       = "_id"
	UpdateAtFieldName = "updateat"
)

type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UpdateAtField struct {
	UpdateAt int64 `bson:"updateat"`
}

var NewObjID = primitive.NewObjectID()

func NewObjectIDWithValue(id fmt.Stringer) {
	NewObjID = func() primitive.ObjectID {
		return objid.MustFromID(id)
	}
}

var UpdateAt = func() int64 {
	return time.Now().UnixNano()
}

// Set 返回一个$set document
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert 返回一个$setOnInsert document
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

// ZeroOrDoesNotExist 生成字段为0，或该字段不存在的表达式
func ZeroOrDoesNotExist(field string, zero interface{}) bson.M {
	return bson.M{
		"$or": []bson.M{
			{
				field: zero,
			},
			{
				field: bson.M{
					"$exists": false,
				},
			},
		},
	}
}
