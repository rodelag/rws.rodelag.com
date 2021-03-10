package formularios

import "github.com/graphql-go/graphql"

var ArregloPagoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ArregloPago",
	Description: "Arreglo de Pago",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID del registro",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del cliente",
		},
		"apellido": &graphql.Field{
			Type:        graphql.String,
			Description: "Apellido del cliente",
		},
		"cedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula del cliente",
		},
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"telefono": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono del cliente",
		},
		"celular": &graphql.Field{
			Type:        graphql.String,
			Description: "Celular del cliente",
		},
		"direccionDomicilio": &graphql.Field{
			Type:        graphql.String,
			Description: "Dirección del cliente",
		},
		"telefonoTrabajo": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono de trabajo del cliente",
		},
		"lugarTrabajo": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del lugar de trabajo del cliente",
		},
		"sector": &graphql.Field{
			Type:        graphql.String,
			Description: "Sector de cliente",
		},
		"motivoArregloPago": &graphql.Field{
			Type:        graphql.String,
			Description: "Motivo del arreglo de pago",
		},
		"fotoCedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto de la cédula",
		},
		"comprobanteAbono": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto del comprobante de abono",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado",
		},
	},
})
