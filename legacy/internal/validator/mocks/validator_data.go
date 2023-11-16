package mocks

import (
	entity "vrata/legacy/internal/domain/entity"
	"vrata/legacy/internal/railway"
)

var UserOK *entity.User = &entity.User{
	Id:   1,
	Name: "panxo",
	Age:  40,
}

var UserFail *entity.User = &entity.User{
	Id:   -1,
	Name: "juan",
	Age:  40,
}

var SuccessOk *railway.Result = &railway.Result{
	Value: true,
}
