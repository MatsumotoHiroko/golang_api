package main

import (
	"encoding/json"
    "net/http"
    "net/http/httptest"
	"testing"
	"bytes"
    "os"
    "strconv"

	"gopkg.in/mgo.v2/bson"
    "github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	//. "github.com/MatsumotoHiroko/golang_api/config"
	//. "github.com/MatsumotoHiroko/golang_api/dao"
    . "github.com/MatsumotoHiroko/golang_api/models"
    auth "github.com/MatsumotoHiroko/golang_api/auth"
)

func TestCreateBook(t *testing.T) {
    token := getAuthToken()
    book := &Book{
        Name: "testbook1",
		Price: 300,
		Difficulty: 1,
		Released: true,
    }
    jsonBook, _ := json.Marshal(book)
    request, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBook))
    request.Header.Set("Authorization", "Bearer " + token)
	response := httptest.NewRecorder()
    router := mux.NewRouter()
    router.Handle("/books", auth.JwtMiddleware.Handler(http.HandlerFunc(CreateBookEndPoint))).Methods("POST")
    router.ServeHTTP(response, request)
    assert.Equal(t, 201, response.Code, "OK response is expected")
    assert.Contains(t, response.Body.String(), book.Name)
}

func TestUpdateBook(t *testing.T) {
    token := getAuthToken()
    base_book := insertBook()
    book := &Book{
        Name: "testbook2",
		Price: 200,
		Difficulty: 2,
		Released: false,
    }
	jsonBook, _ := json.Marshal(book)
    request, _ := http.NewRequest("PUT", "/books/" + base_book.ID.Hex(), bytes.NewBuffer(jsonBook))
    request.Header.Set("Authorization", "Bearer " + token)
	response := httptest.NewRecorder()
    router := mux.NewRouter()
    router.Handle("/books/{id}", auth.JwtMiddleware.Handler(http.HandlerFunc(UpdateBookEndPoint))).Methods("PUT")
	router.ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")

    create_book, _ := dao.FindById(base_book.ID.Hex())
    assert.Equal(t, book.Name, create_book.Name)
    assert.Equal(t, book.Price, create_book.Price)
    assert.Equal(t, book.Difficulty, create_book.Difficulty)
    assert.Equal(t, book.Released, create_book.Released)
}

func TestDeleteBook(t *testing.T) {
    token := getAuthToken()
    base_book := insertBook()
    request, _ := http.NewRequest("DELETE", "/books/" + base_book.ID.Hex(), bytes.NewBuffer(nil))
    request.Header.Set("Authorization", "Bearer " + token)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
    router.Handle("/books/{id}", auth.JwtMiddleware.Handler(http.HandlerFunc(DeleteBookEndPoint))).Methods("DELETE")
	router.ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
    create_book, _ := dao.FindById(base_book.ID.Hex())
    assert.Empty(t, create_book) 
}

func TestCreateReviewBook(t *testing.T) {
	base_book := insertBook()
    review := &Review{
        Name: "Aaron",
        Comment: "That's cool!",
        Value: 1,
    }
    jsonReview, _ := json.Marshal(review)
    request, _ := http.NewRequest("POST", "/books/" + base_book.ID.Hex() + "/review", bytes.NewBuffer(jsonReview))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
    router.HandleFunc("/books/{id}/review", CreateReviewBookEndPoint).Methods("POST")
    router.ServeHTTP(response, request)
    assert.Equal(t, 201, response.Code, "OK response is expected")
    assert.Contains(t, response.Body.String(), base_book.ID.Hex())
    assert.Contains(t, response.Body.String(), "\"value\":\"" + strconv.Itoa(review.Value) + "\"")
}

func TestGetAllBooks(t *testing.T) {
	base_book := insertBook()
    request, _ := http.NewRequest("GET", "/books", nil)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
    router.HandleFunc("/books", AllBooksEndPoint).Methods("GET")
    router.ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
    assert.Contains(t, response.Body.String(), base_book.ID.Hex())
}

func TestFindBook(t *testing.T) {
	base_book := insertBook()
    request, _ := http.NewRequest("GET", "/books/" + base_book.ID.Hex(), nil)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
    router.HandleFunc("/books/{id}", FindBookEndPoint).Methods("GET")
    router.ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
    assert.Contains(t, response.Body.String(), base_book.ID.Hex())
}


func getAuthToken() (string) {
    request, _ := http.NewRequest("GET", "/auth", nil)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
    router.HandleFunc("/auth", auth.GetTokenHandler).Methods("GET")
	router.ServeHTTP(response, request)
    return response.Body.String()
}

func insertBook() (Book) {
	book := Book{
        Name: "testbook_base",
		Price: 100,
		Difficulty: 1,
		Released: true,
	}
	book.ID = bson.NewObjectId()
	dao.Insert(book)
	return book
}

func setup() {
    println("setup")
}

func teardown() {
    println("teardown")
}

func TestMain(m *testing.M) {
    setup()
    ret := m.Run()
    teardown()
    os.Exit(ret)
}