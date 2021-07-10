package validate

type SignUpParam struct {
	Username   string `json:"username" validate:"SignUpParamUsernameValidate"`
	Password   string `json:"password" validate:"required,min=8,max=64"`
	RePassword string `json:"re_password" validate:"required,eqfield=Password"`
}
