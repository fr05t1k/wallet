package operation

import (
	"github.com/fr05t1k/wallet/wallet"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const collectionName = "Operation"

type Manager struct {
	DB *mgo.Database
}

func ( manager *Manager)AddOperation(operation *wallet.Operation) (error) {
	return manager.DB.C(collectionName).Insert(operation)
}

func (manager *Manager)GetTotal() (*wallet.Total, error) {
	pipe := manager.DB.C(collectionName).Pipe(
		[]bson.M{{
			"$group" : bson.M{
				"_id" : bson.NewObjectId(),
				"total" : bson.M{"$sum": "$value"},
			},
		}},
	)
	resp := bson.M{}
	err := pipe.One(&resp)
	total := wallet.Total{}

	if err != nil {
		log.Printf("Can't get total from db %v", err)
	} else {
		total.Value = resp["total"].(int64)
	}

	return &total, nil
}
