package converter

import (
	"client/internal/entities"
	"fmt"
)

func CommentToString(c *entities.Comment) string {
	return fmt.Sprintf("id: %d\n%s\n%s\nauthor: %s", c.CommentId, c.Text, c.Date.Format("2 January 2006 at 15:04:05"), c.AuthorMail)
}
