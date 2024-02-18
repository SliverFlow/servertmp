package request

type UserCreateReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDeleteReq struct {
	Id int64 `json:"id" binding:"required"`
}

type UserUpdateReq struct {
	Id       int64  `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
