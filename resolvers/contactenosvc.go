package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ContactenosVC struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Telefono,
	ActividadEconomica,
	DetalleSolicitud,
	FechaRegistro string
}

func ListarContactenosVC() []ContactenosVC {
	rows, err := conexion().Query("SELECT * FROM formulario_contactenosvc;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	ct := ContactenosVC{}
	cts := []ContactenosVC{}

	for rows.Next() {
		err := rows.Scan(
			&ct.ID,
			&ct.Nombre,
			&ct.Apellido,
			&ct.Cedula,
			&ct.Correo,
			&ct.Telefono,
			&ct.ActividadEconomica,
			&ct.DetalleSolicitud,
			&ct.FechaRegistro,
		)
		logError("Problemas leer los datos: ", err)
		cts = append(cts, ContactenosVC{
			ID:                 ct.ID,
			Nombre:             ct.Nombre,
			Apellido:           ct.Apellido,
			Cedula:             ct.Cedula,
			Correo:             ct.Correo,
			Telefono:           ct.Telefono,
			ActividadEconomica: ct.ActividadEconomica,
			DetalleSolicitud:   ct.DetalleSolicitud,
			FechaRegistro:      ct.FechaRegistro,
		})
	}
	return cts
}

func CrearContactenosVC(nombre, apellido, cedula, correo, telefono, actividadEconomica, detalleSolicitud string) ContactenosVC {
	ct := ContactenosVC{
		Nombre:             nombre,
		Apellido:           apellido,
		Cedula:             cedula,
		Correo:             correo,
		Telefono:           telefono,
		ActividadEconomica: actividadEconomica,
		DetalleSolicitud:   detalleSolicitud,
		FechaRegistro:      time.Now().Format("2006-01-02 15:04:05"),
	}

	conn, err := conexion().Prepare("INSERT INTO formulario_contactenosvc (nombre, apellido, cedula, correo, telefono, actividadEconomica, detalleSolicitud, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		ct.Nombre,
		ct.Apellido,
		ct.Cedula,
		ct.Correo,
		ct.Telefono,
		ct.ActividadEconomica,
		ct.DetalleSolicitud,
		ct.FechaRegistro,
	)

	return ct
}
