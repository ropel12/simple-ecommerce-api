package service

import (
	"context"

	"github.com/ropel12/simple-ecommerce-api/model"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, request model.TransactionRequest, csid int64) error
	UpdateTmpTransaction(ctx context.Context, request model.TempUpdateTransactionRequest) error
	DeleteTmpTransaction(ctx context.Context, idtemptrx int64, cusid int64) error
	FindAllTmpTransactionCustomer(ctx context.Context, cusid int64) ([]*model.TempTransaction, error)
	FindAllTransactionCustomer(ctx context.Context) ([]model.TransactionAdmin, error)
	FindTrxByTransactionid(ctx context.Context, trxid int64) (*model.TransactionAdmin, error)
	FindAllTransactionByStatus(ctx context.Context, status int, cusid int) ([]model.TransactionCus, error)
	FindAllTransactionById(ctx context.Context, cusid int) ([]model.TransactionCus, error)
	FindAllTrxByTransactionid(ctx context.Context, trxid int64, cusid int64) (*model.TransactionCus, error)
	InsertTmpTransaction(ctx context.Context, req model.TempTransactionRequest, csid int64) error
	UploadProof(ctx context.Context, proof string, trxid string) error
}
