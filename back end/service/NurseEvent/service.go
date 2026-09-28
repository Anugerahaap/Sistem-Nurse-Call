package nurseevent

import (
	"backend/model/web"
	"backend/service/device"
	"backend/utils"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Service struct {
	Repo           Repository
	Db             *sql.DB
	deviceServices device.Service
}

func NewService(repo Repository, db *sql.DB, deviceServices device.Service) *Service {
	return &Service{Repo: repo, Db: db, deviceServices: deviceServices}
}

// mendapatkan semua event terakhir dari device-id yang terdaftar
func (s *Service) GetLastEventByID(ctx context.Context) ([]web.NurseEvent, error) {

	query := "select distinct on (device_id) device_id,event,tanggal,waktu from nurse_events order by device_id ,tanggal desc,waktu desc"

	dme, err := s.Repo.GetNurseEvents(ctx, query, s.Db)
	if err != nil {
		return nil, err
	}

	if len(dme) == 0 {
		return nil, utils.NotFoundNurse
	}

	return utils.ConvertDomainNurseEventsIntoSlices(dme), nil
}

// mendapatkan event terakhir berdasarkan device-id
func (s *Service) GetLatestEvent(ctx context.Context, deviceID string) (*web.NurseEvent, error) {

	query := "select distinct on (device_id) device_id,event,tanggal,waktu from nurse_events where device_id=$1 order by device_id ,tanggal desc,waktu desc"

	wbe, err := s.Repo.GetNurseEvent(ctx, query, s.Db, deviceID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.NotFoundNurse
		default:
			log.Println(err)
			return nil, err
		}

	}

	return utils.ConvertDomainNurseEventIntoWeb(wbe), nil
}

func (s *Service) GetAllEvents(ctx context.Context) ([]web.NurseEvent, error) {

	query := "select device_id, event,tanggal,waktu from nurse_events"

	dme, err := s.Repo.GetNurseEvents(ctx, query, s.Db)
	if err != nil {
		return nil, err
	}

	if len(dme) == 0 {
		return nil, utils.NotFoundNurse
	}

	return utils.ConvertDomainNurseEventsIntoSlices(dme), nil
}

func (s *Service) CreateNewEvent(ctx context.Context, payload *web.NurseEventCreatePayload) error {

	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	//pastikan device-id terdaftar
	if _, err := s.deviceServices.GetDeviceByID(ctx, payload.DeviceId); err != nil {
		return err
	}

	query := "insert into nurse_events (device_id,event,tanggal,waktu) values($1,$2,current_date,current_time)"
	result, err := s.Repo.Execute(ctx, query, tx, payload.DeviceId, payload.Event)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows affected ,message :%v", err)
	}
	return nil
}
