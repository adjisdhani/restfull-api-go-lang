package controller

import (
	"belajar_golang_restful_api/helper"
	"belajar_golang_restful_api/model/web"
	"belajar_golang_restful_api/service"
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createCategoryRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &createCategoryRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), createCategoryRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Data:   categoryResponse,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	createUpdateRequest := web.CategoryUpdateRequest{}

	err := decoder.Decode(&createUpdateRequest)
	helper.PanicIfError(err)

	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	createUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), createUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Data:   categoryResponse,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Data:   nil,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	size, _ := strconv.Atoi(request.URL.Query().Get("size"))
	if size == 0 {
		size = 10
	}

	categoryResponses, total := controller.CategoryService.FindAll(request.Context(), page, size)

	webResponse := web.WebResponse{
		Code:   200,
		Data:   categoryResponses,
		Status: "OK",
		Paging: map[string]interface{}{
			"current_page": page,
			"total_page":   int(math.Ceil(float64(total) / float64(size))),
			"total_item":   total,
		},
	}

	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Data:   categoryResponse,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(webResponse)
	helper.PanicIfError(err)
}
