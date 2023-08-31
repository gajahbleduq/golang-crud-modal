package mahasiswamodel

import (
	"database/sql"

	"gitlab.com/tomimulhartono/go-crud-modal/config"
	"gitlab.com/tomimulhartono/go-crud-modal/entities"
)

type MahasiswaModel struct {
	db *sql.DB
}

func New() *MahasiswaModel {
	db, err := config.DBConnection()

	if err != nil {
		panic(err)
	}

	return &MahasiswaModel{db: db}
}

func (m *MahasiswaModel) FindAll(mahasiswa *[]entities.Mahasiswa) error {
	rows, err := m.db.Query("select * from mahasiswa")

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Mahasiswa
		rows.Scan(
			&data.Id,
			&data.FullName,
			&data.Gender,
			&data.Birthplace,
			&data.Birthdate,
			&data.Address)

		*mahasiswa = append(*mahasiswa, data)
	}

	return nil
}

func (m *MahasiswaModel) Create(mahasiswa *entities.Mahasiswa) error {
	result, err := m.db.Exec("insert into mahasiswa (fullname, gender, birthdate, birthplace, address) values (?,?,?,?,?)",
		mahasiswa.FullName, mahasiswa.Gender, mahasiswa.Birthplace, mahasiswa.Birthdate, mahasiswa.Address)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	mahasiswa.Id = lastInsertId
	return nil
}

func (m *MahasiswaModel) Find(id int64, mahasiswa *entities.Mahasiswa) error {
	return m.db.QueryRow("select * from mahasiswa where id = ?", id).Scan(
		&mahasiswa.Id,
		&mahasiswa.FullName,
		&mahasiswa.Gender,
		&mahasiswa.Birthplace,
		&mahasiswa.Birthdate,
		&mahasiswa.Address)
}

func (m *MahasiswaModel) Update(id int64, mahasiswa *entities.Mahasiswa) error {
	_, err := m.db.Exec("update mahasiswa set fullname = ?, gender = ?, birthplace = ?, birthdate = ?, address = ? where id = ?",
		mahasiswa.FullName, mahasiswa.Gender, mahasiswa.Birthplace, mahasiswa.Birthdate, mahasiswa.Address, mahasiswa.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *MahasiswaModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from mahasiswa where id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
