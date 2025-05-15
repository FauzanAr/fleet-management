package fleet

import (
	"context"
)

type Repository interface {
	GetFleet(context.Context) error
}