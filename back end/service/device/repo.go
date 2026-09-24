package device

import (
	"backend/model/domain"
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAllDevice(ctx context.Context, tx *sql.Tx) (*[]domain.Device, error)
	GetByDeviceId(ctx context.Context, deviceId string, tx *sql.Tx) (*domain.Device, error)
	Create(ctx context.Context, dvs *domain.Device, tx *sql.Tx) error
	Update(ctx context.Context, query string, tx *sql.Tx, args ...any) error
	Delete(ctx context.Context, deviceId string, tx *sql.Tx) error
}

type RepositoryImplementaion struct {
}

func NewRepository() Repository {
	return &RepositoryImplementaion{}
}

func (r *RepositoryImplementaion) GetAllDevice(ctx context.Context, tx *sql.Tx) (*[]domain.Device, error) {

	dvs := []domain.Device{}

	rows, err := tx.QueryContext(ctx, "select device_id,room_name,tanggal,waktu from devices;")
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
	return &dvs, nil
}

func (r *RepositoryImplementaion) GetByDeviceId(ctx context.Context, deviceId string, tx *sql.Tx) (*domain.Device, error) {
	Device := &domain.Device{}
	err := tx.QueryRowContext(ctx, "select  device_id,room_name,tanggal,waktu from devices where device_id = $1 ", deviceId).Scan(&Device.DeviceId, &Device.RoomName, &Device.Tanggal, &Device.Waktu)
	if err != nil {
		return nil, err
	}

	return Device, nil
}

func (r *RepositoryImplementaion) Create(ctx context.Context, dvs *domain.Device, tx *sql.Tx) error {

	result, err := tx.ExecContext(ctx, "insert into devices (device_id,room_name) values($1,$2)", dvs.DeviceId, dvs.RoomName)
	if err != nil {
		fmt.Println("error:", err.Error())
		return err
	}

	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		return fmt.Errorf("no rows affected ,message :%v", err)
	}

	return nil
}

func (r *RepositoryImplementaion) Update(ctx context.Context, query string, tx *sql.Tx, args ...any) error {

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		return fmt.Errorf("no rows affected ,message :%v", err)
	}

	return nil
}

func (r *RepositoryImplementaion) Delete(ctx context.Context, deviceId string, tx *sql.Tx) error {
	result, err := tx.ExecContext(ctx, "delete from devices where devices_id = $1", deviceId)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		return fmt.Errorf("no rows affected ,message :%v", err)
	}
	return nil
}
