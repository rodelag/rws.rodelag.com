package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
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

func ListarNotificacionDescuento() []NotificacionDescuento {
	rows, err := conexion().Query("SELECT * FROM formulario_notificaciondescuento;")
	logError("Problemas al listar los registros de la base de datos: ", err)
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
		logError("Problemas leer los datos: ", err)
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

	conn, err := conexion().Prepare("INSERT INTO formulario_notificaciondescuento (nombre, apellido, sucursal, fecha, nombreProducto, fotoPrecio, codigoParte, cantidadVendidas, nombreCompetidor, precioCompetidor, precioRodelag, fotoCotizacion, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
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
