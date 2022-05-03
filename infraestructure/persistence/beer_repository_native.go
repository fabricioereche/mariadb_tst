package persistence

import (
	"database/sql"
	"fmt"

	"github.com/fabricioereche/mariadb_tst/core/beer"
)

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		DB: db,
	}
}

func (s *repository) GetAll() ([]*beer.Beer, error) {

	var result []*beer.Beer

	rows, err := s.DB.Query("select id, name, type, style from beer")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var b beer.Beer
		err = rows.Scan(&b.ID, &b.Name, &b.Type, &b.Style)
		if err != nil {
			return nil, err
		}

		result = append(result, &b)
	}
	return result, nil
}

func (s *repository) Get(ID int64) (*beer.Beer, error) {

	var b beer.Beer

	stmt, err := s.DB.Prepare("select id, name, type, style from beer where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&b.ID, &b.Name, &b.Type, &b.Style)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (s *repository) Store(b *beer.Beer) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into beer(id, name, type, style) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(b.ID, b.Name, b.Type, b.Style)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (s *repository) Update(b *beer.Beer) error {
	if b.ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update beer set name=?, type=?, style=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(b.Name, b.Type, b.Style, b.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (s *repository) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beer where id=?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
