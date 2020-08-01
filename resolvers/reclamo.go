package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Reclamo struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Telefono,
	TipoReclamo,
	Detalle,
	AdjuntoDocumento,
	FechaRegistro string
}

func ListarReclamo() []Reclamo {
	connMySQL := conexion()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_reclamo;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	re := Reclamo{}
	res := []Reclamo{}

	for rows.Next() {
		err := rows.Scan(
			&re.ID,
			&re.Nombre,
			&re.Apellido,
			&re.Cedula,
			&re.Correo,
			&re.Telefono,
			&re.TipoReclamo,
			&re.Detalle,
			&re.AdjuntoDocumento,
			&re.FechaRegistro,
		)
		logError("Problemas leer los datos: ", err)
		res = append(res, Reclamo{
			ID:               re.ID,
			Nombre:           re.Nombre,
			Apellido:         re.Apellido,
			Cedula:           re.Cedula,
			Correo:           re.Correo,
			Telefono:         re.Telefono,
			TipoReclamo:      re.TipoReclamo,
			Detalle:          re.Detalle,
			AdjuntoDocumento: re.AdjuntoDocumento,
			FechaRegistro:    re.FechaRegistro,
		})
	}
	return res
}

func CrearReclamo(nombre, apellido, cedula, correo, telefono, tipoReclamo, detalle, adjuntoDocumento string) Reclamo {
	re := Reclamo{
		Nombre:           nombre,
		Apellido:         apellido,
		Cedula:           cedula,
		Correo:           correo,
		Telefono:         telefono,
		TipoReclamo:      tipoReclamo,
		Detalle:          detalle,
		AdjuntoDocumento: adjuntoDocumento,
		FechaRegistro:    time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexion()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_reclamo (nombre, apellido, cedula, correo, telefono, tipoReclamo, detalle, adjuntoDocumento, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		re.Nombre,
		re.Apellido,
		re.Cedula,
		re.Correo,
		re.Telefono,
		re.TipoReclamo,
		re.Detalle,
		re.AdjuntoDocumento,
		re.FechaRegistro,
	)

	return re
}
