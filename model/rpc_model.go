package model

type Args_AgregarFarmaco struct {
	Nombre      string
	Laboratorio string
	Existencias uint
	Precio      float32
}

type Args_ComprarFarmaco struct {
	Nombre   string
	Cantidad uint
}
