package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func SosQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"sos_ver": {
			Type:        types.SosType,
			Description: "Ver Afiliación de Sos",
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
					return resolvers.VerSos(id), nil
				}
				return nil, nil
			},
		},
		"sos_busqueda": {
			Type:        graphql.NewList(types.SosType),
			Description: "Búsqueda de Afiliación de Sos",
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
					return resolvers.BusquedaSos(busqueda), nil
				}
				return nil, nil
			},
		},
		"sos_listar": {
			Type:        graphql.NewList(types.SosType),
			Description: "Listar Afiliación de Sos",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarSos(), nil
			},
		},
	}
	return schemas
}

func SosMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"sos_crear": &graphql.Field{
			Type:        types.SosType,
			Description: "Creación de Afiliación de Sos",
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

				return resolvers.CrearSos(nombre, apellido, cedula, correo, telefono), nil
			},
		},
		"sos_editar": &graphql.Field{
			Type:        types.SosType,
			Description: "Edición de Afiliación de Sos",
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

				return resolvers.EditarSos(id, estado), nil
			},
		},
	}
	return schemas
}
