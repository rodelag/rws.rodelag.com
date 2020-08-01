package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Contactenos struct {
	ID int
	Nombre,
	Apellido,
	Correo,
	Telefono,
	Mensaje,
	FechaRegistro string
}

func ListarContactenos() []Contactenos {
	connMySQL := conexion()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_contactenos;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	ct := Contactenos{}
	cts := []Contactenos{}

	for rows.Next() {
		err := rows.Scan(
			&ct.ID,
			&ct.Nombre,
			&ct.Apellido,
			&ct.Correo,
			&ct.Telefono,
			&ct.Mensaje,
			&ct.FechaRegistro,
		)
		logError("Problemas leer los datos: ", err)
		cts = append(cts, Contactenos{
			ID:            ct.ID,
			Nombre:        ct.Nombre,
			Apellido:      ct.Apellido,
			Correo:        ct.Correo,
			Telefono:      ct.Telefono,
			Mensaje:       ct.Mensaje,
			FechaRegistro: ct.FechaRegistro,
		})
	}
	return cts
}

func CrearContactenos(nombre, apellido, correo, telefono, mensaje string) Contactenos {
	ct := Contactenos{
		Nombre:        nombre,
		Apellido:      apellido,
		Correo:        correo,
		Telefono:      telefono,
		Mensaje:       mensaje,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexion()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_contactenos (nombre, apellido, correo, telefono, mensaje, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		ct.Nombre,
		ct.Apellido,
		ct.Correo,
		ct.Telefono,
		ct.Mensaje,
		ct.FechaRegistro,
	)

	return ct
}
