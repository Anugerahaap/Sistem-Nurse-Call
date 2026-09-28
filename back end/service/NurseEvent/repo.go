package nurseevent

import (
	"backend/model/domain"
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type Repository interface {
	// Create(ctx context.Context, ns domain.NurseEvent, tx *sql.Tx) (string, error)
	// GetByDeviceID(ctx context.Context, deviceID string, tx *sql.Tx) ([]domain.NurseEvent, error)
	// GetNurseEvents(ctx context.Context, tx *sql.Tx) ([]domain.NurseEvent, error)
	// GetLastEvents(ctx context.Context, tx *sql.Tx) ([]domain.NurseEvent, error)
	GetNurseEvents(ctx context.Context, query string, dbtx DBTX, args ...any) ([]domain.NurseEvent, error)
	GetNurseEvent(ctx context.Context, query string, dbtx DBTX, args ...any) (*domain.NurseEvent, error)
	Execute(ctx context.Context, query string, dbtx DBTX, args ...any) (sql.Result, error)
}

type RepositoryImplementaion struct {
}

func NewRepostiroy() Repository {
	return &RepositoryImplementaion{}
}

// func (r *RepositoryImplementaion) Create(ctx context.Context, ns domain.NurseEvent, tx *sql.Tx) (string, error) {

// 	result, err := tx.ExecContext(ctx, "insert into nurse_events (device_id,event,tanggal,waktu) values($1,$2,current_date,current_time)", ns.DeviceId, ns.Event)
// 	if err != nil {
// 		return "", err
// 	}

// 	if rowsAffected, err := result.RowsAffected(); err != nil {
// 		return "", err
// 	} else if rowsAffected == 0 {
// 		return "", fmt.Errorf("no rows affected ,message :%v", err)
// 	}

// 	return ns.DeviceId, nil
// }

// func (r *RepositoryImplementaion) GetByDeviceID(ctx context.Context, deviceID string, tx *sql.Tx) ([]domain.NurseEvent, error) {

// 	rows, err := tx.QueryContext(ctx, "select device_id, event,tanggal,waktu from nurse_events where device_id = $1", deviceID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var nurseEvents []domain.NurseEvent

// 	for rows.Next() {
// 		nurseEvent := &domain.NurseEvent{}
// 		if err := rows.Scan(&nurseEvent.DeviceId, &nurseEvent.Event, &nurseEvent.Tanggal, &nurseEvent.Waktu); err != nil {
// 			return nil, err
// 		}
// 		nurseEvents = append(nurseEvents, *nurseEvent)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return nurseEvents, nil

// }

func (r *RepositoryImplementaion) GetNurseEvents(ctx context.Context, query string, dbtx DBTX, args ...any) ([]domain.NurseEvent, error) {
	//  "select device_id, event,tanggal,waktu from nurse_events"
	rows, err := dbtx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var nurseEvents []domain.NurseEvent

	for rows.Next() {
		nurseEvent := &domain.NurseEvent{}
		if err := rows.Scan(&nurseEvent.DeviceId, &nurseEvent.Event, &nurseEvent.Tanggal, &nurseEvent.Waktu); err != nil {
			return nil, err
		}
		nurseEvents = append(nurseEvents, *nurseEvent)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nurseEvents, nil

}

// func (r *RepositoryImplementaion) GetLastEvents(ctx context.Context, tx *sql.Tx) ([]domain.NurseEvent, error) {
// 	rows, err := tx.QueryContext(ctx, "select distinct on (device_id) device_id,event,tanggal,waktu from nurse_events order by device_id ,tanggal desc,waktu desc")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var nurseEvents []domain.NurseEvent
// 	for rows.Next() {
// 		nurseEvent := &domain.NurseEvent{}
// 		if err := rows.Scan(&nurseEvent.DeviceId, &nurseEvent.Event, &nurseEvent.Tanggal, &nurseEvent.Waktu); err != nil {
// 			return nil, err
// 		}
// 		nurseEvents = append(nurseEvents, *nurseEvent)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return nurseEvents, nil
// }

func (r *RepositoryImplementaion) GetNurseEvent(ctx context.Context, query string, dbtx DBTX, args ...any) (*domain.NurseEvent, error) {
	NurseEvent := &domain.NurseEvent{}
	err := dbtx.QueryRowContext(ctx, query, args...).Scan(&NurseEvent.DeviceId, &NurseEvent.Event, &NurseEvent.Tanggal, &NurseEvent.Waktu)
	if err != nil {
		return nil, err
	}
	return NurseEvent, nil

}

func (r *RepositoryImplementaion) Execute(ctx context.Context, query string, dbtx DBTX, args ...any) (sql.Result, error) {
	return dbtx.ExecContext(ctx, query, args...)
}
