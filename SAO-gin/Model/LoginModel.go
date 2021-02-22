package model

type UserFilter struct {
	Username string `json:"username", form:"username"`
	Password string `json:"password", form:"password"`
}

type Password struct {
	Password string `db:"Password" json:"password"`
}

type User struct {
	Username string `db:"Username" json:"username"`
	Nickname string `db:"Nickname" json:"nickname"`
	Role     string `db:"Role" json:"role"`
}
