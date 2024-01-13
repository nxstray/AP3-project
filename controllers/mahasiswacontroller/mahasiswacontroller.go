package mahasiswacontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/nxstray/AP3-project/libraries"

	"github.com/nxstray/AP3-project/models"

	"github.com/nxstray/AP3-project/entities"
)

var validation = libraries.NewValidation()
var mahasiswaModel = models.NewMahasiswaModel()

func Index(response http.ResponseWriter, request *http.Request) {

	mahasiswa, _ := mahasiswaModel.FindAll()

	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}

	temp, err := template.ParseFiles("views/mahasiswa/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/mahasiswa/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var mahasiswa entities.Mahasiswa
		mahasiswa.Npm = request.Form.Get("npm")
		mahasiswa.NamaLengkap = request.Form.Get("nama_lengkap")
		mahasiswa.Fakultas = request.Form.Get("fakultas")
		mahasiswa.Kelas = request.Form.Get("kelas")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(mahasiswa)

		if vErrors != nil {
			data["mahasiswa"] = mahasiswa
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data mahasiswa berhasil disimpan"
			mahasiswaModel.Create(mahasiswa)
		}

		temp, _ := template.ParseFiles("views/mahasiswa/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var mahasiswa entities.Mahasiswa
		mahasiswaModel.Find(id, &mahasiswa)

		data := map[string]interface{}{
			"mahasiswa": mahasiswa,
		}

		temp, err := template.ParseFiles("views/mahasiswa/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var mahasiswa entities.Mahasiswa
		mahasiswa.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		mahasiswa.Npm = request.Form.Get("npm")
		mahasiswa.NamaLengkap = request.Form.Get("nama_lengkap")
		mahasiswa.Fakultas = request.Form.Get("fakultas")
		mahasiswa.Kelas = request.Form.Get("kelas")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(mahasiswa)

		if vErrors != nil {
			data["mahasiswa"] = mahasiswa
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data mahasiswa berhasil diperbarui"
			mahasiswaModel.Update(mahasiswa)
		}

		temp, _ := template.ParseFiles("views/mahasiswa/edit.html")
		temp.Execute(response, data)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	mahasiswaModel.Delete(id)

	http.Redirect(response, request, "/mahasiswa", http.StatusSeeOther)

}
