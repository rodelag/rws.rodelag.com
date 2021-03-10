package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func EstadoCuentaQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"estadocuenta_ver": {
			Type:        types.SolicitudEstadoCuentaType,
			Description: "Ver detalle del estado de cuenta",
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
					return resolvers.VerSolicitudEstadoCuenta(id), nil
				}
				return nil, nil
			},
		},
		"estadocuenta_busqueda": {
			Type:        graphql.NewList(types.SolicitudEstadoCuentaType),
			Description: "Búsqueda de estados de cuenta",
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
					return resolvers.BusquedaEstadoCuenta(busqueda), nil
				}
				return nil, nil
			},
		},
		"estadocuenta_listar": {
			Type:        graphql.NewList(types.SolicitudEstadoCuentaType),
			Description: "Listado de estados de cuenta",
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
		"estadocuenta_editar": &graphql.Field{
			Type:        types.SolicitudEstadoCuentaType,
			Description: "Edición de Solicitud de Estado de Cuenta",
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

				return resolvers.EditarEstadoCuenta(id, estado), nil
			},
		},
	}
	return schemas
}
