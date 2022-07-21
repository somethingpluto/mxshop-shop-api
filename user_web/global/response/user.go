package response

import "time"

type UserResponse struct {
	Id       int32     `json:"id"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Mobile   string    `json:"mobile"`
	Birthday time.Time `json:"birthday"`
}
