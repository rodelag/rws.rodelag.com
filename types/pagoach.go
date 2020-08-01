package types

import "github.com/graphql-go/graphql"

var PagoACHType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PagoACH",
	Description: "Formulario de pago por ACH",
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
		"titularCuenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Titular de la cuenta",
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
		"compraOrigen": &graphql.Field{
			Type:        graphql.String,
			Description: "Lugar donde hizo la compra",
		},
		"numeroOrden": &graphql.Field{
			Type:        graphql.String,
			Description: "Pedido o Cotización",
		},
		"fotoComprobante": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto del comprobante de pago",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
	},
})
