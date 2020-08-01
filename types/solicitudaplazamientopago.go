package types

import "github.com/graphql-go/graphql"

var SolicitudAplazamientoPagoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "SolicitudAplazamientoPago",
	Description: "Solicitud de Aplazamiento de Pagos",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "id del registro",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del cliente",
		},
		"apellido": &graphql.Field{
			Type:        graphql.String,
			Description: "Apellido del cliente",
		},
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"telefonoCasa": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono de la casa del cliente",
		},
		"cedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula del cliente",
		},
		"celular": &graphql.Field{
			Type:        graphql.String,
			Description: "Celular del cliente",
		},
		"tipoProducto": &graphql.Field{
			Type:        graphql.String,
			Description: "Tipo de Producto",
		},
		"tipoCliente": &graphql.Field{
			Type:        graphql.String,
			Description: "Tipo de cliente",
		},
		"tipoActividadEconomica": &graphql.Field{
			Type:        graphql.String,
			Description: "Tipo de Actividad Económica",
		},
		"lugarTrabajo": &graphql.Field{
			Type:        graphql.String,
			Description: "Lugar de Trabajo",
		},
		"motivoSolicitud": &graphql.Field{
			Type:        graphql.String,
			Description: "Motivo de la solicitud",
		},
		"detalleMotivo": &graphql.Field{
			Type:        graphql.String,
			Description: "Detalle del Motivo",
		},
		"talonario": &graphql.Field{
			Type:        graphql.String,
			Description: "Talonario",
		},
		"cartaMotivo": &graphql.Field{
			Type:        graphql.String,
			Description: "Carta del motivo",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
	},
})
