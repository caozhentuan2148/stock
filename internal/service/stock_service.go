package service

import (
	"fmt"
)

type StockService struct {
}

func NewStockService() *StockService {
	return &StockService{}
}

func (s *StockService) AnalyzeStock(symbol string) (string, error) {
	return fmt.Sprintf("正在分析股票: %s", symbol), nil
}
