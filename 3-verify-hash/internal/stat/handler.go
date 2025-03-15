package stat

import (
	"errors"
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/response"
	"log"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type statHandler struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandlerDependencies struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, dependencies StatHandlerDependencies) {
	handler := &statHandler{
		StatRepository: dependencies.StatRepository,
		Config:         dependencies.Config,
	}
	router.HandleFunc("GET /stat", handler.GetStat())
}

func (handler *statHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		fromString := query.Get("from")
		toString := query.Get("to")
		period := query.Get("period")
		if period != GroupByDay && period != GroupByMonth {
			response.Json(w, errors.New("period is not valid"), http.StatusBadRequest)
			return
		}
		// Parse dates
		fmt.Println(fromString, toString, period)
		from, err := time.Parse("2006-01-02", fromString)
		if err != nil {
			log.Println(err.Error())
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", toString)
		if err != nil {
			log.Println(err.Error())
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Get stats
		stats := handler.StatRepository.GetStatsByPeriod(from, to, period)
		data := GetStatResponse{
			Data: stats,
		}
		// Response
		response.Json(w, data, http.StatusOK)
	}
}
