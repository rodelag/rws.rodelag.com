package types

import "github.com/graphql-go/graphql"

var NotificacionDescuentoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NotificacionDescuento",
	Description: "Notificación para descuento",
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
		"sucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal",
		},
		"fecha": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha",
		},
		"nombreProducto": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del producto",
		},
		"fotoPrecio": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto del precio",
		},
		"codigoParte": &graphql.Field{
			Type:        graphql.String,
			Description: "Código de parte",
		},
		"cantidadVendidas": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad vendidas",
		},
		"nombreCompetidor": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del competidor",
		},
		"precioCompetidor": &graphql.Field{
			Type:        graphql.String,
			Description: "Precio del competidor",
		},
		"precioRodelag": &graphql.Field{
			Type:        graphql.String,
			Description: "Precio Rodelag",
		},
		"fotoCotizacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto de la cotización",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
	},
})
