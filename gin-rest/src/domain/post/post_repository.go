package post

type PostRepository interface {
	GetAll() ([]*Post, error)
	FindById(id string) (*Post, error)
	Create(post *Post) error
	Edit(post *Post) error
	Delete(id string) error
}
