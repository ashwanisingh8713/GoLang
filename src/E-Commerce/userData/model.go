package userData

type CreateWishlistInput struct {
	UserId    string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type CreateUserIdInput struct {
	UserId string `json:"userId" binding:"required"`
}

type StatusMsg struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}
