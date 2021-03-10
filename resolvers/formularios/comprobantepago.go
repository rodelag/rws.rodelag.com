package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
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
	Estado          string
	FechaRegistro   string
}

func conexionComprobantePago() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		"rodelag_ovnicom",
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Comprobante de pago) ", true, errMySQL)
	}
	return connMySQL
}

func VerComprobantePago(id int) ComprobantePago {
	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	comprobantePago := ComprobantePago{}

	err := connMySQL.QueryRow("SELECT * FROM formulario_comprobantepago WHERE id = ?;", id).Scan(
		&comprobantePago.ID,
		&comprobantePago.Nombre,
		&comprobantePago.Apellido,
		&comprobantePago.Cedula,
		&comprobantePago.Correo,
		&comprobantePago.Telefono,
		&comprobantePago.Estado,
		&comprobantePago.ComprobantePago,
		&comprobantePago.FechaRegistro,
	)
	utils.LogError("Problemas al leer registro: (Comprobante de pago) ", false, err)

	return comprobantePago
}

func ListarComprobantePago() []ComprobantePago {
	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_comprobantepago;")
	utils.LogError("Problemas al listar los registros de la base de datos: (Comprobante de pago) ", true, err)
	defer rows.Close()

	cp := ComprobantePago{}
	cps := []ComprobantePago{}

	for rows.Next() {
		err := rows.Scan(&cp.ID, &cp.Nombre, &cp.Apellido, &cp.Cedula, &cp.Correo, &cp.Telefono, &cp.Estado, &cp.ComprobantePago, &cp.FechaRegistro)
		utils.LogError("Problemas leer los datos: (Comprobante de pago) ", true, err)
		cps = append(cps, ComprobantePago{
			ID:              cp.ID,
			Nombre:          cp.Nombre,
			Apellido:        cp.Apellido,
			Cedula:          cp.Cedula,
			Correo:          cp.Correo,
			Telefono:        cp.Telefono,
			Estado:          cp.Estado,
			ComprobantePago: cp.ComprobantePago,
			FechaRegistro:   cp.FechaRegistro,
		})
	}
	return cps
}

func BusquedaComprobantePago(busqueda string) []ComprobantePago {
	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(
		fmt.Sprintf("SELECT * FROM formulario_comprobantepago WHERE (nombre LIKE '%%%s%%') OR (apellido LIKE '%%%s%%') OR (cedula LIKE '%%%s%%') OR (estado LIKE '%%%s%%');",
			busqueda,
			busqueda,
			busqueda,
			busqueda))

	utils.LogError("Problemas al listar los registros de la base de datos: (Comprobante de pago) ", true, err)
	defer rows.Close()

	cp := ComprobantePago{}
	cps := []ComprobantePago{}

	for rows.Next() {
		err := rows.Scan(&cp.ID, &cp.Nombre, &cp.Apellido, &cp.Cedula, &cp.Correo, &cp.Telefono, &cp.Estado, &cp.ComprobantePago, &cp.FechaRegistro)
		utils.LogError("Problemas leer los datos: (Comprobante de pago) ", true, err)
		cps = append(cps, ComprobantePago{
			ID:              cp.ID,
			Nombre:          cp.Nombre,
			Apellido:        cp.Apellido,
			Cedula:          cp.Cedula,
			Correo:          cp.Correo,
			Telefono:        cp.Telefono,
			Estado:          cp.Estado,
			ComprobantePago: cp.ComprobantePago,
			FechaRegistro:   cp.FechaRegistro,
		})
	}
	return cps
}

func EditarComprobantePago(id int, estado string) ComprobantePago {
	cp := ComprobantePago{
		ID:     id,
		Estado: estado,
	}

	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("UPDATE formulario_comprobantepago SET estado = ? WHERE id = ?;")
	utils.LogError("Problemas al crear el registro en la base de datos: (Comprobante de pago) ", true, err)
	defer conn.Close()

	conn.Exec(cp.Estado, cp.ID)

	return cp
}

func CrearComprobantePago(nombre, apellido, cedula, correo, telefono, comprobantePago string) ComprobantePago {
	cp := ComprobantePago{
		Nombre:          nombre,
		Apellido:        apellido,
		Cedula:          cedula,
		Correo:          correo,
		Telefono:        telefono,
		ComprobantePago: comprobantePago,
		FechaRegistro:   time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comprobantepago (nombre, apellido, cedula, correo, telefono, comprobantePago, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (Comprobante de pago) ", true, err)
	defer conn.Close()

	conn.Exec(cp.Nombre, cp.Apellido, cp.Cedula, cp.Correo, cp.Telefono, cp.ComprobantePago, cp.FechaRegistro)

	return cp
}
