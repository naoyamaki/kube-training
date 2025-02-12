package usecase

import (
	"github.com/gin-gonic/gin"

	"gin-rest/domain/post"
)

type PostUsecase struct {
	postRepository post.PostRepository
}

func NewPostUsecase(postRepository post.PostRepository) *PostUsecase {
	return &PostUsecase{
		postRepository: postRepository,
	}
}

func (uc *PostUsecase) GetAllPosts(c *gin.Context) (int, gin.H) {
	posts, _ := uc.postRepository.GetAll()
	return 200, gin.H{
		"id":           posts[0].Id,
		"userId":       posts[0].UserId,
		"postStatusId": posts[0].PostStatusId,
		"title":        posts[0].Title,
		"body":         posts[0].Body,
	}
}

func (uc *PostUsecase) GetPostByID(c *gin.Context) (int, gin.H) {
	id := c.Param("id")

	var post, _ = uc.postRepository.FindById(id)
	return 200, gin.H{
		"id":           post.Id,
		"userId":       post.UserId,
		"postStatusId": post.PostStatusId,
		"title":        post.Title,
		"body":         post.Body,
	}
}

func (uc *PostUsecase) CreatePost(c *gin.Context) (int, gin.H) {
	post, _ := post.NewPost("hoge", "hoge", 1, "hoge", "hogehogehogehoge")
	result := uc.postRepository.Create(post)
	if result != nil {
		return 500, gin.H{"message": "error"}
	}
	return 200, gin.H{"message": "Post created"}
}

func (uc *PostUsecase) DeletePost(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	result := uc.postRepository.Delete(id)
	if result != nil {
		return 500, gin.H{"message": "error"}
	}
	return 200, gin.H{"message": "Deleted post", "id": id}
}

func (uc *PostUsecase) EditPost(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	post, _ := post.NewPost("hoge", "hoge", 1, "hoge", "hogehogehogehoge")
	result := uc.postRepository.Edit(post)
	if result != nil {
		return 500, gin.H{"message": "error"}
	}
	return 200, gin.H{"message": "Edited post", "id": id}
}
