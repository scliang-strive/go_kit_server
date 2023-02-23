package article

type CreateReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	CateId  int64  `json:"cate_id"`
}

type CreateResp struct {
	Id int64 `json:"id"`
}

type DetailReq struct {
	Id int64 `json:"id"`
}

type DetailResp struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	CateId  int64  `json:"cate_id"`
	UserId  int64  `json:"user_id"`
}
