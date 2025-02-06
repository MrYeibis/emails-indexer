package emails

import (
	"encoding/json"
	"net/http"

	"github.com/mryeibis/indexer/internal/models"
	"github.com/mryeibis/indexer/internal/zincsearch"
	"github.com/mryeibis/indexer/pkg/logger"
)

type handler struct {
	logs       *logger.Logger
	zincSearch *zincsearch.ZincSearch[models.Email]
}

func newHandler(logs *logger.Logger, zincSearch *zincsearch.ZincSearch[models.Email]) *handler {
	return &handler{
		logs:       logs,
		zincSearch: zincSearch,
	}
}

func (h *handler) getAll(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	params := filterFromValues(&values)

	searchParams := getAllsearchParamsFromFilter(params)

	response, err := h.zincSearch.GetAll(*searchParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(getAllResponseFromZincSearchResponse(response))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}
}
