package forms

// SendSmsForm
// @Description: 手机验证码表单验证
//
type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Type   int    `form:"type" json:"type" binding:"required,oneof=1 2"`
}
