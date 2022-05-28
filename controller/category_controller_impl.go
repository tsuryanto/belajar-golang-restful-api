package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/tsuryanto/belajar-golang-restful-api/helper"
	"github.com/tsuryanto/belajar-golang-restful-api/model/web"
	"github.com/tsuryanto/belajar-golang-restful-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	id, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id // set categoryId
	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
