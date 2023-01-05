package model

type Farmaco struct {
	Nro                 uint                `gorm:"primaryKey"`
	Args_AgregarFarmaco Args_AgregarFarmaco `gorm:"embedded"`
	// Nombre      string
	// Laboratorio string
	// Existencias uint
	// Precio      float32
}

type Compra struct {
	Nro        uint `gorm:"primaryKey"`
	FarmacoNro uint
	Farmaco    Farmaco `gorm:"foreignKey:FarmacoNro"`
	Cantidad   uint
}
