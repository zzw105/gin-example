package types

// BaseResponse 响应结构
type BaseResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Options struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 用户信息
type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username" binding:"required,min=3,max=30"` // 字符串，必填，长度在 3-30 之间
	Password string  `json:"password" binding:"required,min=6,max=20"` // 字符串，必填，长度在 6-20 之间
	Email    *string `json:"email"`                                    // 必填，必须是合法的邮箱格式
}

// 工作日历信息
type WorkCalendar struct {
	ID      int    `json:"id"`
	Day     string `json:"day"`     // 字符串，必填，长度在 3-30 之间
	Message string `json:"message"` // 字符串，必填，长度在 6-20 之间
}
