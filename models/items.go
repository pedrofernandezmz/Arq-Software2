package model

type Item struct {
	Titulo      string `bson:"title"`
	Descripcion string `bson:"description"`
	Ciudad      string `bson:"city"`
	Estado      string `bson:"state"`
	Imagen      string `bson:"image"`
	Vendedor    string `bson:"seller"`
	//Precio float64 `bson:"price"`
}
