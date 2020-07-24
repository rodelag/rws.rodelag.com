package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ComprobantePago struct {
	ID              int
	Nombre          string
	Apellido        string
	Cedula          string
	Correo          string
	Telefono        string
	ComprobantePago string
	FechaRegistro   string
}

func ListarComprobantePago() []ComprobantePago {
	rows, err := conexion().Query("SELECT * FROM formulario_comprobantepago;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	cp := ComprobantePago{}
	cps := []ComprobantePago{}

	for rows.Next() {
		err := rows.Scan(&cp.ID, &cp.Nombre, &cp.Apellido, &cp.Cedula, &cp.Correo, &cp.Telefono, &cp.ComprobantePago, &cp.FechaRegistro)
		logError("Problemas leer los datos: ", err)
		cps = append(cps, ComprobantePago{
			ID:              cp.ID,
			Nombre:          cp.Nombre,
			Apellido:        cp.Apellido,
			Cedula:          cp.Cedula,
			Correo:          cp.Correo,
			Telefono:        cp.Telefono,
			ComprobantePago: cp.ComprobantePago,
			FechaRegistro:   cp.FechaRegistro,
		})
	}
	return cps
}

func CrearComprobantePago(nombre string, apellido string, cedula string, correo string, telefono string, comprobantePago string) ComprobantePago {
	cp := ComprobantePago{
		Nombre:          nombre,
		Apellido:        apellido,
		Cedula:          cedula,
		Correo:          correo,
		Telefono:        telefono,
		ComprobantePago: comprobantePago,
		FechaRegistro:   time.Now().Format("2006-01-02 15:04:05"),
	}

	conn, err := conexion().Prepare("INSERT INTO formulario_comprobantepago (nombre, apellido, cedula, correo, telefono, comprobantePago, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(cp.Nombre, cp.Apellido, cp.Cedula, cp.Correo, cp.Telefono, cp.ComprobantePago, cp.FechaRegistro)

	return cp
}
