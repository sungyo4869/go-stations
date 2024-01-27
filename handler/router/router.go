package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	// エンドポイントの定義
	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	return mux
}
