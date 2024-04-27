package models

import (
	"net/http"

	"github.com/RifkyAkhsanul/kbs-be/db"
)

type Info struct {
	Id          int     `json:"id"`
	NamaSpesies string  `json:"namaspesies"`
	BahasaLatin string  `json:"bahasa_latin"`
	KoordinatX  float64 `json:"koordinatx"`
	KoordinatY  float64 `json:"koordinaty"`
	Status      string  `json:"status"`
	FunFact     string  `json:"funfact"`
	Program     string  `json:"program"`
	Dampak      string  `json:"dampak"`
	PetaHabitat string  `json:"peta_habitat"`
	Dikunjungi  bool    `json:"dikunjungi"`
	Gambar      string  `json:"gambar"`
	GambarKecil string  `json:"gambar_kecil"`
}

func GetInfo() (Response, error) {
	var obj Info
	var arrobj []Info
	var res Response

	cons, err := db.NewDriver()
	if err != nil {
		return res, err
	}

	sqlstatement := "SELECT * FROM info"

	rows, err := cons.Query(sqlstatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaSpesies, &obj.BahasaLatin, &obj.KoordinatX, &obj.KoordinatY, &obj.Status, &obj.FunFact, &obj.Program, &obj.Dampak, &obj.PetaHabitat, &obj.Dikunjungi, &obj.Gambar, &obj.GambarKecil)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func GetInfofromID(id int) (Response, error) {
	var obj Info
	var arrobj []Info
	var res Response

	con, err := db.NewDriver()

	sqlStatament := "SELECT * FROM info WHERE id=?"

	stmt, err := con.Prepare(sqlStatament)
	if err != nil {
		return res, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaSpesies, &obj.BahasaLatin, &obj.KoordinatX, &obj.KoordinatY, &obj.Status, &obj.FunFact, &obj.Program, &obj.Dampak, &obj.PetaHabitat, &obj.Dikunjungi, &obj.Gambar, &obj.GambarKecil)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Inserted"
	res.Data = arrobj

	return res, nil
}
