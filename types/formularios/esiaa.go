package formularios

import "github.com/graphql-go/graphql"

var EsiaaComentarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "EsiaaComentarios",
	Description: "Comentario de registro para la encuesta de satisfacción de instalación de AA",
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

var EsiaaType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ESIAA",
	Description: "Encuesta de Satisfacción Instalación AA",
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
			Description: "Apellido del lciente",
		},
		"cedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula del cliente",
		},
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"calificacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Calificación del cliente",
		},
		"atencion": &graphql.Field{
			Type:        graphql.String,
			Description: "Atención del cliente",
		},
		"resolverInstalacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Instalación resuelta",
		},
		"tiempoRazonable": &graphql.Field{
			Type:        graphql.String,
			Description: "Tiempo razonable",
		},
		"recomendacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Recomendación",
		},
		"calificacionManera": &graphql.Field{
			Type:        graphql.String,
			Description: "Calificación del cliente en general",
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
			Type:        graphql.NewList(EsiaaComentarioType),
			Description: "Comentarios del registro",
		},
	},
})
