package post

import "unicode/utf8"

type Post struct {
	Id           string
	UserId       string
	PostStatusId int
	Title        string
	Body         string
}

var (
	titleLengthMin = 1
	titleLengthMax = 128
	titleLengthErr = "タイトルの長さが正しくないです"
	bodyLengthMin  = 1
	bodyLengthMax  = 3000
	bodyLengthErr  = "本文の長さが正しくないです"
)

func NewPost(
	id string,
	userId string,
	postStatusId int,
	title string,
	body string,
) (*Post, *string) {
	return newPost(
		id,
		userId,
		postStatusId,
		title,
		body,
	)
}

func newPost(
	id string,
	userId string,
	postStatusId int,
	title string,
	body string,
) (*Post, *string) {

	if utf8.RuneCountInString(title) <= titleLengthMin || utf8.RuneCountInString(title) >= titleLengthMax {
		return nil, &titleLengthErr
	}
	if utf8.RuneCountInString(body) <= bodyLengthMin || utf8.RuneCountInString(body) >= bodyLengthMax {
		return nil, &bodyLengthErr
	}

	return &Post{
		Id:           id,
		UserId:       userId,
		PostStatusId: postStatusId,
		Title:        title,
		Body:         body,
	}, nil
}
