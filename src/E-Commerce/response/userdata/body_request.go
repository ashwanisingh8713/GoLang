package userdata

type BodySelectedProductInput struct {
	UserId    string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type BodyUserIdInput struct {
	UserId string `json:"userId" binding:"required"`
}

type BodyActionOnIdInput struct {
	UserId     string `json:"userId" binding:"required"`
	ActionOnId string `json:"actionOnId" binding:"required"`
}

type BodyCreateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LoginMode string `json:"login_mode"`
}

type StatusMsg struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

type BodyAddressInput struct {
	UserId       string
	AddressType  string
	AddressLine1 string
	AddressLine2 string
	IsPreferred  bool
	Zip          string
	City         string
	State        string
	Country      string
	Mobile1      string
	Mobile2      string
}
