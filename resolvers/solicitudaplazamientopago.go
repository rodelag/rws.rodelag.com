package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type SolicitudAplazamientoPago struct {
	ID int
	Nombre,
	Apellido,
	Correo,
	TelefonoCasa,
	Celular,
	TipoProducto,
	TipoCliente,
	TipoActividadEconomica,
	LugarTrabajo,
	MotivoSolicitud,
	DetalleMotivo,
	Cedula,
	Talonario,
	CartaMotivo,
	FechaRegistro string
}

func ListarSolicitudAplazamientoPago() []SolicitudAplazamientoPago {
	rows, err := conexion().Query("SELECT * FROM formulario_aplazamientopago;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	sap := SolicitudAplazamientoPago{}
	saps := []SolicitudAplazamientoPago{}

	for rows.Next() {
		err := rows.Scan(
			&sap.ID,
			&sap.Nombre,
			&sap.Apellido,
			&sap.Correo,
			&sap.TelefonoCasa,
			&sap.Celular,
			&sap.TipoProducto,
			&sap.TipoCliente,
			&sap.TipoActividadEconomica,
			&sap.LugarTrabajo,
			&sap.MotivoSolicitud,
			&sap.DetalleMotivo,
			&sap.Cedula,
			&sap.Talonario,
			&sap.CartaMotivo,
			&sap.FechaRegistro,
		)
		logError("Problemas leer los datos: ", err)
		saps = append(saps, SolicitudAplazamientoPago{
			ID:                     sap.ID,
			Nombre:                 sap.Nombre,
			Apellido:               sap.Apellido,
			Correo:                 sap.Correo,
			TelefonoCasa:           sap.TelefonoCasa,
			Celular:                sap.Celular,
			TipoProducto:           sap.TipoProducto,
			TipoCliente:            sap.TipoCliente,
			TipoActividadEconomica: sap.TipoActividadEconomica,
			LugarTrabajo:           sap.LugarTrabajo,
			MotivoSolicitud:        sap.MotivoSolicitud,
			DetalleMotivo:          sap.DetalleMotivo,
			Cedula:                 sap.Cedula,
			Talonario:              sap.Talonario,
			CartaMotivo:            sap.CartaMotivo,
			FechaRegistro:          sap.FechaRegistro,
		})
	}
	return saps
}

func CrearSolicitudAplazamientoPago(nombre, apellido, correo, telefonoCasa, celular, tipoProducto, tipoCliente, tipoActividadEconomica, lugarTrabajo, motivoSolicitud, detalleMotivo, cedula, talonario, cartaMotivo string) SolicitudAplazamientoPago {
	sap := SolicitudAplazamientoPago{
		Nombre:                 nombre,
		Apellido:               apellido,
		Correo:                 correo,
		TelefonoCasa:           telefonoCasa,
		Celular:                celular,
		TipoProducto:           tipoProducto,
		TipoCliente:            tipoCliente,
		TipoActividadEconomica: tipoActividadEconomica,
		LugarTrabajo:           lugarTrabajo,
		MotivoSolicitud:        motivoSolicitud,
		DetalleMotivo:          detalleMotivo,
		Cedula:                 cedula,
		Talonario:              talonario,
		CartaMotivo:            cartaMotivo,
		FechaRegistro:          time.Now().Format("2006-01-02 15:04:05"),
	}

	conn, err := conexion().Prepare("INSERT INTO formulario_aplazamientopago (nombre, apellido, correo, telefonoCasa, celular, tipoProducto, tipoCliente, tipoActividadEconomica, lugarTrabajo, motivoSolicitud, detalleMotivo, cedula, talonario, cartaMotivo, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		sap.Nombre,
		sap.Apellido,
		sap.Correo,
		sap.TelefonoCasa,
		sap.Celular,
		sap.TipoProducto,
		sap.TipoCliente,
		sap.TipoActividadEconomica,
		sap.LugarTrabajo,
		sap.MotivoSolicitud,
		sap.DetalleMotivo,
		sap.Cedula,
		sap.Talonario,
		sap.CartaMotivo,
		sap.FechaRegistro,
	)

	return sap
}
