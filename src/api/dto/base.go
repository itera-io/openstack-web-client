package dto

type AuthUtils struct {
	BaseUrl string
	Token   string
}

type CommonDto struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
