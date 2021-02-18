package repositories

import (
	"con-system/common"
	"con-system/datamodels"
	"database/sql"
)

// 第一步：开发对应的接口
// 第二步：实现定义的接口

type IProduct interface {
	// 连接数据库
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	Select(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct{
	return &ProductManager{ table: table, mysqlConn: db}
}

// 数据库连接
func (p *ProductManager) Conn() error{
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return nil
}

// 插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error){
	if err := p.Conn(); err == nil {
		return
	}
	sql := "insert product set productName=?, productNum=?, productImage=?, productUrl=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
