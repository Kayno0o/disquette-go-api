package repository

import entity "go-api-test.kayn.ooo/Api/Entity"

type VoteRepositoryInterface struct {
	GenericRepository
}

func (ur *VoteRepositoryInterface) Init() {
	DB.RegisterModel(&entity.Vote{})
	_, err := DB.NewCreateTable().Model(&entity.Vote{}).IfNotExists().Exec(Ctx)
	if err != nil {
		panic(err)
	}

	err = DB.ResetModel(Ctx, &entity.Vote{})
	if err != nil {
		panic(err)
	}
}
