package mocks

import (
	"context"
	"math/rand"
	"time"

	"github.com/N1ktarchik/FinOps-Daemon/internal/core/errors"
	"github.com/N1ktarchik/FinOps-Daemon/internal/domain"
)

type MockBillingProvider struct {
}

func (m *MockBillingProvider) GetBalance(ctx context.Context, creds domain.Credentials) (float64, error) {
	time.Sleep(100 * time.Millisecond)

	return 1, nil
}

func (m *MockBillingProvider) TopUpBalance(ctx context.Context, creds domain.Credentials, amount float64, idempotencyKey string) error {
	num := rand.Int() % 11 //рандомное число от 0 до 10

	switch num {
	case 0, 1:
		time.Sleep(10 * time.Second)
		return context.DeadlineExceeded
	case 2:
		return errors.TopUpBalanceMockError()
	default:
		return nil

	}
}
