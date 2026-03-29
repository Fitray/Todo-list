package core_http_responce

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/Fitray/Todo-list/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponceHandler struct {
	log *core_logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponceHandler(
	log *core_logger.Logger,
	w http.ResponseWriter,
) *HTTPResponceHandler {
	return &HTTPResponceHandler{
		log: log,
		w:   w,
	}
}

func (h *HTTPResponceHandler) PanicResponce(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.w.WriteHeader(statusCode)

	responce := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.w).Encode(responce); err != nil {
		h.log.Error("write HTTP responce", zap.Error(err))
	}

}
