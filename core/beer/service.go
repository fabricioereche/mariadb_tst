package beer

import (
	"fmt"
)

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) GetAll() ([]*Beer, error) {

	rows, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	fmt.Println("rows: ", len(rows))

	return rows, nil

}

func (s *Service) Get(ID int64) (*Beer, error) {

	beer, err := s.repository.Get(ID)

	fmt.Println("Return from get:", beer)

	return beer, err

}

func (s *Service) Store(b *Beer) error {

	fmt.Println("Tenando gravar uma cerveja ", b.Type.String(), " daquele estilo ", b.Style.String(), " no capricho xD")

	err := s.repository.Store(b)

	return err

}

func (s *Service) Update(b *Beer) error {
	return s.repository.Update(b)
}

func (s *Service) Remove(ID int64) error {
	return s.repository.Remove(ID)
}
