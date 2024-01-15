package app

import (
	"EffectiveMobileTask/internal/database"
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
	Sql    *sql.Sql
	Logger *logrus.Logger
}

func NewApplication() *App {
	return &App{
		Server: &http.Server{Addr: ":8080"},
		Sql:    sql.NewDatabase(),
		Logger: logrus.New(),
	}
}

func (a *App) Run() {
	err := a.Sql.Connect()
	if err != nil {
		a.Logger.Infof("ошибка при подключении - %s", err)
		return
	}

	a.Logger.Info("успешное подключении к БД")

	http.HandleFunc("/", a.LoggingMiddleware(http.MethodGet, a.GetAllRecordsHandler))
	http.HandleFunc("/create_record", a.LoggingMiddleware(http.MethodPost, a.CreateRecordHandler))
	http.HandleFunc("/read_record_by_full_name", a.LoggingMiddleware(http.MethodGet, a.ReadRecordByFullNameHandler))
	http.HandleFunc("/read_record_by_age", a.LoggingMiddleware(http.MethodGet, a.ReadRecordByAgeHandler))
	http.HandleFunc("/read_record_by_gender", a.LoggingMiddleware(http.MethodGet, a.ReadRecordByGenderHandler))
	http.HandleFunc("/read_record_by_country", a.LoggingMiddleware(http.MethodGet, a.ReadRecordByCountryHandler))
	http.HandleFunc("/update_record_by_name", a.LoggingMiddleware(http.MethodPut, a.UpdateRecordByNameHandler))
	http.HandleFunc("/update_record_by_surname", a.LoggingMiddleware(http.MethodPut, a.UpdateRecordBySurnameHandler))
	http.HandleFunc("/update_record_by_patronymic", a.LoggingMiddleware(http.MethodPut, a.UpdateRecordByPatronymicHandler))
	http.HandleFunc("/update_record_by_age", a.LoggingMiddleware(http.MethodPut, a.UpdateRecordByAgeHandler))
	http.HandleFunc("/delete_record", a.LoggingMiddleware(http.MethodDelete, a.DeleteRecordHandler))

	err = a.Server.ListenAndServe()
	if err != nil {
		a.Logger.Info("завершение работы сервера")
	}
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Info("завершение...")
	}
}
