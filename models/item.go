package model

type Item struct {
	Titulo      string `bson:"title"`
	Descripcion string `bson:"description"`
	Direccion   string `bson:"direction"`
	Ciudad      string `bson:"city"`
	Provincia   string `bson:"province"`
	Imagen      string `bson:"image"`
	Imagen2     string `bson:"image2"`
	Vendedor    string `bson:"seller"`
}
