package request

type UserSignInRequest struct {
  Email     string `json:"email"`
  Password  string `json:"password"`
}

type UserEntryRequest struct {
  Id       int64  `json:"id"`
  UserName string `json:"userName"`
  Des      string `json:"des"`
}

type UserDeleteRequest struct {
  Id int64 `json:"id"`
}
