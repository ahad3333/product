package models

type CreateUser struct {
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	BranchId  string `json:"branch_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserPrimeryKey struct {
	Id string `json:"id"`
}

type GetListUserResponse struct {
	Count int64   `json:"count"`
	Users []*User `json:"user"`
}

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetListUserRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	BranchId  string `json:"branch_id"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUserSwag struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	BranchId string `json:"branch_id"`
}
