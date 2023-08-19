package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/pol905/word/entities"
	"gorm.io/gorm"
)

type BookRepository interface {
	Find(ctx context.Context) ([]entities.Book, error)
	FindById(ctx context.Context, id uuid.UUID) (entities.Book, error)
	Create(ctx context.Context, book *entities.Book) error
	Update(ctx context.Context, book *entities.Book) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type book struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &book{db}
}

func (b *book) Find(ctx context.Context) ([]entities.Book, error) {
	books := []entities.Book{}
	res := b.db.WithContext(ctx).Find(&books)
	return books, res.Error
}

func (b *book) FindById(ctx context.Context, id uuid.UUID) (entities.Book, error) {
	book := entities.Book{}
	res := b.db.WithContext(ctx).Find(&book, map[string]string{
		"id": id.String(),
	})

	return book, res.Error
}

func (b *book) Create(ctx context.Context, book *entities.Book) error {
	res := b.db.WithContext(ctx).Create(book)

	return res.Error
}

func (b *book) Update(ctx context.Context, book *entities.Book) error {
	res := b.db.WithContext(ctx).Model(&book).Updates(book)

	return res.Error
}

func (b *book) Delete(ctx context.Context, id uuid.UUID) error {
	res := b.db.WithContext(ctx).Delete(entities.Book{ID: id})

	return res.Error
}
