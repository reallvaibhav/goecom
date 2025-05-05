package usecase

type StatisticsUsecase interface {
	ProcessOrderCreated(data []byte)
	ProcessInventoryUpdated(data []byte)
}

type statisticsUsecase struct {
	repo Repository
}

func NewStatisticsUsecase(repo Repository) StatisticsUsecase {
	return &statisticsUsecase{repo: repo}
}

func (u *statisticsUsecase) ProcessOrderCreated(data []byte) {
	// Parse data and update statistics
}

func (u *statisticsUsecase) ProcessInventoryUpdated(data []byte) {
	// Parse data and update statistics
}
