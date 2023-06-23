package sqrs

import entities "github.com/mohammadMghi/sqrs/entities"

type Command interface{
	Handler () entities.Product
}