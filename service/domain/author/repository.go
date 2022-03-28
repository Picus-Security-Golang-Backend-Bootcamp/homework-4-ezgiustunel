package author

import (
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) Migration() {
	r.db.AutoMigrate(&Author{})
}

// InsertData: insert data to db
func (b *AuthorRepository) InsertData(author Author) {
	b.db.Where(Author{ID: author.ID}).Attrs(Author{AuthorName: author.AuthorName}).FirstOrCreate(&author)
}

// FindAll: finds all elements in db
func (b *AuthorRepository) FindAll() []Author {
	var authors []Author
	b.db.Find(&authors)

	return authors
}
