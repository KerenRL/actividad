package domain

type ITienda interface {
	SaveTienda(nombre string, ubicacion string)
	GetAll()
	UpdateTienda(id int32, nombre string, ubicacion string) error
	DeleteTienda(id int32) error
}

type tienda struct {
	ID     int32   `json:"id"`
	Nombre  string  `json:"nombre"`
	Ubicacion string  `json:"ubicacion"`
}

func NewTienda(nombre string, ubicacion string) *tienda {
	return &tienda{ID:1, Nombre: nombre, Ubicacion: ubicacion}
}

func (p *tienda) GetAll() ([]tienda, error) {
	return []tienda{}, nil
}

func (p *tienda) SetNombre(nombre string) {
	p.Nombre =nombre
}