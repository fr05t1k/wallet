package operation

import (
	"github.com/fr05t1k/wallet/wallet"
	"gopkg.in/mgo.v2"
)

type Manager struct {
	DB         *mgo.Database
	collection string
}

func ( manager *Manager)AddOperation(operation *wallet.Operation) (error) {
	return manager.DB.C("Operation").Insert(operation)
}

func (manager *Manager)GetTotal() (*wallet.Total, error) {
	total := wallet.Total{Value:0}

	return &total, nil
}
