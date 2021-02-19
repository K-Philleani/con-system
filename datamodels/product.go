package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" cons:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" cons:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" cons:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" cons:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" cons:"ProductUrl"`
}
