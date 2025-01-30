package infraestructure

import (
	"actividad/src/perfumes/application"
	"encoding/json"
	"net/http"
)

type CreatePerfumeController struct {
	useCase application.CreatePerfume
}

func NewCreatePerfumeController(useCase application.CreatePerfume) *CreatePerfumeController {
	return &CreatePerfumeController{useCase: useCase}
}

type RequestBody struct {
	Marca  string  `json:"marca"`
	Modelo string  `json:"modelo"`
	Precio float32 `json:"precio"`
}

func (cp_c *CreatePerfumeController) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error al leer el JSON", http.StatusBadRequest)
		return
	}

	cp_c.useCase.Execute(body.Marca, body.Modelo, body.Precio)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Perfume agregado correctamente"})
}
