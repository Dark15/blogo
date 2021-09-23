package response

type CoverImage struct {
	Url      string `json:"url" form:"url"`
	Image    string `json:"image" form:"image" `
	Mime     string `json:"mime" form:"mime"`
	FileName string `json:"filename" form:"filename"`
}
