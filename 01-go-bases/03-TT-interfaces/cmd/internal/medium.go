package product

type ConfigMedium struct {
	Price float64
}

func NewMedium(conf *ConfigMedium) (m *Medium) {
	m = &Medium{
		PriceProd: conf.Price,
	}
	return
}

type Medium struct {
	PriceProd float64
}

func (m Medium) Price() (price float64) {
	price = m.PriceProd + m.PriceProd*0.03
	return
}
