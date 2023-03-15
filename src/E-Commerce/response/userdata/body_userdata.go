package userdata

type BodyWishlistInput struct {
	UserId    string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type BodyUserIdInput struct {
	UserId string `json:"userId" binding:"required"`
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
