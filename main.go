package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	type Mahasiswa struct {
		Nim int `json:"id"`
		Nama string `json:"nama"`
		Prodi string `json:"prodi"`
		Fakultas string `json:"fakultas"`
	}

	mahasiswa :=[]Mahasiswa{
		{Nim: 251242, Nama: "Adib", Prodi: "Teknik Elektro", Fakultas: "Teknik"},
		{Nim: 255683, Nama: "Rizky", Prodi: "Teknik Elektro", Fakultas: "Teknik"},
	}

	http.HandleFunc("GET /" , func(w http.ResponseWriter , r *http.Request) {
		var input Mahasiswa
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(mahasiswa)
	})

	http.ListenAndServe(":3000", nil)
}