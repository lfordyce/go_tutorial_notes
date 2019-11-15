package interview

type WaterBottle interface {
	Liters() int
}

type plasticWaterBottles struct {
	liters int
}

func (b *plasticWaterBottles) Liters() int {
	return b.liters
}

func NewWaterBottle() WaterBottle {
	return &plasticWaterBottles{liters: 1}
}
