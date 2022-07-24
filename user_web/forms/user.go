package forms

// PasswordLoginForm
// @Description: 手机号 密码登录表单规则
//
type PasswordLoginForm struct {
	Mobile    string `from:"mobile" json:"mobile" binding:"required,mobile"`
	Password  string `form:"mobile" json:"password" binding:"required,min=3,max=10"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
