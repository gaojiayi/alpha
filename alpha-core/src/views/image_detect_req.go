package views

/**
图片检测请求参数
 */

type ImageDetectRequest struct {
	// 如果是application/json   则使用json 而不是form
	FileName     string `json:"file_name"`
	FileSaveName string `json:"file_save_name"`
}
