package api

import (
	nurseevent "backend/service/NurseEvent"
	"backend/service/device"
	"backend/utils"
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type ApiServer struct {
	Addr string
	Db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{Addr: addr, Db: db}
}

func (s *ApiServer) Run() error {
	router := httprouter.New()

	router.NotFound = utils.NotFoundHandler()
	var validator *validator.Validate = validator.New()

	deviceRepository := device.NewRepository()
	deviceService := device.NewService(deviceRepository, s.Db)
	deviceHandler := device.NewHandler(*deviceService, validator)
	deviceHandler.RegisterRoute(router)

	nurseEventRepository := nurseevent.NewRepostiroy()
	nurseEventService := nurseevent.NewService(nurseEventRepository, s.Db, *deviceService)
	nurseEventHandler := nurseevent.NewHandler(*nurseEventService, validator)
	nurseEventHandler.RegisterRoute(router)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"refresh_token", "Content-Type", "Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	handler := c.Handler(router)

	return http.ListenAndServe(s.Addr, handler)

}
