package controllers

import (
	"book_store/pkg/models"
	"book_store/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res,_ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	bookId:=params["bookId"]
	ID, err:=strconv.ParseInt(bookId, 0, 0)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookDetails, _ :=models.GetBookById(ID)
	res,_ :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	createbook:=models.Book{}
	utils.ParseBody(r, &createbook)
	book:=createbook.CreateBook()
	res,_ :=json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
    params:=mux.Vars(r)
	bookId:=params["bookId"]
	ID, err:=strconv.ParseInt(bookId, 0, 0)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookDetails,_:=models.GetBookById(ID)
	if bookDetails.ID==0{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if updateBook.Name!=""{
		bookDetails.Name=updateBook.Name
	}
	if updateBook.Author!=""{
		bookDetails.Author=updateBook.Author
	}
	if updateBook.Publication!=""{
		bookDetails.Publication=updateBook.Publication
	}
	bookDetails.UpdateBook()
	res,_ :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	bookId:=params["bookId"]
	ID, err:=strconv.ParseInt(bookId, 0, 0)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	models.DeleteBook(ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))
}