package models

type CreateClient struct {
	Name string `json:"name"`
}

type ClientPrimeryKey struct {
	Id string `json:"id"`
}

type Client struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
