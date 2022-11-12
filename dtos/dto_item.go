package dtos

type ItemDTO struct {
	Titulo      string `json:"title"`
	Id          string `json:"id"`
	Descripcion string `json:"description"`
	Direccion   string `json:"direction"`
	Ciudad      string `json:"city"`
	Provincia   string `json:"province"`
	Imagen      string `json:"image"`
	Imagen2     string `json:"image2"`
	Vendedor    string `json:"seller"`
}
