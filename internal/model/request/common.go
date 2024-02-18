package request

type IdReq struct {
	Id int64 `json:"id" binding:"required"`
}
