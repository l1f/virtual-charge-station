package app

import (
	"github.com/l1f/virtual-charge-station/internal/data"
	"go.uber.org/zap"
)

type App struct {
	Logger *zap.SugaredLogger
	Data   *data.Data
}
