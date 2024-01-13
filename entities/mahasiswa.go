package entities

type Mahasiswa struct {
	Id          int64
	Npm         string `validate:"required" label:"NPM"`
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	Fakultas    string `validate:"required"`
	Kelas       string `validate:"required"`
}
