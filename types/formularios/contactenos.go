package formularios

import "github.com/graphql-go/graphql"

var ContactenosType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Contactenos",
	Description: "Contáctenos",
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
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"telefono": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono del cliente",
		},
		"mensaje": &graphql.Field{
			Type:        graphql.String,
			Description: "Mensaje del cliente",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
	},
})
