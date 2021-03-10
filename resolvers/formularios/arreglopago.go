package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type ArregloPago struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Telefono,
	Celular,
	DireccionDomicilio,
	TelefonoTrabajo,
	LugarTrabajo,
	Sector,
	MotivoArregloPago,
	FotoCedula,
	ComprobanteAbono,
	FechaRegistro,
	Estado string
}

func conexionArregloPago() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		"rodelag_ovnicom",
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (ArregloPago) ", true, errMySQL)
	}
	return connMySQL
}

func VerArregloPago(id int) ArregloPago {
	connMySQL := conexionArregloPago()
	defer connMySQL.Close()

	sec := ArregloPago{}

	err := connMySQL.QueryRow("SELECT * FROM formulario_arreglopago WHERE id = ?;", id).Scan(
		&sec.ID,
		&sec.Nombre,
		&sec.Apellido,
		&sec.Cedula,
		&sec.Correo,
		&sec.Telefono,
		&sec.Celular,
		&sec.DireccionDomicilio,
		&sec.TelefonoTrabajo,
		&sec.LugarTrabajo,
		&sec.Sector,
		&sec.MotivoArregloPago,
		&sec.FotoCedula,
		&sec.ComprobanteAbono,
		&sec.FechaRegistro,
		&sec.Estado,
	)
	utils.LogError("Problemas al leer registro: (ArregloPago) ", false, err)

	return sec
}

func ListarArregloPago() []ArregloPago {
	connMySQL := conexionArregloPago()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_arreglopago;")
	utils.LogError("Problemas al listar los registros de la base de datos: (ArregloPago) ", true, err)
	defer rows.Close()

	sec := ArregloPago{}
	ssec := []ArregloPago{}

	for rows.Next() {
		err := rows.Scan(&sec.ID, &sec.Nombre, &sec.Apellido, &sec.Cedula, &sec.Correo, &sec.Telefono, &sec.Celular, &sec.DireccionDomicilio, &sec.TelefonoTrabajo, &sec.LugarTrabajo, &sec.Sector, &sec.MotivoArregloPago, &sec.FotoCedula, &sec.ComprobanteAbono, &sec.FechaRegistro, &sec.Estado)
		utils.LogError("Problemas leer los datos: (ArregloPago) ", true, err)
		ssec = append(ssec, ArregloPago{
			ID:                 sec.ID,
			Nombre:             sec.Nombre,
			Apellido:           sec.Apellido,
			Cedula:             sec.Cedula,
			Correo:             sec.Correo,
			Telefono:           sec.Telefono,
			Celular:            sec.Celular,
			DireccionDomicilio: sec.DireccionDomicilio,
			TelefonoTrabajo:    sec.TelefonoTrabajo,
			LugarTrabajo:       sec.LugarTrabajo,
			Sector:             sec.Sector,
			MotivoArregloPago:  sec.MotivoArregloPago,
			FotoCedula:         sec.FotoCedula,
			ComprobanteAbono:   sec.ComprobanteAbono,
			FechaRegistro:      sec.FechaRegistro,
			Estado:             sec.Estado,
		})
	}
	return ssec
}

func BusquedaArregloPago(busqueda string) []ArregloPago {
	connMySQL := conexionArregloPago()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(
		fmt.Sprintf("SELECT * FROM formulario_arreglopago WHERE (nombre LIKE '%%%s%%') OR (apellido LIKE '%%%s%%') OR (cedula LIKE '%%%s%%') OR (estado LIKE '%%%s%%');",
			busqueda,
			busqueda,
			busqueda,
			busqueda))

	utils.LogError("Problemas al listar los registros de la base de datos: (Solicitud Estado Cuenta) ", true, err)
	defer rows.Close()

	sec := ArregloPago{}
	ssec := []ArregloPago{}

	for rows.Next() {
		err := rows.Scan(&sec.ID, &sec.Nombre, &sec.Apellido, &sec.Cedula, &sec.Correo, &sec.Telefono, &sec.Celular, &sec.DireccionDomicilio, &sec.TelefonoTrabajo, &sec.LugarTrabajo, &sec.Sector, &sec.MotivoArregloPago, &sec.FotoCedula, &sec.ComprobanteAbono, &sec.FechaRegistro, &sec.Estado)
		utils.LogError("Problemas leer los datos: (ArregloPago) ", true, err)
		ssec = append(ssec, ArregloPago{
			ID:                 sec.ID,
			Nombre:             sec.Nombre,
			Apellido:           sec.Apellido,
			Cedula:             sec.Cedula,
			Correo:             sec.Correo,
			Telefono:           sec.Telefono,
			Celular:            sec.Celular,
			DireccionDomicilio: sec.DireccionDomicilio,
			TelefonoTrabajo:    sec.TelefonoTrabajo,
			LugarTrabajo:       sec.LugarTrabajo,
			Sector:             sec.Sector,
			MotivoArregloPago:  sec.MotivoArregloPago,
			FotoCedula:         sec.FotoCedula,
			ComprobanteAbono:   sec.ComprobanteAbono,
			FechaRegistro:      sec.FechaRegistro,
			Estado:             sec.Estado,
		})
	}
	return ssec
}

func EditarArregloPago(id int, estado string) ArregloPago {
	sec := ArregloPago{
		ID:     id,
		Estado: estado,
	}

	connMySQL := conexionArregloPago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("UPDATE formulario_arreglopago SET estado = ? WHERE id = ?;")
	utils.LogError("Problemas al crear el registro en la base de datos: (Solicitud Estado Cuenta) ", true, err)
	defer conn.Close()

	conn.Exec(sec.Estado, sec.ID)

	return sec
}

func CrearArregloPago(nombre, apellido, cedula, correo, telefono, celular, direccionDomicilio, telefonoTrabajo, lugarTrabajo, sector, motivoArregloPago, fotoCedula, comprobanteAbono string) ArregloPago {
	sec := ArregloPago{
		Nombre:             nombre,
		Apellido:           apellido,
		Cedula:             cedula,
		Correo:             correo,
		Telefono:           telefono,
		Celular:            celular,
		DireccionDomicilio: direccionDomicilio,
		TelefonoTrabajo:    telefonoTrabajo,
		LugarTrabajo:       lugarTrabajo,
		Sector:             sector,
		MotivoArregloPago:  motivoArregloPago,
		FotoCedula:         fotoCedula,
		ComprobanteAbono:   comprobanteAbono,
		FechaRegistro:      time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionArregloPago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_arreglopago (nombre, apellido, cedula, correo, telefono, celular, direccion_domicilio, telefono_trabajo, lugar_trabajo, sector, motivo_arreglo_pago, foto_cedula, comprobante_abono, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (ArregloPago) ", false, err)
	defer conn.Close()

	conn.Exec(sec.Nombre, sec.Apellido, sec.Cedula, sec.Correo, sec.Telefono, sec.Celular, sec.DireccionDomicilio, sec.TelefonoTrabajo, sec.LugarTrabajo, sec.Sector, sec.MotivoArregloPago, sec.FotoCedula, sec.ComprobanteAbono, sec.FechaRegistro)

	return sec
}
