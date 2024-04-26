package models

import (
	"net/http"

	"github.com/RifkyAkhsanul/kbs-be/db"
)

type Info struct {
	Id          int    `json:"id"`
	NamaSpesies string `json:"namaspesies"`
	KoordinatX  int    `json:"koordinatx"`
	KoordinatY  int    `json:"koordinaty"`
	Status      string `json:"status"`
	FunFact     string `json:"funfact"`
	Program     string `json:"program"`
	Dampak      string `json:"dampak"`
	Games       string `json:"games"`
	Jawaban     string `json:"jawaban"`
	Dikunjungi  bool   `json:"dikunjungi"`
}

func GetInfo() (Response, error) {
	var obj Info
	var arrobj []Info
	var res Response

	cons, err := db.NewDriver()
	if err != nil {
		return res, err
	}

	sqlstatement := "SELECT + FROM info"

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
