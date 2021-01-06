package tarjeta

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

type Cliente struct {
	RegistroCedula,
	RegistroCuenta,
	RegistroNombre,
	RegistroLimite,
	RegistroUltimoSaldo,
	RegistroFechaUltimaCompra,
	RegistroMontoUltimaCompra,
	RegistroFechaApertura,
	RegistroFechaCorte,
	RegistroPagarAntesDe,
	RegistroCorte,
	RegistroEstado,
	RegistroFechaUltimoPago,
	RegistroMontoUltimoPago,
	RegistroPagoMinimo,
	RegistroCorreo,
	RegistroTelefono,
	RegistroFechaNacimiento,
	RegistroSexo,
	RegistroSucursal,
	RegistroTipoCliente,
	RegistroFechaEstoCuenta,
	RegistroNumeroCuenta,
	RegistroLimiteCreditoEstadoCuenta,
	RegistroSaldoDisponible,
	RegistroFechaInicioTran,
	RegistroFechaFinTran,
	RegistroSaldoAnterior,
	RegistroPagoMinimoEstadoCuenta,
	RegistroSaldoCorriente,
	RegistroSaldo30Dias,
	RegistroSaldo60Dias,
	RegistroSaldo90Dias,
	RegistroSaldo120Dias,
	RegistroFechaFinPago string
	RegistroEstadoCuentaDetalle []EstadoCuentaDetalle
}

type EstadoCuentaDetalle struct {
	RegistroFechaTran,
	RegistroNumeroDocumento,
	RegistroDescripcionDocumento,
	RegistroDebito,
	RegistroCredito,
	RegistroSaldo string
}

func conexionTarjeta() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		viper.GetString("basedatos.mysql.rodelag.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Tarjeta) ", true, errMySQL)
	}
	return connMySQL
}

func ClienteTarjetaRodelag(cedula string) Cliente {
	connMySQL := conexionTarjeta()
	defer connMySQL.Close()

	cliente := Cliente{
		RegistroEstadoCuentaDetalle: func() []EstadoCuentaDetalle {
			rows, err := connMySQL.Query(consultaEstadoCuentaDetalle(cedula))
			utils.LogError("Problemas al leer registro del detalle del estado de cuenta: (Tarjeta) ", false, err)
			defer rows.Close()

			d, detalle := EstadoCuentaDetalle{}, []EstadoCuentaDetalle{}

			for rows.Next() {
				err := rows.Scan(
					&d.RegistroFechaTran,
					&d.RegistroNumeroDocumento,
					&d.RegistroDescripcionDocumento,
					&d.RegistroDebito,
					&d.RegistroCredito,
					&d.RegistroSaldo,
				)
				utils.LogError("Problemas al iterar los registro del detalle del estado de cuenta: (Tarjeta) ", false, err)
				detalle = append(detalle, EstadoCuentaDetalle{
					RegistroFechaTran:            d.RegistroFechaTran,
					RegistroNumeroDocumento:      d.RegistroNumeroDocumento,
					RegistroDescripcionDocumento: d.RegistroDescripcionDocumento,
					RegistroDebito:               d.RegistroDebito,
					RegistroCredito:              d.RegistroCredito,
					RegistroSaldo:                d.RegistroSaldo,
				})
			}
			return detalle
		}(),
	}

	err := connMySQL.QueryRow(consultaCliente(cedula)).Scan(
		&cliente.RegistroCedula,
		&cliente.RegistroCuenta,
		&cliente.RegistroNombre,
		&cliente.RegistroLimite,
		&cliente.RegistroUltimoSaldo,
		&cliente.RegistroFechaUltimaCompra,
		&cliente.RegistroMontoUltimaCompra,
		&cliente.RegistroFechaApertura,
		&cliente.RegistroFechaCorte,
		&cliente.RegistroPagarAntesDe,
		&cliente.RegistroCorte,
		&cliente.RegistroEstado,
		&cliente.RegistroFechaUltimoPago,
		&cliente.RegistroMontoUltimoPago,
		&cliente.RegistroPagoMinimo,
		&cliente.RegistroCorreo,
		&cliente.RegistroTelefono,
		&cliente.RegistroFechaNacimiento,
		&cliente.RegistroSexo,
		&cliente.RegistroSucursal,
		&cliente.RegistroTipoCliente,
		&cliente.RegistroFechaEstoCuenta,
		&cliente.RegistroNumeroCuenta,
		&cliente.RegistroLimiteCreditoEstadoCuenta,
		&cliente.RegistroSaldoDisponible,
		&cliente.RegistroFechaInicioTran,
		&cliente.RegistroFechaFinTran,
		&cliente.RegistroSaldoAnterior,
		&cliente.RegistroPagoMinimoEstadoCuenta,
		&cliente.RegistroSaldoCorriente,
		&cliente.RegistroSaldo30Dias,
		&cliente.RegistroSaldo60Dias,
		&cliente.RegistroSaldo90Dias,
		&cliente.RegistroSaldo120Dias,
		&cliente.RegistroFechaFinPago,
	)
	utils.LogError("Problemas al leer registro: (Tarjeta) ", false, err)

	return cliente
}

func consultaCliente(cedula string) string {
	consulta := `
		SELECT DISTINCT
			saldo.registroCedula,
			saldo.registroCuenta,
			saldo.registroNombre,
			saldo.registroLimite,
			saldo.registroUltimoSaldo,
			saldo.registroFechaUltimaCompra,
			saldo.registroMontoUltimaCompra,
			saldo.registroFechaApertura,
			saldo.registroFechaCorte,
			saldo.registroPagarAntesDe,
			saldo.registroCorte,
			saldo.registroEstado,
			saldo.registroFechaUltimoPago,
			saldo.registroMontoUltimoPago,
			saldo.registroPagoMinimo,
			saldo.registroCorreo,
			saldo.registroTelefono,
			saldo.registroFechaNacimiento,
			saldo.registroSexo,
			saldo.registroSucursal,
			saldo.registroTipoCliente,
			IFNULL(estadoCuenta.registroFechaEstoCuenta, '') AS registroFechaEstoCuenta,
			IFNULL(estadoCuenta.registroNumeroCuenta, '') AS registroNumeroCuenta,
			IFNULL(estadoCuenta.registroLimiteCredito, '') AS registroLimiteCredito,
			IFNULL(estadoCuenta.registroSaldoDisponible, '') AS registroSaldoDisponible,
			IFNULL(estadoCuenta.registroFechaInicioTran, '') AS registroFechaInicioTran,
			IFNULL(estadoCuenta.registroFechaFinTran, '') AS registroFechaFinTran,
			IFNULL(estadoCuenta.registroSaldoAnterior, '') AS registroSaldoAnterior,
			IFNULL(estadoCuenta.registroPagoMinimo, '') AS registroPagoMinimo,
			IFNULL(estadoCuenta.registroSaldoCorriente, '') AS registroSaldoCorriente,
			IFNULL(estadoCuenta.registroSaldo30Dias, '') AS registroSaldo30Dias,
			IFNULL(estadoCuenta.registroSaldo60Dias, '') AS registroSaldo60Dias,
			IFNULL(estadoCuenta.registroSaldo90Dias, '') AS registroSaldo90Dias,
			IFNULL(estadoCuenta.registroSaldo120Dias, '') AS registroSaldo120Dias,
			IFNULL(estadoCuenta.registroFechaFinPago, '') AS registroFechaFinPago
		FROM
			rodelag_tarjetarodelag.saldo AS saldo
			LEFT JOIN rodelag_tarjetarodelag.estadoCuenta AS estadoCuenta
				ON saldo.registroCedula = estadoCuenta.registroIdentificacion
		WHERE
			saldo.registroCedula = '%s'
	`
	return fmt.Sprintf(consulta, cedula)
}

func consultaEstadoCuentaDetalle(cedula string) string {
	consulta := `
		SELECT
			registroFechaTran,
			registroNumeroDocumento,
			registroDescripcionDocumento,
			registroDebito,
			registroCredito,
			registroSaldo
		FROM
			rodelag_tarjetarodelag.estadoCuenta
		WHERE
			registroIdentificacion = '%s'
	`
	return fmt.Sprintf(consulta, cedula)
}
