package queries

import(

	entities "github.com/mohammadMghi/sqrs/entities"
)
type Query interface{
	Handler() entities.Product
}