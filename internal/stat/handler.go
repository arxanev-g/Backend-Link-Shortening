package stat

import (
	"net/http"
	"time"

	"github.com/arxanev/adv/config"
	"github.com/arxanev/adv/middleware"
	"github.com/arxanev/adv/pkg/res"
)

const (
	GroupByDay    = "day"
	GroupByMounth = "mounth"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *config.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (h *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-01", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from params", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-01", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid from params", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMounth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		stats := h.StatRepository.GetSTats(by, from, to)
		res.Json(w, stats, 200)
	}
}
