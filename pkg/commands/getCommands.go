package commands

import "fmt"

func Get() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		Register, Login, GetProfile, WritePost, GetPosts, EditPost,
		GetPost, DeletePost, WriteComment, GetComments, DeleteComment)
}
