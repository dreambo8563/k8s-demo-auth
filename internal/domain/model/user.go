package model

// User struct
type User struct {
	ID string `json:"uid,omitempty"`
}

//GetID -
func (u *User) GetID() string {
	return u.ID
}
