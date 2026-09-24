package device

import (
	"backend/model/domain"
	"backend/model/web"
	"backend/utils"
	"context"
	"database/sql"
	"fmt"
)

type Service struct {
	Repo Repository
	Db   *sql.DB
}

func NewService(repo Repository, db *sql.DB) *Service {
	return &Service{Repo: repo, Db: db}
}

func (s *Service) GetDevices(ctx context.Context) ([]web.Device, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)

	dvc, err := s.Repo.GetAllDevice(ctx, tx)
	if err != nil {
		return nil, err
	}

	if len(*dvc) < 1 {
		return nil, utils.NotFoundDevices
	}

	return utils.ConvertDomainDevicesIntoSlices(*dvc), nil

}

func (s *Service) GetDeviceByID(ctx context.Context, DeviceID string) (*web.Device, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)

	Device, err := s.Repo.GetByDeviceId(ctx, DeviceID, tx)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.NotFoundDevices
		default:
			return nil, err
		}

	}

	return utils.ConvertDomainDeviceIntoWeb(Device), nil

}

func (s *Service) RegisterNewDevice(ctx context.Context, device *web.DevicePayload) error {
	tx, err := s.Db.Begin()
	if err != nil {
		fmt.Println("error:", err.Error())

		return err
	}
	defer utils.CommitOrRollback(tx)

	//make sure device tidak(belum terdaftar) terdaftar
	_, err = s.GetDeviceByID(ctx, device.DeviceId)
	if err == nil {

		return utils.DevicesIdAlrRegistered
	}

	return s.Repo.Create(ctx, &domain.Device{DeviceId: device.DeviceId, RoomName: device.RoomName}, tx)

}

func (s *Service) UpdateRoomName(ctx context.Context, device *web.DevicePayload) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	// make sure device-id terdaftar
	_, err = s.GetDeviceByID(ctx, device.DeviceId)
	if err != nil {
		return err
	}

	query := "update devices set room_name = $1 , tanggal = current_date , waktu = current_time where device_id = $2"

	return s.Repo.Update(ctx, query, tx, device.RoomName, device.DeviceId)

}

func (s *Service) UpdateDeviceID(ctx context.Context, oldDeviceID string, device *web.DevicePayload) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	// make sure device-id terdaftar
	oldDevice, err := s.GetDeviceByID(ctx, oldDeviceID)
	if err != nil {
		return err
	}

	if oldDevice.DeviceId == device.DeviceId {
		return utils.DeviceSameParameter
	}

	// check if device-id already exist

	if _, err := s.GetDeviceByID(ctx, device.DeviceId); err == nil {
		fmt.Println(err)
		return utils.DevicesIdAlrRegistered
	}

	query := "update devices set device_id = $1 ,tanggal = current_date,waktu = current_time where device_id = $2"

	return s.Repo.Update(ctx, query, tx, device.DeviceId, oldDeviceID)

}

func (s *Service) DeleteDevice(ctx context.Context, device *web.DevicePayload) (*web.Device, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)
	//make sure device sudah terdaftar
	dvc, err := s.GetDeviceByID(ctx, device.DeviceId)
	if err != nil {
		return nil, err
	}

	return dvc, s.Repo.Delete(ctx, device.DeviceId, tx)
}
