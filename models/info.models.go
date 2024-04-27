package models

import (
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
	Games       string  `json:"games"`
	Jawaban     string  `json:"jawaban"`
	Dikunjungi  bool    `json:"dikunjungi"`
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
		err = rows.Scan(&obj.Id, &obj.NamaSpesies, &obj.KoordinatX, &obj.KoordinatY, &obj.Status, &obj.FunFact, &obj.Program, &obj.Dampak, &obj.Games, &obj.Jawaban, &obj.Dikunjungi)

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
		err = rows.Scan(&obj.Id, &obj.NamaSpesies, &obj.KoordinatX, &obj.KoordinatY, &obj.Status, &obj.FunFact, &obj.Program, &obj.Dampak, &obj.Games, &obj.Jawaban, &obj.Dikunjungi)

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
