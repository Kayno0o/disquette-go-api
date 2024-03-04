package repository

var (
	UserRepository = &UserRepositoryStruct{}
)

type UserRepositoryStruct struct {
	GenericRepository
}
