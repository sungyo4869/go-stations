package handler

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// レスポンスするメッセージを定義
	res := &model.HealthzResponse{
		Message: "OK",
	}
	// JSONにシリアライズ
	err := json.NewEncoder(w).Encode(res)
	// 失敗した場合、エラーを出力
	if err != nil {
		log.Print(err)
	}
}
