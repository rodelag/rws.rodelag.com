package formularios

import "github.com/graphql-go/graphql"

var PrevencionFraudeComentarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PrevencionFraudeComentarios",
	Description: "Comentario de registro para Prevención de fraude",
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

var PrevencionFraudeType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PrevencionFraude",
	Description: "Prevención de fraude",
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
		"fechaNacimiento": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de nacimiento",
		},
		"lugarResidencia": &graphql.Field{
			Type:        graphql.String,
			Description: "Lugar de residencia",
		},
		"celular": &graphql.Field{
			Type:        graphql.String,
			Description: "Celular del cliente",
		},
		"fotoCedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto de la cédula",
		},
		"fotoTarjeta": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto frontal de la tarjeta",
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
			Type:        graphql.NewList(PrevencionFraudeComentarioType),
			Description: "Comentarios del registro",
		},
	},
})
