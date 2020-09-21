package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func ContactenosVCQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"contactenosvc_listar": {
			Type: graphql.NewList(types.ContactenosVCType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarContactenosVC(), nil
			},
		},
	}
	return schemas
}

func ContactenosVCMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"contactenosvc_crear": &graphql.Field{
			Type:        types.ContactenosVCType,
			Description: "Contáctenos ventas comerciales",
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
				"actividadEconomica": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Actividad económica del cliente",
				},
				"detalleSolicitud": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Detalle de solicitud o cotización.",
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
				actividadEconomica, _ := params.Args["actividadEconomica"].(string)
				detalleSolicitud, _ := params.Args["detalleSolicitud"].(string)

				return resolvers.CrearContactenosVC(nombre, apellido, cedula, correo, telefono, actividadEconomica, detalleSolicitud), nil
			},
		},
	}
	return schemas
}
