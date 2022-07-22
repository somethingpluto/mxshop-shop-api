package response

import "time"

// UserResponse
// @Description: 用户信息
//
type UserResponse struct {
	Id       int32     `json:"id"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Mobile   string    `json:"mobile"`
	Birthday time.Time `json:"birthday"`
}
