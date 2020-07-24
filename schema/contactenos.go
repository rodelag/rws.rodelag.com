package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func ContactenosQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"contactenos_listar": {
			Type: graphql.NewList(types.ContactenosType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarContactenos(), nil
			},
		},
	}
	return schemas
}

func ContactenosMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"contactenos_crear": &graphql.Field{
			Type:        types.ContactenosType,
			Description: "Contáctenos",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"telefono": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono del cliente",
				},
				"mensaje": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Mensaje del cliente",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(params.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				nombre, _ := params.Args["nombre"].(string)
				apellido, _ := params.Args["apellido"].(string)
				correo, _ := params.Args["correo"].(string)
				telefono, _ := params.Args["telefono"].(string)
				mensaje, _ := params.Args["mensaje"].(string)

				return resolvers.CrearContactenos(nombre, apellido, correo, telefono, mensaje), nil
			},
		},
	}
	return schemas
}
