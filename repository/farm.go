package repository

type FarmData struct {
	ID   string
	Name string
}

type FarmRepository interface {
	GetFarm(id string) *FarmData
}

type FakeFarmRepository struct {
	// TODO put state stuff in here
}

func (*FakeFarmRepository) GetFarm(id string) *FarmData {
	return nil
}
