package api

import (
	"net/http"

	"github.com/Hudson256/AskMeAnything-Server/internal/store/pgstore"
)

type apiHandler struct {
	q *pgstore.Queries
}

func (h apiHandler) ServeHTTP(http.ResponseWriter, r *http.Request)

func NewHandler(q *pgstore.Queries) http.Handler {

}
