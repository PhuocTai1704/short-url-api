package request

type RequestLink struct {
	Url string `json:"url" binding:"required,url"`
}