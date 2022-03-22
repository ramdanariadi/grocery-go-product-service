package transaction

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TransactionController interface {
	FindByUserId(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Save(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}
