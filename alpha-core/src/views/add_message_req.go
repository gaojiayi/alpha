package views

/**
参数校验后续完善
 */
type AddMessageRequest struct {
	Context        string `json:"context"`
	ParentMsgId    string `json:"parent_msg_id"`
	ParentUsername string `json:"parent_username" `
	Username       string `json:"username"`
}
