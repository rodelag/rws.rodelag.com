package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type SolicitudEstadoCuenta struct {
	ID            int
	Nombre        string
	Apellido      string
	Correo        string
	Telefono      string
	Cedula        string
	FechaRegistro string
}

func ListarEstadoCuenta() []SolicitudEstadoCuenta {
	connMySQL := conexion()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_estadocuenta;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	sec := SolicitudEstadoCuenta{}
	ssec := []SolicitudEstadoCuenta{}

	for rows.Next() {
		err := rows.Scan(&sec.ID, &sec.Nombre, &sec.Apellido, &sec.Correo, &sec.Telefono, &sec.Cedula, &sec.FechaRegistro)
		logError("Problemas leer los datos: ", err)
		ssec = append(ssec, SolicitudEstadoCuenta{
			ID:            sec.ID,
			Nombre:        sec.Nombre,
			Apellido:      sec.Apellido,
			Correo:        sec.Correo,
			Telefono:      sec.Telefono,
			Cedula:        sec.Cedula,
			FechaRegistro: sec.FechaRegistro,
		})
	}
	return ssec
}

func CrearSolicitudEstadoCuenta(nombre string, apellido string, correo string, telefono string, cedula string) SolicitudEstadoCuenta {
	sec := SolicitudEstadoCuenta{
		Nombre:        nombre,
		Apellido:      apellido,
		Correo:        correo,
		Telefono:      telefono,
		Cedula:        cedula,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexion()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_estadocuenta (nombre, apellido, correo, telefono, cedula, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(sec.Nombre, sec.Apellido, sec.Correo, sec.Telefono, sec.Cedula, sec.FechaRegistro)

	return sec
}
