package controllers

import (
	"encoding/json"
	"github.com/pluralsight/webservice/models"
	"net/http"
	"regexp"
)

type transactionController struct {
	transactionIDPattern *regexp.Regexp
}

func (uc transactionController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/transactions" || r.URL.Path == "/transactions/" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.transactionIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id := matches[1]
		
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *transactionController) getAll(w http.ResponseWriter, r *http.Request) {
	us := models.GetTransactions()
	if us != nil {
		encodeResponseAsJSON(us, w)
	} else {
		encodeResponseAsJSON("There are no any transactions", w)
	}
}

func (uc *transactionController) get(id string, w http.ResponseWriter) {
	u, err := models.GetTransactionByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *transactionController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Transaction object"))
		return
	}
	u, err = models.AddTransaction(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *transactionController) put(id string, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Transaction object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted transaction must match ID in URL"))
		return
	}
	u, err = models.UpdateTransaction(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *transactionController) delete(id string, w http.ResponseWriter) {
	err := models.RemoveTransactionById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *transactionController) parseRequest(r *http.Request) (models.Transaction, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Transaction
	err := dec.Decode(&u)
	if err != nil {
		return models.Transaction{}, err
	}
	return u, nil
}

func newTransactionController() *transactionController {
	return &transactionController{
		transactionIDPattern: regexp.MustCompile(`^/transactions/(\w+)/?`),
	}
}
