package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func EsiaaQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"esiaa_listar": {
			Type: graphql.NewList(types.EsiaaType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarEsiaa(), nil
			},
		},
	}
	return schemas
}

func EsiaaMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"esiaa_crear": &graphql.Field{
			Type:        types.EsiaaType,
			Description: "Contáctenos ventas comerciales",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del lciente",
				},
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"calificacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Calificación del cliente",
				},
				"atencion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Atención del cliente",
				},
				"resolverInstalacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Instalación resuelta",
				},
				"tiempoRazonable": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tiempo razonable",
				},
				"recomendacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Recomendación",
				},
				"calificacionManera": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Calificación del cliente en general",
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
				calificacion, _ := params.Args["calificacion"].(string)
				atencion, _ := params.Args["atencion"].(string)
				resolverInstalacion, _ := params.Args["resolverInstalacion"].(string)
				tiempoRazonable, _ := params.Args["tiempoRazonable"].(string)
				recomendacion, _ := params.Args["recomendacion"].(string)
				calificacionManera, _ := params.Args["calificacionManera"].(string)

				return resolvers.CrearEsiaa(nombre, apellido, cedula, correo, calificacion, atencion, resolverInstalacion, tiempoRazonable, recomendacion, calificacionManera), nil
			},
		},
	}
	return schemas
}