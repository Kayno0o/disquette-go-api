package repository

import entity "go-api-test.kayn.ooo/Api/Entity"

type DisquetteRepositoryInterface struct {
	GenericRepository
}

func (ur *DisquetteRepositoryInterface) Init() {
	DB.RegisterModel(&entity.Disquette{})
	_, err := DB.NewCreateTable().Model(&entity.Disquette{}).IfNotExists().Exec(Ctx)
	if err != nil {
		panic(err)
	}

	err = DB.ResetModel(Ctx, &entity.Disquette{})
	if err != nil {
		panic(err)
	}
}
