package infraestructure

func Init() {
	ps := NewMySQL()
	SetupRouter(ps)
}
