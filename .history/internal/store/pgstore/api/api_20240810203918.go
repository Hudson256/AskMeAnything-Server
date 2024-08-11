package api

import "github.com/Hudson256/AskMeAnything-Server/internal/store/pgstore"

type apiHandler struct {
	q *pgstore.Queries
}

func (h apiHandler) ServeHTTP

func NewHandler(q *pgstore.Queries) http.Handler {

}
