package request

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"passoword"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}
