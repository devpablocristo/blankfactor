package mysql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"

	port "github.com/devpablocristo/blankfactor/event-service/internal/application/port"
	domain "github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

type eventDTO struct {
	id        int
	StartTime string
	EndTime   string
}

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) port.EventRepo {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *domain.Event) error {
	query := "INSERT INTO `events` (start_time, end_time) VALUES (?, ?)"

	result, err := r.db.ExecContext(ctx, query, event.StartTime, event.EndTime)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = int(id)

	return nil
}

func (r *EventRepository) GetEventByID(ctx context.Context, id int) (*domain.Event, error) {
	query := "SELECT id, start_time, end_time FROM events_service WHERE id = ?"

	row := r.db.QueryRowContext(ctx, query, id)

	var event domain.Event
	if err := row.Scan(&event.ID, &event.StartTime, &event.EndTime); err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *EventRepository) UpdateEvent(ctx context.Context, event *domain.Event) error {
	query := "UPDATE events_service SET start_time = ?, end_time = ? WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query, event.StartTime, event.EndTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) DeleteEvent(ctx context.Context, id int) error {
	query := "DELETE FROM events_service WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	query := "SELECT id, start_time, end_time FROM `events`"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.New("rows - " + err.Error())
	}
	defer rows.Close()

	events := make([]*domain.Event, 0)

	for rows.Next() {
		var e eventDTO
		if err := rows.Scan(&e.id, &e.StartTime, &e.EndTime); err != nil {
			return nil, errors.New("error de tipo - " + err.Error())
		}

		event := convertEventDtoToDomainEvent(e)
		events = append(events, &event)
	}

	return events, nil
}

func convertEventDtoToDomainEvent(e eventDTO) domain.Event {
	var event domain.Event
	var err error

	event.ID = e.id

	event.StartTime, err = time.Parse("2006-01-02 15:04:05", e.StartTime)
	if err != nil {
		panic(err)
	}

	event.EndTime, err = time.Parse("2006-01-02 15:04:05", e.EndTime)
	if err != nil {
		panic(err)
	}

	return event
}
