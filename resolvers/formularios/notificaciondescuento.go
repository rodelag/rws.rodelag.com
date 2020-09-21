package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type NotificacionDescuento struct {
	ID int
	Nombre,
	Apellido,
	Sucursal,
	Fecha,
	NombreProducto,
	FotoPrecio,
	CodigoParte,
	CantidadVendidas,
	NombreCompetidor,
	PrecioCompetidor,
	PrecioRodelag,
	FotoCotizacion,
	FechaRegistro string
}

func conexionNotificacionDescuento() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		viper.GetString("basedatos.mysql.rodelag.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func ListarNotificacionDescuento() []NotificacionDescuento {
	connMySQL := conexionNotificacionDescuento()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_notificaciondescuento;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	nd := NotificacionDescuento{}
	nds := []NotificacionDescuento{}

	for rows.Next() {
		err := rows.Scan(
			&nd.ID,
			&nd.Nombre,
			&nd.Apellido,
			&nd.Sucursal,
			&nd.Fecha,
			&nd.NombreProducto,
			&nd.FotoPrecio,
			&nd.CodigoParte,
			&nd.CantidadVendidas,
			&nd.NombreCompetidor,
			&nd.PrecioCompetidor,
			&nd.PrecioRodelag,
			&nd.FotoCotizacion,
			&nd.FechaRegistro,
		)
		utils.LogError("Problemas leer los datos: ", err)
		nds = append(nds, NotificacionDescuento{
			ID:               nd.ID,
			Nombre:           nd.Nombre,
			Apellido:         nd.Apellido,
			Sucursal:         nd.Sucursal,
			Fecha:            nd.Fecha,
			NombreProducto:   nd.NombreProducto,
			FotoPrecio:       nd.FotoPrecio,
			CodigoParte:      nd.CodigoParte,
			CantidadVendidas: nd.CantidadVendidas,
			NombreCompetidor: nd.NombreCompetidor,
			PrecioCompetidor: nd.PrecioCompetidor,
			PrecioRodelag:    nd.PrecioRodelag,
			FotoCotizacion:   nd.FotoCotizacion,
			FechaRegistro:    nd.FechaRegistro,
		})
	}
	return nds
}

func CrearNotificacionDescuento(nombre, apellido, sucursal, fecha, nombreProducto, fotoPrecio, codigoParte, cantidadVendidas, nombreCompetidor, precioCompetidor, precioRodelag, fotoCotizacion string) NotificacionDescuento {
	nd := NotificacionDescuento{
		Nombre:           nombre,
		Apellido:         apellido,
		Sucursal:         sucursal,
		Fecha:            fecha,
		NombreProducto:   nombreProducto,
		FotoPrecio:       fotoPrecio,
		CodigoParte:      codigoParte,
		CantidadVendidas: cantidadVendidas,
		NombreCompetidor: nombreCompetidor,
		PrecioCompetidor: precioCompetidor,
		PrecioRodelag:    precioRodelag,
		FotoCotizacion:   fotoCotizacion,
		FechaRegistro:    time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionNotificacionDescuento()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_notificaciondescuento (nombre, apellido, sucursal, fecha, nombreProducto, fotoPrecio, codigoParte, cantidadVendidas, nombreCompetidor, precioCompetidor, precioRodelag, fotoCotizacion, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		nd.Nombre,
		nd.Apellido,
		nd.Sucursal,
		nd.Fecha,
		nd.NombreProducto,
		nd.FotoPrecio,
		nd.CodigoParte,
		nd.CantidadVendidas,
		nd.NombreCompetidor,
		nd.PrecioCompetidor,
		nd.PrecioRodelag,
		nd.FotoCotizacion,
		nd.FechaRegistro,
	)

	return nd
}
