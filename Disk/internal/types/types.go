// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Code     string `json:"code"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
}

type RegisterRes struct {
}

type EmailSendReq struct {
	Email string `json:"email"`
}

type EmailSendRes struct {
}

type UpdateUserDetailReq struct {
}

type GetUserDetailReq struct {
	UserId int64 `path:"userId"`
}

type UpdateAvatarReq struct {
}
