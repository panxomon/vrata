package mocks

import (
	entity "vrata/internal/domain/entity"
	"vrata/internal/railway"
)

var UserOK *entity.User = &entity.User{
	Id:   1,
	Name: "panxo",
	Age:  40,
}

var UserFail *entity.User = &entity.User{
	Id:   -1,
	Name: "panxo",
	Age:  40,
}

var SuccessOk *railway.Success = &railway.Success{
	Value: true,
}
