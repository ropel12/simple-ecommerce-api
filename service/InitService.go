package service

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/ropel12/simple-ecommerce-api/repository"
)

type InitService struct {
	DB                    *sql.DB
	Validate              *validator.Validate
	UserRepository        repository.UserRepository
	ProdukRepostory       repository.ProductRepo
	TransactionRepository repository.TransactionRepo
	PaymentsRepository    repository.PaymentRepo
	CacheData             map[int]interface{}
}

func RunService(DB *sql.DB, validate *validator.Validate, user repository.UserRepository, produk repository.ProductRepo, transaction repository.TransactionRepo, payment repository.PaymentRepo) (ProductService, TransactionService, UserService_impl, PaymentService) {
	return &InitService{
			DB:                    DB,
			Validate:              validate,
			UserRepository:        user,
			ProdukRepostory:       produk,
			TransactionRepository: transaction,
			PaymentsRepository:    payment,
			CacheData:             make(map[int]interface{}),
		}, &InitService{
			DB:                    DB,
			Validate:              validate,
			UserRepository:        user,
			ProdukRepostory:       produk,
			TransactionRepository: transaction,
			PaymentsRepository:    payment,
			CacheData:             make(map[int]interface{}),
		}, &InitService{
			DB:                    DB,
			Validate:              validate,
			UserRepository:        user,
			ProdukRepostory:       produk,
			TransactionRepository: transaction,
			PaymentsRepository:    payment,
			CacheData:             make(map[int]interface{}),
		},
		&InitService{
			DB:                    DB,
			Validate:              validate,
			UserRepository:        user,
			ProdukRepostory:       produk,
			TransactionRepository: transaction,
			PaymentsRepository:    payment,
			CacheData:             make(map[int]interface{}),
		}

}
