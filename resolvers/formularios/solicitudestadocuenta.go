package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type SolicitudEstadoCuenta struct {
	ID            int
	Nombre        string
	Apellido      string
	Correo        string
	Telefono      string
	Cedula        string
	Estado        string
	FechaRegistro string
}

func conexionSolicitudEstadoCuenta() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		"rodelag_ovnicom",
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (SolicitudEstadoCuenta) ", true, errMySQL)
	}
	return connMySQL
}

func VerSolicitudEstadoCuenta(id int) SolicitudEstadoCuenta {
	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	sec := SolicitudEstadoCuenta{}

	err := connMySQL.QueryRow("SELECT * FROM formulario_estadocuenta WHERE id = ?;", id).Scan(
		&sec.ID,
		&sec.Nombre,
		&sec.Apellido,
		&sec.Correo,
		&sec.Telefono,
		&sec.Cedula,
		&sec.Estado,
		&sec.FechaRegistro,
	)
	utils.LogError("Problemas al leer registro: (SolicitudEstadoCuenta) ", false, err)

	return sec
}

func ListarEstadoCuenta() []SolicitudEstadoCuenta {
	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_estadocuenta;")
	utils.LogError("Problemas al listar los registros de la base de datos: (SolicitudEstadoCuenta) ", true, err)
	defer rows.Close()

	sec := SolicitudEstadoCuenta{}
	ssec := []SolicitudEstadoCuenta{}

	for rows.Next() {
		err := rows.Scan(&sec.ID, &sec.Nombre, &sec.Apellido, &sec.Correo, &sec.Telefono, &sec.Cedula, &sec.Estado, &sec.FechaRegistro)
		utils.LogError("Problemas leer los datos: (SolicitudEstadoCuenta) ", true, err)
		ssec = append(ssec, SolicitudEstadoCuenta{
			ID:            sec.ID,
			Nombre:        sec.Nombre,
			Apellido:      sec.Apellido,
			Correo:        sec.Correo,
			Telefono:      sec.Telefono,
			Cedula:        sec.Cedula,
			Estado:        sec.Estado,
			FechaRegistro: sec.FechaRegistro,
		})
	}
	return ssec
}

func BusquedaEstadoCuenta(busqueda string) []SolicitudEstadoCuenta {
	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(
		fmt.Sprintf("SELECT * FROM formulario_estadocuenta WHERE (nombre LIKE '%%%s%%') OR (apellido LIKE '%%%s%%') OR (cedula LIKE '%%%s%%') OR (estado LIKE '%%%s%%');",
			busqueda,
			busqueda,
			busqueda,
			busqueda))

	utils.LogError("Problemas al listar los registros de la base de datos: (Solicitud Estado Cuenta) ", true, err)
	defer rows.Close()

	sec := SolicitudEstadoCuenta{}
	ssec := []SolicitudEstadoCuenta{}

	for rows.Next() {
		err := rows.Scan(&sec.ID, &sec.Nombre, &sec.Apellido, &sec.Correo, &sec.Telefono, &sec.Cedula, &sec.Estado, &sec.FechaRegistro)
		utils.LogError("Problemas leer los datos: (SolicitudEstadoCuenta) ", true, err)
		ssec = append(ssec, SolicitudEstadoCuenta{
			ID:            sec.ID,
			Nombre:        sec.Nombre,
			Apellido:      sec.Apellido,
			Correo:        sec.Correo,
			Telefono:      sec.Telefono,
			Cedula:        sec.Cedula,
			Estado:        sec.Estado,
			FechaRegistro: sec.FechaRegistro,
		})
	}
	return ssec
}

func EditarEstadoCuenta(id int, estado string) SolicitudEstadoCuenta {
	sec := SolicitudEstadoCuenta{
		ID:     id,
		Estado: estado,
	}

	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("UPDATE formulario_estadocuenta SET estado = ? WHERE id = ?;")
	utils.LogError("Problemas al crear el registro en la base de datos: (Solicitud Estado Cuenta) ", true, err)
	defer conn.Close()

	conn.Exec(sec.Estado, sec.ID)

	return sec
}

func CrearSolicitudEstadoCuenta(nombre, apellido, correo, telefono, cedula string) SolicitudEstadoCuenta {
	sec := SolicitudEstadoCuenta{
		Nombre:        nombre,
		Apellido:      apellido,
		Correo:        correo,
		Telefono:      telefono,
		Cedula:        cedula,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_estadocuenta (nombre, apellido, correo, telefono, cedula, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (SolicitudEstadoCuenta) ", false, err)
	defer conn.Close()

	conn.Exec(sec.Nombre, sec.Apellido, sec.Correo, sec.Telefono, sec.Cedula, sec.FechaRegistro)

	return sec
}
