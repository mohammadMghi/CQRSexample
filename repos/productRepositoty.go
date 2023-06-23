package sqrs

import (
 

	entities "github.com/mohammadMghi/sqrs/entities"
	sqrs "github.com/mohammadMghi/sqrs/entities"
	"github.com/mohammadMghi/sqrs/mysql"
)
 

type ProductRepository struct{
	mysql mysql.MysqlDB
}


func(u *ProductRepository) Save (Product sqrs.Product) error{

	// Save product in database 

	err := u.Save(Product)

	if err != nil{
		return err
	}

	return nil

}

func(u *ProductRepository) GetProduct () (*entities.Product , error){

	// Save product in database 

	err , product := u.mysql.GetProduct()

	if err != nil{
		return nil , err
	}

	return product , nil

}