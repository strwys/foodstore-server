package service

import (
	"context"
	"time"

	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/repository"
	"github.com/cecepsprd/foodstore-server/utils/logger"
)

type InvoiceService interface {
	Read(ctx context.Context, orderid string) (model.Invoice, error)
}

type invoice struct {
	invoiceRepository repository.InvoiceRepository
	contextTimeout    time.Duration
}

func NewInvoiceService(
	invoiceRepository repository.InvoiceRepository,
	timeout time.Duration) InvoiceService {
	return &invoice{
		invoiceRepository: invoiceRepository,
		contextTimeout:    timeout,
	}
}

func (s *invoice) Read(ctx context.Context, orderid string) (model.Invoice, error) {
	invoice, err := s.invoiceRepository.Read(ctx, orderid)
	if err != nil {
		logger.Log.Error(err.Error())
		return model.Invoice{}, err
	}

	return invoice, nil
}
