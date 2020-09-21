package formularios

import "github.com/graphql-go/graphql"

var PagoACHComentarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PagoACHComentarios",
	Description: "Comentario de registro para los pagos por ACH",
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
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado del registro",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
		"comentarios": &graphql.Field{
			Type:        graphql.NewList(PagoACHComentarioType),
			Description: "Comentarios del registro",
		},
	},
})
