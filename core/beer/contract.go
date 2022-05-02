package beer

import "database/sql"

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int64) error
}

type Contract struct {
	DB *sql.DB
}

func NewContract(db *sql.DB) *Contract {
	return &Contract{
		DB: db,
	}
}
