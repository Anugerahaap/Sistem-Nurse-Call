package nurseevent

import (
	"backend/model/domain"
	"backend/model/web"
	"backend/service/device"
	"backend/utils"
	"context"
	"database/sql"
)

type Service struct {
	Repo           Repository
	Db             *sql.DB
	deviceServices device.Service
}

func NewService(repo Repository, db *sql.DB, deviceServices device.Service) *Service {
	return &Service{Repo: repo, Db: db, deviceServices: deviceServices}
}

func (s *Service) GetNurseEvents(ctx context.Context) ([]web.NurseEvent, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)
	nurseEvents, err := s.Repo.GetNurseEvents(ctx, tx)
	if err != nil {
		return nil, err
	}

	if len(nurseEvents) < 1 {
		return nil, utils.NotFoundNurse
	}

	return utils.ConvertDomainNurseEventsIntoSlices(nurseEvents), nil

}

func (s *Service) GetByDeviceID(ctx context.Context, deviceID string) ([]web.NurseEvent, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)

	nurseEvents, err := s.Repo.GetByDeviceID(ctx, deviceID, tx)
	if err != nil {
		return nil, err
	}

	if len(nurseEvents) < 1 {
		return nil, utils.NotFoundNurse
	}

	return utils.ConvertDomainNurseEventsIntoSlices(nurseEvents), nil
}

func (s *Service) CreateNurseEvent(ctx context.Context, ns *web.NurseEventCreatePayload) (string, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return "", err
	}
	defer utils.CommitOrRollback(tx)

	nurseEvent := domain.NurseEvent{
		DeviceId: ns.DeviceId,
		Event:    ns.Event,
	}

	// check if device exists
	if _, err := s.deviceServices.GetDeviceByID(ctx, nurseEvent.DeviceId); err != nil {
		return "", err
	}

	deviceID, err := s.Repo.Create(ctx, nurseEvent, tx)
	if err != nil {
		return "", err
	}

	return deviceID, nil
}

func (s *Service) GetLastEvents(ctx context.Context) ([]web.NurseEvent, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)

	nurseEvents, err := s.Repo.GetLastEvents(ctx, tx)
	if err != nil {
		return nil, err
	}

	if len(nurseEvents) < 1 {
		return nil, utils.NotFoundNurse
	}

	return utils.ConvertDomainNurseEventsIntoSlices(nurseEvents), nil
}
