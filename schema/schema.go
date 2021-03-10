package schema

import (
	"github.com/graphql-go/graphql"
	"rws/schema/central"
	"rws/schema/conteo"
	"rws/schema/formularios"
	"rws/schema/inventario"
	"rws/schema/inventario_actualizado"
	"rws/schema/proveedores"
	"rws/schema/tarjeta"
	"rws/schema/ventas_detalles"
)

var Schema graphql.Schema

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//Formulario de Pago por ACH
			"pagoach_ver":    formularios.PagoACHQuery()["pagoach_ver"],
			"pagoach_listar": formularios.PagoACHQuery()["pagoach_listar"],

			//Formulario de Afiliación de Sos
			"sos_ver":      formularios.SosQuery()["sos_ver"],
			"sos_listar":   formularios.SosQuery()["sos_listar"],
			"sos_busqueda": formularios.SosQuery()["sos_busqueda"],

			//Formulario de Solicitud de Estado de Cuenta
			"estadocuenta_ver":      formularios.EstadoCuentaQuery()["estadocuenta_ver"],
			"estadocuenta_listar":   formularios.EstadoCuentaQuery()["estadocuenta_listar"],
			"estadocuenta_busqueda": formularios.EstadoCuentaQuery()["estadocuenta_busqueda"],

			//Formulario de Desempleo
			"desempleo_ver":      formularios.DesempleoQuery()["desempleo_ver"],
			"desempleo_listar":   formularios.DesempleoQuery()["desempleo_listar"],
			"desempleo_busqueda": formularios.DesempleoQuery()["desempleo_busqueda"],

			//Formulario de Arreglo de Pago
			"arreglopago_ver":      formularios.ArregloPagoQuery()["arreglopago_ver"],
			"arreglopago_listar":   formularios.ArregloPagoQuery()["arreglopago_listar"],
			"arreglopago_busqueda": formularios.ArregloPagoQuery()["arreglopago_busqueda"],

			//Formulario de Comprobante de Pago
			"comprobantepago_ver":      formularios.ComprobantePagoQuery()["comprobantepago_ver"],
			"comprobantepago_listar":   formularios.ComprobantePagoQuery()["comprobantepago_listar"],
			"comprobantepago_busqueda": formularios.ComprobantePagoQuery()["comprobantepago_busqueda"],

			//Formulario de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19
			"solicitudaplazamientopago_ver":    formularios.SolicitudAplazamientoPagoQuery()["solicitudaplazamientopago_ver"],
			"solicitudaplazamientopago_listar": formularios.SolicitudAplazamientoPagoQuery()["solicitudaplazamientopago_listar"],

			//Formulario de Solicitud Tarjeta Rodelag
			"solicitudtarjetarodelag_ver":    formularios.SolicitudTarjetaRodelagQuery()["solicitudtarjetarodelag_ver"],
			"solicitudtarjetarodelag_listar": formularios.SolicitudTarjetaRodelagQuery()["solicitudtarjetarodelag_listar"],

			//Formulario de Notificación Descuento
			"notificaciondescuento_ver":    formularios.NotificacionDescuentoQuery()["notificaciondescuento_ver"],
			"notificaciondescuento_listar": formularios.NotificacionDescuentoQuery()["notificaciondescuento_listar"],

			//Encuesta de Satisfacción Instalación AA
			"esiaa_ver":    formularios.EsiaaQuery()["esiaa_ver"],
			"esiaa_listar": formularios.EsiaaQuery()["esiaa_listar"],

			//Formulario de Prevención de Fraude
			"prevencionfraude_ver":    formularios.PrevencionFraudeQuery()["prevencionfraude_ver"],
			"prevencionfraude_listar": formularios.PrevencionFraudeQuery()["prevencionfraude_listar"],

			//Formulario de Reclamo
			"reclamo_ver":    formularios.ReclamoQuery()["reclamo_ver"],
			"reclamo_listar": formularios.ReclamoQuery()["reclamo_listar"],

			//Formulario de Contáctenos
			"contactenos_ver":    formularios.ContactenosQuery()["contactenos_ver"],
			"contactenos_listar": formularios.ContactenosQuery()["contactenos_listar"],

			//Formulario de Contáctenos de ventas comerciales
			"contactenosvc_ver":    formularios.ContactenosVCQuery()["contactenosvc_ver"],
			"contactenosvc_listar": formularios.ContactenosVCQuery()["contactenosvc_listar"],

			//Inventario de los proveedores
			"proveedor_inventario_listar": proveedores.ProveedorInventarioQuery()["proveedor_inventario_listar"],

			//Ventas de los proveedores
			"proveedor_ventas_listar": proveedores.ProveedorVentasQuery()["proveedor_ventas_listar"],

			//Conteo de clientes en las tiendas
			"conteo_listar": conteo.ConteoQuery()["conteo_listar"],

			//Cliente Tarjeta Rodelag
			"cliente_tarjeta_rodelag":               tarjeta.TarjetaRodelagQuery()["cliente"],
			"cliente_tarjeta_rodelag_cedula_cuenta": tarjeta.TarjetaRodelagQuery()["cliente_cedula_cuenta"],

			//Central teléfonica
			"central_telefonica": central.CentralTelefonicaQuery()["central"],

			//Ventas de Rodelag
			"ventas_listar": ventas_detalles.VentasDetallesQuery()["ventas_listar"],

			//Inventario de Rodelag
			"inventario_listar": inventario.InventarioQuery()["inventario_listar"],

			//Inventario Actualizado de Rodelag
			"inventario_actualizado_detalle_listar":       inventario_actualizado.InventarioActualizadoQuery()["inventario_actualizado_detalle_listar"],
			"inventario_actualizado_tiendas_listar":       inventario_actualizado.InventarioActualizadoQuery()["inventario_actualizado_tiendas_listar"],
			"inventario_actualizado_departamentos_listar": inventario_actualizado.InventarioActualizadoQuery()["inventario_actualizado_departamentos_listar"],
			"inventario_actualizado_categorias_listar":    inventario_actualizado.InventarioActualizadoQuery()["inventario_actualizado_categorias_listar"],
			"inventario_actualizado_producto_listar":      inventario_actualizado.InventarioActualizadoQuery()["inventario_actualizado_producto_listar"],
		},
	})

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//Formulario de Pago por ACH
			"pagoach_crear":            formularios.PagoACHMutation()["pagoach_crear"],
			"pagoach_crear_comentario": formularios.PagoACHMutation()["pagoach_crear_comentario"],

			//Formulario de Afiliación de Sos
			"sos_crear":  formularios.SosMutation()["sos_crear"],
			"sos_editar": formularios.SosMutation()["sos_editar"],

			//Formulario de Desempleo
			"desempleo_crear":  formularios.DesempleoMutation()["desempleo_crear"],
			"desempleo_editar": formularios.DesempleoMutation()["desempleo_editar"],

			//Formulario de Arreglo de Pago
			"arreglopago_crear":  formularios.ArregloPagoMutation()["arreglopago_crear"],
			"arreglopago_editar": formularios.ArregloPagoMutation()["arreglopago_editar"],

			//Formulario de Solicitud de Estado de Cuenta
			"estadocuenta_crear":  formularios.EstadoCuentaMutation()["estadocuenta_crear"],
			"estadocuenta_editar": formularios.EstadoCuentaMutation()["estadocuenta_editar"],

			//Formulario de Comprobante de Pago
			"comprobantepago_crear":  formularios.ComprobantePagoMutation()["comprobantepago_crear"],
			"comprobantepago_editar": formularios.ComprobantePagoMutation()["comprobantepago_editar"],

			//Formulario de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19
			"solicitudaplazamientopago_crear":            formularios.SolicitudAplazamientoPagoMutation()["solicitudaplazamientopago_crear"],
			"solicitudaplazamientopago_modificar":        formularios.SolicitudAplazamientoPagoMutation()["solicitudaplazamientopago_modificar"],
			"solicitudaplazamientopago_crear_comentario": formularios.SolicitudAplazamientoPagoMutation()["solicitudaplazamientopago_crear_comentario"],

			//Formulario de Solicitud Tarjeta Rodelag
			"solicitudtarjetarodelag_crear":            formularios.SolicitudTarjetaRodelagMutation()["solicitudtarjetarodelag_crear"],
			"solicitudtarjetarodelag_crear_comentario": formularios.SolicitudTarjetaRodelagMutation()["solicitudtarjetarodelag_crear_comentario"],

			//Formulario de Notificación Descuento
			"notificaciondescuento_crear":            formularios.NotificacionDescuentoMutation()["notificaciondescuento_crear"],
			"notificaciondescuento_crear_comentario": formularios.NotificacionDescuentoMutation()["notificaciondescuento_crear_comentario"],

			//Encuesta de Satisfacción Instalación AA
			"esiaa_crear":            formularios.EsiaaMutation()["esiaa_crear"],
			"esiaa_crear_comentario": formularios.EsiaaMutation()["esiaa_crear_comentario"],

			//Formulario de Prevención de Fraude
			"prevencionfraude_crear":            formularios.PrevencionFraudeMutation()["prevencionfraude_crear"],
			"prevencionfraude_crear_comentario": formularios.PrevencionFraudeMutation()["prevencionfraude_crear_comentario"],

			//Formulario de Reclamo
			"reclamo_crear":            formularios.ReclamoMutation()["reclamo_crear"],
			"reclamo_crear_comentario": formularios.ReclamoMutation()["reclamo_crear_comentario"],

			//Formulario de Contáctenos
			"contactenos_crear":            formularios.ContactenosMutation()["contactenos_crear"],
			"contactenos_crear_comentario": formularios.ContactenosMutation()["contactenos_crear_comentario"],

			//Formulario de Contáctenos de ventas comerciales
			"contactenosvc_crear":            formularios.ContactenosVCMutation()["contactenosvc_crear"],
			"contactenosvc_crear_comentario": formularios.ContactenosVCMutation()["contactenosvc_crear_comentario"],
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}
