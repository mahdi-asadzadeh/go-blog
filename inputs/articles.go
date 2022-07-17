package inputs

type CreateArticleInput struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Body        string `form:"body" json:"body" binding:"required"`
	Categories  []struct {
		Name        string `form:"name" json:"name" binding:"min=4,max255"`
		Description string `form:"description" json:"description"`
	} `form:"categories" json:"categories"`
	Tags []struct {
		Name        string `form:"name" json:"name" binding:"exists,alphanum,min=4,max=255"`
		Description string `form:"description" json:"description" binding:"exists"`
	} `json:"tags"`
}
