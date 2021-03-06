package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func ComprobantePagoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"comprobantepago_ver": {
			Type:        types.ComprobantePagoType,
			Description: "Ver pago ACH",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "ID del registro",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				if id, ok := p.Args["id"].(int); ok {
					return resolvers.VerComprobantePago(id), nil
				}
				return nil, nil
			},
		},
		"comprobantepago_busqueda": {
			Type:        graphql.NewList(types.ComprobantePagoType),
			Description: "Búsqueda pago ACH",
			Args: graphql.FieldConfigArgument{
				"busqueda": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Palabra a buscar",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				if busqueda, ok := p.Args["busqueda"].(string); ok {
					return resolvers.BusquedaComprobantePago(busqueda), nil
				}
				return nil, nil
			},
		},
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
		"comprobantepago_editar": &graphql.Field{
			Type:        types.ComprobantePagoType,
			Description: "Edición de los comentarios del comprobante de pago",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "ID del registro",
				},
				"estado": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Estado del registro",
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
				estado, _ := params.Args["estado"].(string)
				id, _ := params.Args["id"].(int)

				return resolvers.EditarComprobantePago(id, estado), nil
			},
		},
	}
	return schemas
}
