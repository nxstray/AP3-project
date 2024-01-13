package main

import (
	"net/http"

	"github.com/nxstray/AP3-project/controllers/mahasiswacontroller"
)

func main() {
	http.HandleFunc("/", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/index", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/add", mahasiswacontroller.Add)
	http.HandleFunc("/mahasiswa/edit", mahasiswacontroller.Edit)
	http.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
