package domain

type IPerfume interface {
	SavePerfume(marca string, modelo string, precio float32)
	GetAll()
	UpdatePerfume(id int32, marca string, modelo string, precio float32) error
	DeletePerfume(id int32) error
}

type perfume struct {
	ID     int32   `json:"id"`
	Marca  string  `json:"marca"`
	Modelo string  `json:"modelo"`
	Precio float32 `json:"precio"`
}

func NewPerfume(marca string, modelo string, precio float32) *perfume {
	return &perfume{ID:1, Marca: marca, Modelo: modelo, Precio: precio}
}

func (p *perfume) GetAll() ([]perfume, error) {
	return []perfume{}, nil
}

func (p *perfume) SetMarca(marca string) {
	p.Marca = marca
}