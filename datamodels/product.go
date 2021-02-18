package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"id" cons:"id"`
	ProductName  string `json:"productName" sql:"productName" cons:"productName"`
	ProductNum   int64  `json:"productNum" sql:"productNum" cons:"productNum"`
	ProductImage string `json:"productImage" sql:"productImage" cons:"productImage"`
	ProductUrl   string `json:"productUrl" sql:"productUrl" cons:"productUrl"`
}
