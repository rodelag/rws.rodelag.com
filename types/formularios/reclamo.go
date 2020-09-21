package formularios

import "github.com/graphql-go/graphql"

var ReclamoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Reclamo",
	Description: "Reclamo",
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
		"tipoReclamo": &graphql.Field{
			Type:        graphql.String,
			Description: "Tipo del reclamo",
		},
		"detalle": &graphql.Field{
			Type:        graphql.String,
			Description: "Detalle del reclamo",
		},
		"adjuntoDocumento": &graphql.Field{
			Type:        graphql.String,
			Description: "Documento o foto del reclamo",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
	},
})
