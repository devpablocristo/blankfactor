package api

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"

	application "github.com/devpablocristo/blankfactor/event-service/internal/application"
	eventdb "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driven-adapter/repository/mysql"
	handler "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driver-adapter/handler"
)

func StartApi(wg *sync.WaitGroup, port string) {
	// event1 := domain.Event{
	// 	ID:        1,
	// 	StartTime: time.Date(2023, time.April, 13, 10, 0, 0, 0, time.UTC),
	// 	EndTime:   time.Date(2023, time.April, 13, 11, 0, 0, 0, time.UTC),
	// }

	// event2 := domain.Event{
	// 	ID:        2,
	// 	StartTime: time.Date(2023, time.April, 14, 14, 0, 0, 0, time.UTC),
	// 	EndTime:   time.Date(2023, time.April, 14, 15, 0, 0, 0, time.UTC),
	// }

	// event3 := domain.Event{
	// 	ID:        3,
	// 	StartTime: time.Date(2023, time.April, 15, 16, 0, 0, 0, time.UTC),
	// 	EndTime:   time.Date(2023, time.April, 15, 17, 0, 0, 0, time.UTC),
	// }

	// abre una conexión a la base de datos
	db, err := eventdb.RepositoryConn()
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	// _, err = db.Exec("CREATE DATABASE IF NOT EXISTS events_service DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci")
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = db.Exec("USE events_service")
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = db.Exec("GRANT ALL PRIVILEGES ON events_service.* TO 'tester'@'event-mysql-repo'")
	// if err != nil {
	// 	panic(err)
	// }

	// inserta cada evento en la base de datos
	// _, err = db.Exec("INSERT INTO events_service (id, start_time, end_time) VALUES (?, ?, ?)", event1.ID, event1.StartTime, event1.EndTime)
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = db.Exec("INSERT INTO events_service (id, start_time, end_time) VALUES (?, ?, ?)", event2.ID, event2.StartTime, event2.EndTime)
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = db.Exec("INSERT INTO events_service (id, start_time, end_time) VALUES (?, ?, ?)", event3.ID, event3.StartTime, event3.EndTime)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Eventos insertados en la base de datos")

	msq := eventdb.NewEventRepository(db)
	esv := application.NewEventService(msq)
	han := handler.NewHandler(esv)
	rou := Router(han)

	HttpServer(port, rou)
}
