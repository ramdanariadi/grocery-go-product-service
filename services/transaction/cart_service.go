package transaction

import (
	"go-tunas/customresponses/product"
)

type CartService interface {
	FindById(id string) []product.CartResponse
	Save(userId string, produtId string, total int) bool
	Delete(userId string, produtId string) bool
}
