package repository

type Repository interface {
	Transaction
	UserRepository
	HeyaRepository
	HiqidashiRepository
	HistoryRepository
	TsunaRepository
}
