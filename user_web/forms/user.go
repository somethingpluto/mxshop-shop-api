package forms

// PasswordLoginForm
// @Description: 手机号 密码登录表单规则
//
type PasswordLoginForm struct {
	Mobile   string `from:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"mobile" json:"password" binding:"required,min=3,max=10"`
}
