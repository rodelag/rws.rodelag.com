package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func ComprobantePagoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"comprobantepago_listar": {
			Type: graphql.NewList(types.ComprobantePagoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarComprobantePago(), nil
			},
		},
	}
	return schemas
}

func ComprobantePagoMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"comprobantepago_crear": &graphql.Field{
			Type:        types.ComprobantePagoType,
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
				"comprobantePago": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Comprobante de Pago",
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
				comprobantePago, _ := params.Args["comprobantePago"].(string)
				telefono, _ := params.Args["telefono"].(string)

				return resolvers.CrearComprobantePago(nombre, apellido, cedula, correo, telefono, comprobantePago), nil
			},
		},
	}
	return schemas
}
