// Code generated by gowrap. DO NOT EDIT.
// template: ../../../logger/slog.gotmpl
// gowrap: http://github.com/hexdigest/gowrap

package middleware

import (
	"context"
	"crypto/x509"
	"log/slog"

	_sourceMigration "github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/shared/api"
)

// TargetEndpointWithSlog implements _sourceMigration.TargetEndpoint that is instrumented with slog logger
type TargetEndpointWithSlog struct {
	_log  *slog.Logger
	_base _sourceMigration.TargetEndpoint
}

// NewTargetEndpointWithSlog instruments an implementation of the _sourceMigration.TargetEndpoint with simple logging
func NewTargetEndpointWithSlog(base _sourceMigration.TargetEndpoint, log *slog.Logger) TargetEndpointWithSlog {
	return TargetEndpointWithSlog{
		_base: base,
		_log:  log,
	}
}

// Connect implements _sourceMigration.TargetEndpoint
func (_d TargetEndpointWithSlog) Connect(ctx context.Context) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
	).Debug("TargetEndpointWithSlog: calling Connect")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("TargetEndpointWithSlog: method Connect returned an error")
		} else {
			log.Debug("TargetEndpointWithSlog: method Connect finished")
		}
	}()
	return _d._base.Connect(ctx)
}

// DoBasicConnectivityCheck implements _sourceMigration.TargetEndpoint
func (_d TargetEndpointWithSlog) DoBasicConnectivityCheck() (e1 api.ExternalConnectivityStatus, cp1 *x509.Certificate) {
	_d._log.Debug("TargetEndpointWithSlog: calling DoBasicConnectivityCheck")
	defer func() {
		log := _d._log.With(
			slog.Any("e1", e1),
			slog.Any("cp1", cp1),
		)
		log.Debug("TargetEndpointWithSlog: method DoBasicConnectivityCheck finished")
	}()
	return _d._base.DoBasicConnectivityCheck()
}

// IsWaitingForOIDCTokens implements _sourceMigration.TargetEndpoint
func (_d TargetEndpointWithSlog) IsWaitingForOIDCTokens() (b1 bool) {
	_d._log.Debug("TargetEndpointWithSlog: calling IsWaitingForOIDCTokens")
	defer func() {
		log := _d._log.With(
			slog.Bool("b1", b1),
		)
		log.Debug("TargetEndpointWithSlog: method IsWaitingForOIDCTokens finished")
	}()
	return _d._base.IsWaitingForOIDCTokens()
}
