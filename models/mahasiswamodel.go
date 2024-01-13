package models

import (
	"database/sql"
	"fmt"

	"github.com/nxstray/AP3-project/config"
	"github.com/nxstray/AP3-project/entities"
)

type MahasiswaModel struct {
	conn *sql.DB
}

func NewMahasiswaModel() *MahasiswaModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &MahasiswaModel{
		conn: conn,
	}
}

func (p *MahasiswaModel) FindAll() ([]entities.Mahasiswa, error) {
	rows, err := p.conn.Query("select * from datamahasiswa")
	if err != nil {
		return []entities.Mahasiswa{}, err
	}
	defer rows.Close()

	var dataMahasiswa []entities.Mahasiswa
	for rows.Next() {
		var mahasiswa entities.Mahasiswa
		rows.Scan(&mahasiswa.Id,
			&mahasiswa.Npm,
			&mahasiswa.NamaLengkap,
			&mahasiswa.Fakultas,
			&mahasiswa.Kelas)

		dataMahasiswa = append(dataMahasiswa, mahasiswa)
	}

	return dataMahasiswa, nil
}

func (p *MahasiswaModel) Create(mahasiswa entities.Mahasiswa) bool {

	result, err := p.conn.Exec("insert into datamahasiswa (npm, nama_lengkap, fakultas, kelas) values(?,?,?,?)", mahasiswa.Npm, mahasiswa.NamaLengkap, mahasiswa.Fakultas, mahasiswa.Kelas)

	if err != nil {
		fmt.Println(err)
		return false
	}
	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *MahasiswaModel) Find(id int64, mahasiswa *entities.Mahasiswa) error {
	return p.conn.QueryRow("select * from datamahasiswa where id = ?", id).Scan(&mahasiswa.Id,
		&mahasiswa.Npm,
		&mahasiswa.NamaLengkap,
		&mahasiswa.Fakultas,
		&mahasiswa.Kelas)
}

func (p *MahasiswaModel) Update(mahasiswa entities.Mahasiswa) error {

	_, err := p.conn.Exec(
		"Update datamahasiswa set npm = ?, nama_lengkap = ?, fakultas = ?, kelas = ? where id = ?",
		mahasiswa.Npm, mahasiswa.NamaLengkap, mahasiswa.Fakultas, mahasiswa.Kelas, mahasiswa.Id)

	if err != nil {
		return err
	}

	return nil

}

func (p *MahasiswaModel) Delete(id int64) {
	p.conn.Exec("delete from datamahasiswa where id = ?", id)
}
