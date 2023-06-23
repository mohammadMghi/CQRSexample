package sqrs

import (
	sqrs "github.com/mohammadMghi/sqrs/entities"
	"github.com/mohammadMghi/sqrs/mysql"
)
 

type ProductRepository struct{
	mysql mysql.MysqlDB
}


func(u *ProductRepository) Save (Product sqrs.Product) error{

	// Save product in database 

	u.mysql.


}