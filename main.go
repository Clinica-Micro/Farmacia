package main

import (
	"farmacia/dbinstance"
	"farmacia/model"
	"farmacia/process"
	"log"
)

func main() {
	db, err := dbinstance.CrearConexion()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Farmaco{}, &model.Compra{})

	var p process.Procesos
	far := model.Args_AgregarFarmaco{
		Nombre:      "Amoxicilina",
		Laboratorio: "Genfar",
		Existencias: 30,
		Precio:      24.50,
	}
	r := new(int)

	err = p.AgregarFarmaco(far, r)
	if err != nil {
		log.Fatal(err)
	}

}
