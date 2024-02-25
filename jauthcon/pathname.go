package jauthcon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Menu struct {
	Kategori  string `json:"Kategori"`
	Nama      string `json:"Nama"`
	Harga     int    `json:"Harga"`
	Deskripsi string `json:"Deskripsi"`
	Gambar    []byte `json:"Gambar"`
}

type Toko struct {
	Menu    []Menu   `json:"Menu"`
	Pemilik []string `json:"Pemilik"`
}

func Showall(sila Toko) (map[string]interface{}, error) {
	Data := map[string]interface{}{}
	var err error
	for i, v := range sila.Menu {
		isinya := map[string]interface{}{
			"Deskripsi": v.Deskripsi,
			"Kategori":  v.Kategori,
			"Nama":      v.Nama,
			"Harga":     v.Harga,
			"Gambar":    v.Gambar,
			"Kaki":      sila.Pemilik[i],
		}

		if err != nil {
			return Data, err
		}
		Data[v.Nama] = isinya
	}
	return Data, nil
}

func Bacadata(namaFile string) (Toko, error) {
	file, err := os.Open(namaFile)
	if err != nil {
		return Toko{}, nil
	}
	de := json.NewDecoder(file)
	Data := Toko{}
	de.Decode(&Data)
	return Data, nil
}

func Homepagecan(w http.ResponseWriter, r *http.Request) {
	Data, err := Bacadata("Data.json")
	if err != nil {
		fmt.Println(err)
	}

	temp, _ := template.ParseFiles("tempat/index.html")
	semuamenu, x := Showall(Data)
	if x != nil {
		fmt.Println(x)
	}
	temp.ExecuteTemplate(w, "index.html", semuamenu)
	// temp.Execute(w,subtoko)
	// fmt.Println("Data================", Data)
	// fmt.Println("SUBTOKO==============", subtoko)

}
