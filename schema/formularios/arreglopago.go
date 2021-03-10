package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func ArregloPagoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"arreglopago_ver": {
			Type:        types.ArregloPagoType,
			Description: "Ver detalle del arreglo de pago",
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
					return resolvers.VerArregloPago(id), nil
				}
				return nil, nil
			},
		},
		"arreglopago_busqueda": {
			Type:        graphql.NewList(types.ArregloPagoType),
			Description: "Búsqueda de arreglo de pago",
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
					return resolvers.BusquedaArregloPago(busqueda), nil
				}
				return nil, nil
			},
		},
		"arreglopago_listar": {
			Type:        graphql.NewList(types.ArregloPagoType),
			Description: "Listado de arreglos de pago",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarArregloPago(), nil
			},
		},
	}
	return schemas
}

func ArregloPagoMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"arreglopago_crear": &graphql.Field{
			Type:        types.ArregloPagoType,
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
				"celular": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Celular del cliente",
				},
				"direccion_domicilio": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Dirección del cliente",
				},
				"telefono_trabajo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono de trabajo del cliente",
				},
				"lugar_trabajo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del lugar de trabajo del cliente",
				},
				"sector": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sector de cliente",
				},
				"motivo_arreglo_pago": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Motivo del arreglo de pago",
				},
				"foto_cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto de la cédula",
				},
				"comprobante_abono": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto del comprobante de abono",
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
				celular, _ := params.Args["celular"].(string)
				direccionDomicilio, _ := params.Args["direccion_domicilio"].(string)
				telefonoTrabajo, _ := params.Args["telefono_trabajo"].(string)
				lugarTrabajo, _ := params.Args["lugar_trabajo"].(string)
				sector, _ := params.Args["sector"].(string)
				motivoArregloPago, _ := params.Args["motivo_arreglo_pago"].(string)
				fotoCedula, _ := params.Args["foto_cedula"].(string)
				comprobanteAbono, _ := params.Args["comprobante_abono"].(string)

				return resolvers.CrearArregloPago(nombre, apellido, cedula, correo, telefono, celular, direccionDomicilio, telefonoTrabajo, lugarTrabajo, sector, motivoArregloPago, fotoCedula, comprobanteAbono), nil
			},
		},
		"arreglopago_editar": &graphql.Field{
			Type:        types.ArregloPagoType,
			Description: "Edición de Arreglo de Pago",
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

				return resolvers.EditarArregloPago(id, estado), nil
			},
		},
	}
	return schemas
}
