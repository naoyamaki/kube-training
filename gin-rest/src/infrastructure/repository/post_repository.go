package repository

import (
	"fmt"

	"github.com/naoyamaki/domain/post"
)

type postRepository struct{}

func NewPostRepository() post.PostRepository {
	return &postRepository{}
}

func (r *postRepository) GetAll() ([]*post.Post, error) {
	result, err := post.NewPost("asdfghjkl", "asdf", 1, "titlehoge", "bodyhogehoge")
	if err != nil {
		fmt.Println(err)
	}
	var posts []*post.Post
	posts = append(posts, result, result)
	return posts, nil
}
func (r *postRepository) FindById(id string) (*post.Post, error) {
	result, err := post.NewPost("asdfghjkl", "asdf", 1, "titlehoge", "bodyhogehoge")
	if err != nil {
		fmt.Println(err)
	}
	return result, nil
}
func (r *postRepository) Create(post *post.Post) error {
	return nil
}
func (r *postRepository) Edit(post *post.Post) error {
	return nil
}
func (r *postRepository) Delete(id string) error {
	return nil
}
