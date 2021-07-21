package validate

type SignUpParam struct {
	Username   string `json:"username" validate:"SignUpParamUsernameValidate,min=1,max=20"`
	Password   string `json:"password" validate:"required,min=8,max=64"`
	RePassword string `json:"rePassword" validate:"required,eqfield=Password"`
}

type LoginParam struct {
	Username string `json:"username" validate:"required,min=1,max=20"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type ClimbTicketInfoParam struct {
	FromCity  string `json:"fromCity" validate:"required"`
	ToCity    string `json:"toCity" validate:"required"`
	StartDate string `json:"startDate" validate:"required"`
	EndDate   string `json:"endDate" validate:"required,gtefield=StartDate"`
}
