package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pol905/word/entities"
	"github.com/pol905/word/repositories"
	"github.com/pol905/word/utils"
)

type bookHandlers struct {
	bookRepository repositories.BookRepository
}

type PostBook struct {
}

func BookRouter(bookRepository repositories.BookRepository) http.Handler {
	bookHandler := NewBookHandlers(bookRepository)
	r := chi.NewRouter()

	r.Get("/", bookHandler.getBooks)
	r.Get("/{id}", bookHandler.getBook)
	r.Post("/", bookHandler.createBook)
	r.Put("/{id}", bookHandler.updateBook)
	r.Delete("/{id}", bookHandler.deleteBook)
	return r
}

func NewBookHandlers(bookRepository repositories.BookRepository) *bookHandlers {
	return &bookHandlers{bookRepository}
}

func (bh *bookHandlers) getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookRepository.Find(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, books)
}

func (bh *bookHandlers) getBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := bh.bookRepository.FindById(r.Context(), id)

	utils.WriteJson(w, book)
}

func (bh *bookHandlers) createBook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book

	if err := utils.ReadJson(r, &book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := validator.New()
	if err := v.Struct(book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := bh.bookRepository.Create(r.Context(), &book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, book)
}

func (bh *bookHandlers) updateBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book entities.Book

	if err := utils.ReadJson(r, &book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := validator.New()
	if err := v.Struct(book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book.ID = id
	err = bh.bookRepository.Update(r.Context(), &book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, book)
}

func (bh *bookHandlers) deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bh.bookRepository.Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
