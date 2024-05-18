package service

type Service struct {
	db *DB
}

func New(db *DB) *Service {
	return &Service{db: db}
}
