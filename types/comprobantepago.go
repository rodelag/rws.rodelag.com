package types

import "github.com/graphql-go/graphql"

var ComprobantePagoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ComprobantePago",
	Description: "Comprobante de Pago",
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
		"comprobantePago": &graphql.Field{
			Type:        graphql.String,
			Description: "Comprobante de pago",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Fecha del registro",
		},
	},
})
