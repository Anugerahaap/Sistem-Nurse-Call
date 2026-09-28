package device

import (
	"backend/model/web"
	"backend/utils"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Service struct {
	Repo Repository
	Db   *sql.DB
}

func NewService(repo Repository, db *sql.DB) *Service {
	return &Service{Repo: repo, Db: db}
}

func (s *Service) GetAllDevice(ctx context.Context) ([]web.Device, error) {
	query := "select device_id,room_name,tanggal,waktu from devices;"
	dvc, err := s.Repo.GetDevices(ctx, query, s.Db)
	if err != nil {
		return nil, err
	}

	if len(dvc) == 0 {
		return nil, utils.NotFoundDevices
	}

	return utils.ConvertDomainDevicesIntoSlices(dvc), nil
}

func (s *Service) GetDeviceByID(ctx context.Context, deviceID string) (*web.Device, error) {

	query := "select  device_id,room_name,tanggal,waktu from devices where device_id = $1 "

	device, err := s.Repo.GetDevice(ctx, query, s.Db, deviceID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.NotFoundDevices
		default:
			return nil, err
		}
	}

	return utils.ConvertDomainDeviceIntoWeb(device), nil
}

func (s *Service) RegisterNewDevice(ctx context.Context, payload *web.DevicePayload) (string, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return "", err
	}

	defer utils.CommitOrRollback(tx)

	//make sure device-id belum terdaftar

	_, err = s.GetDeviceByID(ctx, payload.DeviceId)
	if err != nil {
		return "", err
	}

	query := "insert into devices (device_id,room_name) values($1,$2)"

	result, err := s.Repo.Execute(ctx, query, tx, payload.DeviceId, payload.RoomName)
	if err != nil {
		return "", err
	}

	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		log.Println(err)
		return "", fmt.Errorf("no rows affected ,message :%v", err)
	}

	return "", nil
}

// seharusnya ada mac addres sebagai unique id ,karna aku cuman ada 1 unit esp32 jadi aku pakai device-id langsung (ini bukan panduan menggunakan func ini hehe)
// func UpdateDevice(ctx context.Context, mac, oldID string, payload *web.DevicePayload) error {}
func (s *Service) UpdateDeviceID(ctx context.Context, OldID string, payload *web.DevicePayload) (string, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return "", err
	}

	defer utils.CommitOrRollback(tx)

	// cek untuk memastikan oldID(id yg akan diubah) sudah terdaftar
	oldDevice, err := s.GetDeviceByID(ctx, OldID)
	if err != nil {
		return "", err
	}
	// cek jika old id tdk ada perubahan dengan id baru maka return error
	if oldDevice.DeviceId == payload.DeviceId {
		return "", utils.DeviceSameParameter
	}

	// pastikan id yg mau diubah blm terdaftar
	if _, err := s.GetDeviceByID(ctx, payload.DeviceId); err == nil {
		return "", utils.DevicesIdAlrRegistered
	}

	query := "update devices set device_id = $1 ,tanggal = current_date,waktu = current_time where device_id = $2"

	result, err := s.Repo.Execute(ctx, query, tx, payload.DeviceId, OldID)
	if affected, err := result.RowsAffected(); err != nil || affected == 0 {
		log.Println(err)
		return "", fmt.Errorf("no rows affected ,message :%v", err)
	}
	return "", nil
}

func (s *Service) UpdateRoomName(ctx context.Context, payload *web.DevicePayload) (string, error) {

	if _, err := s.GetDeviceByID(ctx, payload.DeviceId); err != nil {
		return "", err
	}

	query := "update devices set room_name = $1 , tanggal = current_date , waktu = current_time where device_id = $2"
	_, err := s.Repo.Execute(ctx, query, s.Db, payload.RoomName, payload.DeviceId)
	if err != nil {
		return "", err
	}

	return payload.RoomName, nil
}
