package models

type User struct {
  Id    string `json:"id,omitempty"`
  Name  string `json:"name"`
  Email string `json:"email"`
}
