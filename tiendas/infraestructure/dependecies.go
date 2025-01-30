package infraestructure

func Iniciar() {
	ps := NewMySQL()
	SetupRouter(ps)
}