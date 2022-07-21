package forms

type PasswordLoginForm struct {
	Mobile   string `from:"mobile" json:"mobile" binding:"required"`
	Password string `form:"mobile" json:"password" binding:"required,min=3,max=10"`
}
