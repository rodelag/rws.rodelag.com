package types

import "github.com/graphql-go/graphql"

var ComprobantePagoComentarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ComprobantePagoComentarios",
	Description: "Comentario de registro de los comporbante de pago",
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
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado del registro",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
		"comentarios": &graphql.Field{
			Type:        graphql.NewList(ComprobantePagoComentarioType),
			Description: "Comentarios del registro",
		},
	},
})
