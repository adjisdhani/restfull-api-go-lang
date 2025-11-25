package test

import (
	"belajar_golang_restful_api/app"
	"belajar_golang_restful_api/controller"
	"belajar_golang_restful_api/helper"
	"belajar_golang_restful_api/middleware"
	"belajar_golang_restful_api/model/domain"
	"belajar_golang_restful_api/repository"
	"belajar_golang_restful_api/service"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setUpNewDB() *sql.DB {
	sqlDB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang_restfull_api_test")
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return sqlDB
}

func setupRouter(db *sql.DB) http.Handler {
	validator := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateTable(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE category")
	helper.PanicIfError(err)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name": "Golang"}`)
	request := httptest.NewRequest("POST", "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest("POST", "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Golangs",
	})
	tx.Commit()

	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name": "Golang"}`)
	request := httptest.NewRequest("PUT", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Golang", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Golangs",
	})
	tx.Commit()

	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest("PUT", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Golangs",
	})
	tx.Commit()

	router := setupRouter(db)
	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	router := setupRouter(db)
	request := httptest.NewRequest("DELETE", "http://localhost:8080/api/categories/1", nil)
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}

func TestFindAll(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Golangs",
	})
	tx.Commit()

	router := setupRouter(db)
	request := httptest.NewRequest("GET", "http://localhost:8080/api/categories", nil)
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, category.Id, int(responseBody["data"].([]interface{})[0].(map[string]interface{})["id"].(float64)))

	paging := responseBody["paging"].(map[string]interface{})
	assert.Equal(t, 1, int(paging["current_page"].(float64)))
	assert.Equal(t, 1, int(paging["total_page"].(float64)))
	assert.Equal(t, 1, int(paging["total_item"].(float64)))
}

func TestFindByIdSuccess(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Golangs",
	})
	tx.Commit()

	router := setupRouter(db)
	request := httptest.NewRequest("GET", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Golangs", responseBody["data"].(map[string]interface{})["name"])
}

func TestFindByIdFailed(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	router := setupRouter(db)
	request := httptest.NewRequest("GET", "http://localhost:8080/api/categories/"+strconv.Itoa(0), nil)
	request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
}

func TestUnauthorized(t *testing.T) {
	db := setUpNewDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Golangs",
	})
	tx.Commit()

	router := setupRouter(db)
	request := httptest.NewRequest("GET", "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), nil)
	// request.Header.Add("X-API-KEY", "adjis_ganteng_banget")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
}
