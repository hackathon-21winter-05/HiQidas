package repository

type Repository interface {
	Transaction
	UserRepository
	HeyaRepository
	HiqidashiRepository
	HistoryRepository
	TsunaRepository
	FavoriteRepository
}

type RepositoryEx interface {
	Repository
	CredentialRepository
}
