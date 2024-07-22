package models

type Request struct{
	Account_id string `json:"account_id"`
	Reference string `json:"reference"`
	Amount float64 `json:"amount"`

}

type Response struct{
	Account_id string `json:"account_id"`
	Reference string `json:"reference"`
	Amount float64 `json:"amount"`
}