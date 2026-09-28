package device

import (
	"backend/model/domain"
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repository interface {
	GetDevices(ctx context.Context, query string, dbtx DBTX, args ...any) ([]domain.Device, error)
	GetDevice(ctx context.Context, query string, dbtx DBTX, args ...any) (*domain.Device, error)
	Execute(ctx context.Context, query string, dbtx DBTX, args ...any) (sql.Result, error)
}

type RepositoryImplementaion struct {
}

func NewRepository() Repository {
	return &RepositoryImplementaion{}
}

func (r *RepositoryImplementaion) GetDevices(ctx context.Context, query string, dbtx DBTX, args ...any) ([]domain.Device, error) {

	dvs := []domain.Device{}

	rows, err := dbtx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		dv := &domain.Device{}
		if err := rows.Scan(&dv.DeviceId, &dv.RoomName, &dv.Tanggal, &dv.Waktu); err != nil {
			return nil, err
		}

		dvs = append(dvs, *dv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return dvs, nil
}

func (r *RepositoryImplementaion) GetDevice(ctx context.Context, query string, dbtx DBTX, args ...any) (*domain.Device, error) {
	Device := &domain.Device{}
	err := dbtx.QueryRowContext(ctx, query, args...).Scan(&Device.DeviceId, &Device.RoomName, &Device.Tanggal, &Device.Waktu)
	if err != nil {
		return nil, err
	}
	return Device, nil
}

func (r *RepositoryImplementaion) Execute(ctx context.Context, query string, dbtx DBTX, args ...any) (sql.Result, error) {
	return dbtx.ExecContext(ctx, query, args...)
}
