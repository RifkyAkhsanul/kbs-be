package models

import (
	"fmt"
	"net/http"

	"github.com/RifkyAkhsanul/kbs-be/db"
)

type Info struct {
	Id          int     `json:"id"`
	NamaSpesies string  `json:"namaspesies"`
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
	BahasaLatin string  `json:"bahasa_latin"`
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
		err = rows.Scan(&obj.Id, &obj.NamaSpesies, &obj.KoordinatX, &obj.KoordinatY, &obj.Status, &obj.FunFact, &obj.Program, &obj.Dampak, &obj.PetaHabitat, &obj.Dikunjungi, &obj.Gambar, &obj.GambarKecil, &obj.BahasaLatin)

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
	res.Message = fmt.Sprintf("id : %d", id)
	con, err := db.NewDriver()

	sqlStatement := "SELECT * FROM info WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	rows, err := stmt.Query(id)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaSpesies, &obj.KoordinatX, &obj.KoordinatY, &obj.Status, &obj.FunFact, &obj.Program, &obj.Dampak, &obj.PetaHabitat, &obj.Dikunjungi, &obj.Gambar, &obj.GambarKecil, &obj.BahasaLatin)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("id : %d", id)
	res.Data = arrobj

	return res, nil
}
