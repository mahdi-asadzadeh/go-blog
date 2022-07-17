package inputs

type CreateCommentInput struct {
	Content string `form:"content" json:"content" binding:"required"`
}
