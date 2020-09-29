package formularios

import "github.com/graphql-go/graphql"

var NotificacionDescuentoComentarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "NotificacionDescuentoComentarios",
	Description: "Comentario de registro para la notificación de descuento",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID del estado",
		},
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado del registro",
		},
		"comentario": &graphql.Field{
			Type:        graphql.String,
			Description: "Comentario del agente para con el registro",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de creación del estado",
		},
		"formulario": &graphql.Field{
			Type:        graphql.String,
			Description: "Formulario al que pertenece el estado",
		},
		"usuario": &graphql.Field{
			Type:        graphql.String,
			Description: "Usuario que gestiona el registro",
		},
		"correoUsuario": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del usuario que gestiona el registro",
		},
		"idFormulario": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID del registro",
		},
	},
})

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
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado del registro",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
		"comentarios": &graphql.Field{
			Type:        graphql.NewList(NotificacionDescuentoComentarioType),
			Description: "Comentarios del registro",
		},
	},
})
