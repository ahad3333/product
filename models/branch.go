package models

type Branch struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BranchPrimeryKey struct {
	Id string `json:"id"`
}

type CreateBranch struct {
	Name string `json:"name"`
}
