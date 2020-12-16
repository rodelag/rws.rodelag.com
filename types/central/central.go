package central

import "github.com/graphql-go/graphql"

var CentralTelefonicaExtensionType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "CentralTelefonicaExtension",
	Description: "Directorio de la Central Teléfonica",
	Fields: graphql.Fields{
		"registroNombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del colaborador",
		},
		"registroExtension": &graphql.Field{
			Type:        graphql.String,
			Description: "Extensión del colaborador",
		},
	},
})

var CentralTelefonicaType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "CentralTelefonica",
	Description: "Directorio de la Central Teléfonica",
	Fields: graphql.Fields{
		"registroDetalle": &graphql.Field{
			Type:        graphql.NewList(CentralTelefonicaExtensionType),
			Description: "Detalle",
		},
	},
})
