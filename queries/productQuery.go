package queries
import(

	entities "github.com/mohammadMghi/sqrs/entities"
	repo "github.com/mohammadMghi/sqrs/repos"
)
type ProductQuery struct{
	query Query
	repo repo.ProductRepository
}


func(p *ProductQuery) Handler () (*entities.Product, error){

	product , err := p.repo.GetProduct()

	if err != nil{
		return nil, err
	}

	return product , nil 
}