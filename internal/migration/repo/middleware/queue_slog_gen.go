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

// QueueRepoWithSlog implements _sourceMigration.QueueRepo that is instrumented with slog logger
type QueueRepoWithSlog struct {
	_log  *slog.Logger
	_base _sourceMigration.QueueRepo
}

// NewQueueRepoWithSlog instruments an implementation of the _sourceMigration.QueueRepo with simple logging
func NewQueueRepoWithSlog(base _sourceMigration.QueueRepo, log *slog.Logger) QueueRepoWithSlog {
	return QueueRepoWithSlog{
		_base: base,
		_log:  log,
	}
}

// Create implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) Create(ctx context.Context, queue _sourceMigration.QueueEntry) (i1 int64, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("queue", queue),
	).Debug("QueueRepoWithSlog: calling Create")
	defer func() {
		log := _d._log.With(
			slog.Int64("i1", i1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method Create returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method Create finished")
		}
	}()
	return _d._base.Create(ctx, queue)
}

// DeleteAllByBatch implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) DeleteAllByBatch(ctx context.Context, batch string) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("batch", batch),
	).Debug("QueueRepoWithSlog: calling DeleteAllByBatch")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method DeleteAllByBatch returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method DeleteAllByBatch finished")
		}
	}()
	return _d._base.DeleteAllByBatch(ctx, batch)
}

// DeleteByUUID implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) DeleteByUUID(ctx context.Context, id uuid.UUID) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("id", id),
	).Debug("QueueRepoWithSlog: calling DeleteByUUID")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method DeleteByUUID returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method DeleteByUUID finished")
		}
	}()
	return _d._base.DeleteByUUID(ctx, id)
}

// GetAll implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) GetAll(ctx context.Context) (q1 _sourceMigration.QueueEntries, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
	).Debug("QueueRepoWithSlog: calling GetAll")
	defer func() {
		log := _d._log.With(
			slog.Any("q1", q1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method GetAll returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method GetAll finished")
		}
	}()
	return _d._base.GetAll(ctx)
}

// GetAllByBatch implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) GetAllByBatch(ctx context.Context, batch string) (q1 _sourceMigration.QueueEntries, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("batch", batch),
	).Debug("QueueRepoWithSlog: calling GetAllByBatch")
	defer func() {
		log := _d._log.With(
			slog.Any("q1", q1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method GetAllByBatch returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method GetAllByBatch finished")
		}
	}()
	return _d._base.GetAllByBatch(ctx, batch)
}

// GetAllByBatchAndState implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) GetAllByBatchAndState(ctx context.Context, batch string, statuses ...api.MigrationStatusType) (q1 _sourceMigration.QueueEntries, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("batch", batch),
		slog.Any("statuses", statuses),
	).Debug("QueueRepoWithSlog: calling GetAllByBatchAndState")
	defer func() {
		log := _d._log.With(
			slog.Any("q1", q1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method GetAllByBatchAndState returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method GetAllByBatchAndState finished")
		}
	}()
	return _d._base.GetAllByBatchAndState(ctx, batch, statuses...)
}

// GetAllByState implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) GetAllByState(ctx context.Context, status ...api.MigrationStatusType) (q1 _sourceMigration.QueueEntries, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("status", status),
	).Debug("QueueRepoWithSlog: calling GetAllByState")
	defer func() {
		log := _d._log.With(
			slog.Any("q1", q1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method GetAllByState returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method GetAllByState finished")
		}
	}()
	return _d._base.GetAllByState(ctx, status...)
}

// GetAllNeedingImport implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) GetAllNeedingImport(ctx context.Context, batch string, needsDiskImport bool) (q1 _sourceMigration.QueueEntries, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.String("batch", batch),
		slog.Bool("needsDiskImport", needsDiskImport),
	).Debug("QueueRepoWithSlog: calling GetAllNeedingImport")
	defer func() {
		log := _d._log.With(
			slog.Any("q1", q1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method GetAllNeedingImport returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method GetAllNeedingImport finished")
		}
	}()
	return _d._base.GetAllNeedingImport(ctx, batch, needsDiskImport)
}

// GetByInstanceUUID implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) GetByInstanceUUID(ctx context.Context, id uuid.UUID) (qp1 *_sourceMigration.QueueEntry, err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("id", id),
	).Debug("QueueRepoWithSlog: calling GetByInstanceUUID")
	defer func() {
		log := _d._log.With(
			slog.Any("qp1", qp1),
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method GetByInstanceUUID returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method GetByInstanceUUID finished")
		}
	}()
	return _d._base.GetByInstanceUUID(ctx, id)
}

// Update implements _sourceMigration.QueueRepo
func (_d QueueRepoWithSlog) Update(ctx context.Context, entry _sourceMigration.QueueEntry) (err error) {
	_d._log.With(
		slog.Any("ctx", ctx),
		slog.Any("entry", entry),
	).Debug("QueueRepoWithSlog: calling Update")
	defer func() {
		log := _d._log.With(
			slog.Any("err", err),
		)
		if err != nil {
			log.Error("QueueRepoWithSlog: method Update returned an error")
		} else {
			log.Debug("QueueRepoWithSlog: method Update finished")
		}
	}()
	return _d._base.Update(ctx, entry)
}
