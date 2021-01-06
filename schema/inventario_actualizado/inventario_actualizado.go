package inventario_actualizado

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/inventario_actualizado"
	types "rws/types/inventario_actualizado"
)

func InventarioActualizadoQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"inventario_actualizado_detalle_listar": {
			Type:        graphql.NewList(types.InventarioActualizadoType),
			Description: "Listar el inventario actualizado.",
			Args: graphql.FieldConfigArgument{
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal",
				},
				"departamento": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Departamento",
				},
				"categoria": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Categoria",
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
				sucursal, _ := p.Args["sucursal"].(string)
				departamento, _ := p.Args["departamento"].(string)
				categoria, _ := p.Args["categoria"].(string)

				return resolvers.ListarInventarioActualizado(sucursal, departamento, categoria), nil
			},
		},
		"inventario_actualizado_tiendas_listar": {
			Type:        graphql.NewList(types.InventarioTiendasType),
			Description: "Listar tidenas con inventario actualizado.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarInventarioActualizadoTiendas(), nil
			},
		},
		"inventario_actualizado_departamentos_listar": {
			Type:        graphql.NewList(types.InventarioDepartamentosType),
			Description: "Listar departamentos con inventario actualizado.",
			Args: graphql.FieldConfigArgument{
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal",
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
				sucursal, _ := p.Args["sucursal"].(string)
				return resolvers.ListarInventarioActualizadoDepartamento(sucursal), nil
			},
		},
		"inventario_actualizado_categorias_listar": {
			Type:        graphql.NewList(types.InventarioCategoriasType),
			Description: "Listar categorías con inventario actualizado.",
			Args: graphql.FieldConfigArgument{
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal",
				},
				"departamento": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Departamento",
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
				sucursal, _ := p.Args["sucursal"].(string)
				departamento, _ := p.Args["departamento"].(string)

				return resolvers.ListarInventarioActualizadoCategorias(sucursal, departamento), nil
			},
		},
		"inventario_actualizado_producto_listar": {
			Type:        types.InventarioProductosType,
			Description: "Listar Producto con inventario actualizado.",
			Args: graphql.FieldConfigArgument{
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal",
				},
				"codigo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Codigo",
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
				sucursal, _ := p.Args["sucursal"].(string)
				codigo, _ := p.Args["codigo"].(string)

				return resolvers.ListarInventarioActualizadoProducto(codigo, sucursal), nil
			},
		},
	}
	return schemas
}
