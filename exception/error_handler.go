package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tsuryanto/belajar-golang-restful-api/helper"
	"github.com/tsuryanto/belajar-golang-restful-api/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}
	if validateError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	if exception, ok := err.(NotFoundError); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)

		return true
	} else {
		return false
	}
}

func validateError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	if exception, ok := err.(validator.ValidationErrors); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
