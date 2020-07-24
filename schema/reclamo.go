package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func ReclamoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"reclamo_listar": {
			Type: graphql.NewList(types.ReclamoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarReclamo(), nil
			},
		},
	}
	return schemas
}

func ReclamoMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"reclamo_crear": &graphql.Field{
			Type:        types.ReclamoType,
			Description: "Creación de reclamo",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"telefono": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono del cliente",
				},
				"tipoReclamo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tipo del reclamo",
				},
				"detalle": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Detalle del reclamo",
				},
				"adjuntoDocumento": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Documento o foto del reclamo",
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
				cedula, _ := params.Args["cedula"].(string)
				correo, _ := params.Args["correo"].(string)
				telefono, _ := params.Args["telefono"].(string)
				tipoReclamo, _ := params.Args["tipoReclamo"].(string)
				detalle, _ := params.Args["detalle"].(string)
				adjuntoDocumento, _ := params.Args["adjuntoDocumento"].(string)

				return resolvers.CrearReclamo(nombre, apellido, cedula, correo, telefono, tipoReclamo, detalle, adjuntoDocumento), nil
			},
		},
	}
	return schemas
}
