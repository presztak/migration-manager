// Code generated by gowrap. DO NOT EDIT.
// template: ../../../logger/slog.gotmpl
// gowrap: http://github.com/hexdigest/gowrap

package middleware

import (
	"context"
	"log/slog"

	_sourceMigration "github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/shared/api"
	"github.com/google/uuid"
)

// InstanceRepoWithSlog implements _sourceMigration.InstanceRepo that is instrumented with slog logger
type InstanceRepoWithSlog struct {
	_log  *slog.Logger
	_base _sourceMigration.InstanceRepo
}

// NewInstanceRepoWithSlog instruments an implementation of the _sourceMigration.InstanceRepo with simple logging
func NewInstanceRepoWithSlog(base _sourceMigration.InstanceRepo, log *slog.Logger) InstanceRepoWithSlog {
	return InstanceRepoWithSlog{
		_base: base,
		_log:  log,
	}
}

// Create implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) Create(ctx context.Context, instance _sourceMigration.Instance) (i1 int64, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("instance", instance),
	).Debug("InstanceRepoWithSlog: calling Create")
	defer func() {
		log := _d._log.With(
			slog.Int64("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method Create returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method Create finished")
		}
	}()
	return _d._base.Create(ctx, instance)
}

// CreateOverrides implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) CreateOverrides(ctx context.Context, overrides _sourceMigration.InstanceOverride) (i1 int64, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("overrides", overrides),
	).Debug("InstanceRepoWithSlog: calling CreateOverrides")
	defer func() {
		log := _d._log.With(
			slog.Int64("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method CreateOverrides returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method CreateOverrides finished")
		}
	}()
	return _d._base.CreateOverrides(ctx, overrides)
}

// DeleteByUUID implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) DeleteByUUID(ctx context.Context, id uuid.UUID) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("id", id),
	).Debug("InstanceRepoWithSlog: calling DeleteByUUID")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method DeleteByUUID returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method DeleteByUUID finished")
		}
	}()
	return _d._base.DeleteByUUID(ctx, id)
}

// DeleteOverridesByUUID implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) DeleteOverridesByUUID(ctx context.Context, id uuid.UUID) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("id", id),
	).Debug("InstanceRepoWithSlog: calling DeleteOverridesByUUID")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method DeleteOverridesByUUID returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method DeleteOverridesByUUID finished")
		}
	}()
	return _d._base.DeleteOverridesByUUID(ctx, id)
}

// GetAll implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetAll(ctx context.Context) (i1 _sourceMigration.Instances, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
	).Debug("InstanceRepoWithSlog: calling GetAll")
	defer func() {
		log := _d._log.With(
			slog.Any("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetAll returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetAll finished")
		}
	}()
	return _d._base.GetAll(ctx)
}

// GetAllByBatch implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetAllByBatch(ctx context.Context, batch string) (i1 _sourceMigration.Instances, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("batch", batch),
	).Debug("InstanceRepoWithSlog: calling GetAllByBatch")
	defer func() {
		log := _d._log.With(
			slog.Any("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetAllByBatch returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetAllByBatch finished")
		}
	}()
	return _d._base.GetAllByBatch(ctx, batch)
}

// GetAllByBatchAndState implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetAllByBatchAndState(ctx context.Context, batch string, status api.MigrationStatusType) (i1 _sourceMigration.Instances, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("batch", batch),
		slog.Any("status", status),
	).Debug("InstanceRepoWithSlog: calling GetAllByBatchAndState")
	defer func() {
		log := _d._log.With(
			slog.Any("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetAllByBatchAndState returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetAllByBatchAndState finished")
		}
	}()
	return _d._base.GetAllByBatchAndState(ctx, batch, status)
}

// GetAllBySource implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetAllBySource(ctx context.Context, source string) (i1 _sourceMigration.Instances, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("source", source),
	).Debug("InstanceRepoWithSlog: calling GetAllBySource")
	defer func() {
		log := _d._log.With(
			slog.Any("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetAllBySource returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetAllBySource finished")
		}
	}()
	return _d._base.GetAllBySource(ctx, source)
}

// GetAllByState implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetAllByState(ctx context.Context, status api.MigrationStatusType) (i1 _sourceMigration.Instances, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("status", status),
	).Debug("InstanceRepoWithSlog: calling GetAllByState")
	defer func() {
		log := _d._log.With(
			slog.Any("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetAllByState returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetAllByState finished")
		}
	}()
	return _d._base.GetAllByState(ctx, status)
}

// GetAllUUIDs implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetAllUUIDs(ctx context.Context) (ua1 []uuid.UUID, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
	).Debug("InstanceRepoWithSlog: calling GetAllUUIDs")
	defer func() {
		log := _d._log.With(
			slog.Any("ua1", ua1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetAllUUIDs returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetAllUUIDs finished")
		}
	}()
	return _d._base.GetAllUUIDs(ctx)
}

// GetAllUnassigned implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetAllUnassigned(ctx context.Context) (i1 _sourceMigration.Instances, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
	).Debug("InstanceRepoWithSlog: calling GetAllUnassigned")
	defer func() {
		log := _d._log.With(
			slog.Any("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetAllUnassigned returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetAllUnassigned finished")
		}
	}()
	return _d._base.GetAllUnassigned(ctx)
}

// GetByUUID implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetByUUID(ctx context.Context, id uuid.UUID) (ip1 *_sourceMigration.Instance, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("id", id),
	).Debug("InstanceRepoWithSlog: calling GetByUUID")
	defer func() {
		log := _d._log.With(
			slog.Any("ip1", ip1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetByUUID returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetByUUID finished")
		}
	}()
	return _d._base.GetByUUID(ctx, id)
}

// GetOverridesByUUID implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) GetOverridesByUUID(ctx context.Context, id uuid.UUID) (ip1 *_sourceMigration.InstanceOverride, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("id", id),
	).Debug("InstanceRepoWithSlog: calling GetOverridesByUUID")
	defer func() {
		log := _d._log.With(
			slog.Any("ip1", ip1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method GetOverridesByUUID returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method GetOverridesByUUID finished")
		}
	}()
	return _d._base.GetOverridesByUUID(ctx, id)
}

// Update implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) Update(ctx context.Context, instance _sourceMigration.Instance) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("instance", instance),
	).Debug("InstanceRepoWithSlog: calling Update")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method Update returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method Update finished")
		}
	}()
	return _d._base.Update(ctx, instance)
}

// UpdateOverrides implements _sourceMigration.InstanceRepo
func (_d InstanceRepoWithSlog) UpdateOverrides(ctx context.Context, overrides _sourceMigration.InstanceOverride) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("overrides", overrides),
	).Debug("InstanceRepoWithSlog: calling UpdateOverrides")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("InstanceRepoWithSlog: method UpdateOverrides returned an error")
		} else {
			log.Debug("InstanceRepoWithSlog: method UpdateOverrides finished")
		}
	}()
	return _d._base.UpdateOverrides(ctx, overrides)
}
