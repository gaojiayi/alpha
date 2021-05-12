package views

type ImageDetectResp struct {
	ImageDetectItems []ImageDetectItem `json:"image_detect_items"`
}

type ImageDetectItem struct {
	FileName           string        `json:"file_name"`
	ImageDetectObjects []ImageDetect `json:"image_detect_objects"`
}
type ImageDetect struct {
	Name   string `json:"name"`
	Percet string `json:"percet"`
	Url    string `json:"url"`
}
