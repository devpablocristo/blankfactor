package api

import (
	"sync"

	application "github.com/devpablocristo/blankfactor/event-service/internal/application"
	eventdb "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driven-adapter/repository/mysql"
	handler "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driver-adapter/handler"
)

func StartApi(wg *sync.WaitGroup, port string) {
	db, err := eventdb.RepositoryConn()
	if err != nil {
		panic(err)
	}

	msq := eventdb.NewEventRepository(db)
	esv := application.NewEventService(msq)
	han := handler.NewHandler(esv)
	rou := Router(han)

	HttpServer(port, rou)
}
