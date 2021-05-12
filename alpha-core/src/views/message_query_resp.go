package views

/**
评论按照时间的降序，回复按照时间的升序
 */
type MessageQueryResponse struct {
	MessageResp []MessageResp
	Total       int64
}

type MessageResp struct {
	Id              string `json:"id"`
	Headphoto       string    `json:"head_photo"`
	Date            string `json:"date"`
	Context         string `json:"context"`
	Parent_msg_id   string `json:"parent_msg_id"`
	Parent_username string `json:"parent_username" `
	Username        string `json:"username"`
	Children        []MessageResp `json:"children"`
}
