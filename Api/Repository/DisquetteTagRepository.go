package repository

import entity "go-api-test.kayn.ooo/Api/Entity"

type DisquetteTagRepositoryInterface struct {
	GenericRepository
}

func (ur *DisquetteTagRepositoryInterface) Init() {
	DB.RegisterModel(&entity.DisquetteTag{})
	_, err := DB.NewCreateTable().Model(&entity.DisquetteTag{}).IfNotExists().Exec(Ctx)
	if err != nil {
		panic(err)
	}

	err = DB.ResetModel(Ctx, &entity.DisquetteTag{})
	if err != nil {
		panic(err)
	}
}
