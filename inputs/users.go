package inputs

type RegisterInput struct {
	FirstName            string `form:"first_name" json:"first_name" binding:"required"`
	LastName             string `form:"last_name" json:"last_name" binding:"required"`
	Email                string `form:"email" json:"email" validate:"email" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required"`
}

type LoginInput struct {
	Email string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,gte=8"`
}
