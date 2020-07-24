package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func EstadoCuentaQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"estadocuenta_listar": {
			Type: graphql.NewList(types.SolicitudEstadoCuentaType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarEstadoCuenta(), nil
			},
		},
	}
	return schemas
}

func EstadoCuentaMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"estadocuenta_crear": &graphql.Field{
			Type:        types.SolicitudEstadoCuentaType,
			Description: "Creación de Solicitud de Estado de Cuenta",
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
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
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
				cedula, _ := params.Args["cedula"].(string)

				return resolvers.CrearSolicitudEstadoCuenta(nombre, apellido, correo, telefono, cedula), nil
			},
		},
	}
	return schemas
}
