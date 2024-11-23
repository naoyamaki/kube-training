package user

import (
	"unicode/utf8"

	"github.com/dlclark/regexp2"
)

type User struct {
	id                string
	email             string
	name              string
	user_status_id_fk int
}

var (
	emailPattern   = `^[a-Z0-9_.+-]+@([a-Z0-9][a-Z0-9-]*[a-Z0-9]*\.)+[a-Z]{2,}$`
	emailLengthMax = 254
	emailErr       = "メールの形式が正しくないです"
	nameLengthMin  = 2
	nameLengthMax  = 32
	nameLengthErr  = "名前の長さが正しくないです"
)

func newUser(
	id string,
	email string,
	name string,
	user_status_id_fk int,
) (*User, *string) {
	re := regexp2.MustCompile(emailPattern, 0)
	isEmail, _ := re.MatchString(email)
	if !isEmail {
		return nil, &emailErr
	}
	if utf8.RuneCountInString(email) >= emailLengthMax {
		return nil, &emailErr
	}
	if utf8.RuneCountInString(name) <= nameLengthMin || utf8.RuneCountInString(name) >= nameLengthMax {
		return nil, &nameLengthErr
	}
	return &User{
		id:                id,
		email:             email,
		name:              name,
		user_status_id_fk: user_status_id_fk,
	}, nil
}
