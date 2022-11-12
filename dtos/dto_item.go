package dtos

type ItemDTO struct {
	Titulo      string `json:"title"`
	Id          string `json:"id"`
	Descripcion string `json:"description"`
	Ciudad      string `json:"city"`
	Estado      string `json:"state"`
	Imagen      string `json:"image"`
	Vendedor    string `json:"seller"`
	//Precio int `json:"price"`
}
