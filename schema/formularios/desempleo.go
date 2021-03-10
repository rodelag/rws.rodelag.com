package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func DesempleoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"desempleo_ver": {
			Type:        types.DesempleoType,
			Description: "Ver Desempleo",
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
					return resolvers.VerDesempleo(id), nil
				}
				return nil, nil
			},
		},
		"desempleo_busqueda": {
			Type:        graphql.NewList(types.DesempleoType),
			Description: "Búsqueda Desempleo",
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
					return resolvers.BusquedaDesempleo(busqueda), nil
				}
				return nil, nil
			},
		},
		"desempleo_listar": {
			Type: graphql.NewList(types.DesempleoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarDesempleo(), nil
			},
		},
	}
	return schemas
}

func DesempleoMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"desempleo_crear": &graphql.Field{
			Type:        types.DesempleoType,
			Description: "Creación de Desempleo",
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
				"edad": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Edad del cliente",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"telefono": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono del cliente",
				},
				"direccionDomicilio": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Dirección del cliente",
				},
				"nombreEmpresa": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Empresa donde labora el cliente",
				},
				"tiempoLaboral": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tiempo laboral del cliente",
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
				edad, _ := params.Args["edad"].(string)
				correo, _ := params.Args["correo"].(string)
				telefono, _ := params.Args["telefono"].(string)
				direccionDomicilio, _ := params.Args["direccionDomicilio"].(string)
				nombreEmpresa, _ := params.Args["nombreEmpresa"].(string)
				tiempoLaboral, _ := params.Args["tiempoLaboral"].(string)

				return resolvers.CrearDesempleo(nombre, apellido, cedula, edad, correo, telefono, direccionDomicilio, nombreEmpresa, tiempoLaboral), nil
			},
		},
		"desempleo_editar": &graphql.Field{
			Type:        types.ComprobantePagoType,
			Description: "Edición Desempleo",
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

				return resolvers.EditarDesempleo(id, estado), nil
			},
		},
	}
	return schemas
}
