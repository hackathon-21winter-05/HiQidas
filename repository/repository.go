package repository

type Repository interface {
	Transaction
	UserRepository
}
