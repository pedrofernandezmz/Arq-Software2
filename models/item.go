package model

type Item struct {
	Titulo      string `bson:"text"`
	Descripcion string `bson:"text"`
	Direccion   string `bson:"text"`
	Ciudad      string `bson:"text"`
	Provincia   string `bson:"text"`
	Imagen      string `bson:"text"`
	Imagen2     string `bson:"text"`
	Vendedor    string `bson:"text"`
}
