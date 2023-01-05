package process

import (
	"farmacia/dbinstance"
	"farmacia/model"
)

type Procesos struct{}

func (p *Procesos) AgregarFarmaco(fargs model.Args_AgregarFarmaco, r *int) error {
	fobj := model.Farmaco{
		Args_AgregarFarmaco: fargs,
	}

	resultado := dbinstance.DB.Create(&fobj)

	if resultado.Error != nil {
		*r = -1
		return resultado.Error
	}

	*r = int(resultado.RowsAffected)
	return nil
}

func (p *Procesos) ComprarFarmaco() error {
	return nil
}
