package converter

import (
	"client/internal/entities"
	"fmt"
)

func PostToString(post *entities.Post) string {
	return fmt.Sprintf("id: %d\n%s\n%s\n%s\nauthor: %s", post.PostId, post.Header, post.Body, post.Date.Format("2 January 2006 at 15:04:05"), post.AuthorMail)
}
