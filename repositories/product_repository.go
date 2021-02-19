package repositories

import (
	"con-system/common"
	"con-system/datamodels"
	"database/sql"
	"strconv"
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

// 删除
func (p *ProductManager) Delete(productID int64) bool{
	if err := p.Conn(); err != nil {
		return false
	}
	sql := "deletd from product where ID=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_, err = stmt.Exec(productID)
	if err != nil {
		return false
	}
	return true
}

// 更新
func (p *ProductManager) Update(product *datamodels.Product) error{
	if err := p.Conn(); err != nil {
		return err
	}
	sql := "update product set productName=?, productNum=?, productImage=?, productUrl=? where ID=" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}

//
func (p ProductManager) Select(productID int64) (productResult *datamodels.Product, err error){
	if err = p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}
	sql := "select * from" + p.table + "where ID =" + strconv.FormatInt(productID, 10)
	row, err := p.mysqlConn.Query(sql)
	if err != nil {
		return &datamodels.Product{}, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	productResult = &datamodels.Product{}
	common.DataToStructByTagSql(result,productResult)
	return

}

//
func (p *ProductManager) SelectAll() ([]*datamodels.Product, error) {
	return nil, nil
}