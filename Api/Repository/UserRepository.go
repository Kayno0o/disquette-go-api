package repository

import entity "go-api-test.kayn.ooo/Api/Entity"

type UserRepositoryInterface struct {
	GenericRepository
}

func (ur *UserRepositoryInterface) Init() {
	DB.RegisterModel(&entity.User{})
	_, err := DB.NewCreateTable().Model(&entity.User{}).IfNotExists().Exec(Ctx)
	if err != nil {
		panic(err)
	}

	err = DB.ResetModel(Ctx, &entity.User{})
	if err != nil {
		panic(err)
	}
}
