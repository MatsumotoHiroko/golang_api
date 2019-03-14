package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

  "gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	
	. "github.com/MatsumotoHiroko/golang_api/config"
	. "github.com/MatsumotoHiroko/golang_api/dao"
	. "github.com/MatsumotoHiroko/golang_api/models"
	auth "github.com/MatsumotoHiroko/golang_api/auth"
)

var config = Config{}
var dao = BooksDAO{}
var review_dao = ReviewsDAO{}

func AllBooksEndPoint(w http.ResponseWriter, r *http.Request) {
	books, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func FindBookEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	respondWithJson(w, http.StatusOK, book)
}

func CreateBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	validate := validator.New() 
	if err := validate.Struct(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	book.ID = bson.NewObjectId()
	book.CreatedAt = time.Now()
	if err := dao.Insert(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, book)
}

func UpdateBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	book, err := dao.FindById(params["id"])
	book_base_id := book.ID
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.ID = book_base_id
	if err := dao.Update(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
  respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	book, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	if err := dao.Delete(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func CreateReviewBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var review Review
	params := mux.Vars(r)
	book, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	validate := validator.New() 
	if err := validate.Struct(review); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	review.ID = bson.NewObjectId()
	review.BookId = book.ID
	review.CreatedAt = time.Now()
	if err := review_dao.Insert(review); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, review)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()

	review_dao.Server = config.Server
	review_dao.Database = config.Database
	review_dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", AllBooksEndPoint).Methods("GET")
	r.Handle("/books", auth.JwtMiddleware.Handler(http.HandlerFunc(CreateBookEndPoint))).Methods("POST")
	r.Handle("/books/{id}", auth.JwtMiddleware.Handler(http.HandlerFunc(UpdateBookEndPoint))).Methods("PUT")
	r.Handle("/books/{id}", auth.JwtMiddleware.Handler(http.HandlerFunc(UpdateBookEndPoint))).Methods("PATCH")
	r.Handle("/books/{id}", auth.JwtMiddleware.Handler(http.HandlerFunc(DeleteBookEndPoint))).Methods("DELETE")
	r.HandleFunc("/books/{id}", FindBookEndPoint).Methods("GET")
	r.HandleFunc("/books/{id}/review", CreateReviewBookEndPoint).Methods("POST")
	r.HandleFunc("/auth", auth.GetTokenHandler).Methods("GET")
    if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}