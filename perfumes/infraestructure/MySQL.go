package infraestructure

import (
	"actividad/src/config"
	"actividad/src/perfumes/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IPerfume = (*MySQL)(nil)

func NewMySQL() domain.IPerfume {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}
func (mysql *MySQL) SavePerfume(marca string, modelo string, precio float32) {
	query := "INSERT INTO perfume (marca, modelo, precio) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, marca, modelo, precio)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Perfume guardado correctamente: Marca: %s Modelo: %s - Precio: %.2f", marca, modelo, precio)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
}

func (mysql *MySQL) GetAll() {
	query := "SELECT * FROM perfume"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	for rows.Next() {
		var id int
		var marca, modelo string
		var precio float32
		if err := rows.Scan(&id, &marca, &modelo, &precio); err != nil {
			fmt.Printf("Error al escanear la fila: %v\n", err)
		}
		fmt.Printf("ID: %d, Marca: %s, Modelo: %s, Precio: %.2f\n", id, marca, modelo, precio)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterando sobre las filas: %v\n", err)
	}
}

func (mysql *MySQL) UpdatePerfume(id int32, marca string, modelo string, precio float32) error {
	query := "UPDATE perfume SET marca = ?, modelo = ?, precio = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, marca, modelo, precio, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

func (mysql *MySQL) DeletePerfume(id int32) error {
	query := "DELETE FROM perfume WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}