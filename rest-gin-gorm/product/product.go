package product

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model //내부적으로도 공통으로 사용하는 필드를 가지고 있음 (ex. ID, CreatedAt, etc)
	Code       string
	Price      uint
}
