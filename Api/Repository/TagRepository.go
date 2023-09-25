package repository

import entity "go-api-test.kayn.ooo/Api/Entity"

type TagRepositoryInterface struct {
	GenericRepository
}

func (ur *TagRepositoryInterface) Init() {
	DB.RegisterModel(&entity.Tag{})
	_, err := DB.NewCreateTable().Model(&entity.Tag{}).IfNotExists().Exec(Ctx)
	if err != nil {
		panic(err)
	}

	err = DB.ResetModel(Ctx, &entity.Tag{})
	if err != nil {
		panic(err)
	}
}
