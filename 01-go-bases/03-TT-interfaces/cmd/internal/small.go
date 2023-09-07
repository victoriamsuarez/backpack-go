package product

type ConfigSmall struct {
	Price float64
}

func NewSmall(conf *ConfigSmall) (s *Small) {
	s = &Small{
		PriceProd: conf.Price,
	}
	return
}

type Small struct {
	PriceProd float64
}

func (s Small) Price() (price float64) {
	return s.PriceProd
}
