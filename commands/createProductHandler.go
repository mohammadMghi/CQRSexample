package sqrs

import (
	"log"

	product "github.com/mohammadMghi/sqrs/entities"
	repo "github.com/mohammadMghi/sqrs/repos"
)
type CreateProductHandler struct{
	command Command
	productRepository repo.ProductRepository
}


func(c *CreateProductHandler) Handler (command product.Product){


	err := c.productRepository.Save(command) ; if err != nil{
		log.Fatal(err.Error())
		return 
	}
}