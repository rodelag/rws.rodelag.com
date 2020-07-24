package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func PrevencionFraudeQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"prevencionfraude_listar": {
			Type: graphql.NewList(types.PrevencionFraudeType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarPrevencionFraude(), nil
			},
		},
	}
	return schemas
}

func PrevencionFraudeMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"prevencionfraude_crear": &graphql.Field{
			Type:        types.PrevencionFraudeType,
			Description: "Creación Prevención de Fraude",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"fechaNacimiento": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha de nacimiento del cliente",
				},
				"lugarResidencia": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Lugar de residencia",
				},
				"celular": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Celular del cliente",
				},
				"fotoCedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto de la cédula o pasaporte",
				},
				"fotoTarjeta": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto frontal de la tarjeta",
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
				fechaNacimiento, _ := params.Args["fechaNacimiento"].(string)
				lugarResidencia, _ := params.Args["lugarResidencia"].(string)
				celular, _ := params.Args["celular"].(string)
				fotoCedula, _ := params.Args["fotoCedula"].(string)
				fotoTarjeta, _ := params.Args["fotoTarjeta"].(string)

				return resolvers.CrearPrevencionFraude(nombre, apellido, fechaNacimiento, lugarResidencia, celular, fotoCedula, fotoTarjeta), nil
			},
		},
	}
	return schemas
}
