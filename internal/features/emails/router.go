package emails

import (
	"github.com/go-chi/chi/v5"
	"github.com/mryeibis/indexer/internal/middlewares"
	"github.com/mryeibis/indexer/internal/models"
	"github.com/mryeibis/indexer/internal/zincsearch"
	"github.com/mryeibis/indexer/pkg/logger"
)

func NewRouter(logs *logger.Logger, zincSearch *zincsearch.ZincSearch[models.Email]) chi.Router {
	h := newHandler(logs, zincSearch)

	r := chi.NewRouter()
	r.Use(middlewares.JSONResponse)
	r.Get("/", h.getAll)
	return r
}
