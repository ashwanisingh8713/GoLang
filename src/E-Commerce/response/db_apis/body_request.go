package db_apis

type BodySelectedProductInput struct {
	UserId    string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type BodyUserIdInput struct {
	UserId string `json:"userId" binding:"required"`
}

type BodyUserLoginInput struct {
	EmailOrMobile string `json:"EmailOrMobile" binding:"required"`
	Password      string `json:"password" binding:"required"`
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
	UserId       string `json:"userId"`
	AddressType  string `json:"addressType"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	IsPreferred  bool   `json:"isPreferred"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Mobile1      string `json:"mobile1"`
	Mobile2      string `json:"mobile2"`
}

type BodyAddressUpdate struct {
	AddressId    string `json:"addressId"`
	UserId       string `json:"userId"`
	AddressType  string `json:"addressType"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	IsPreferred  bool   `json:"isPreferred"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Mobile1      string `json:"mobile1"`
	Mobile2      string `json:"mobile2"`
}

type BodyNameCreate struct {
	Name string `json:"name"`
}

type BodySubCategoryCreate struct {
	Name string `json:"name"`
	CategoryId string `json:"categoryId"`
	Description string `json:"description"`

}


type BodyById struct {
	ID string `json:"id"`
}
