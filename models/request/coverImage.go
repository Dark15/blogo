package request

type CoverImage struct {
	Image    string `json:"image" form:"image" `
	Mime     string `json:"mime" form:"mime"`
	FileName string `json:"filename" form:"filename"`
}
