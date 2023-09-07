package product

type ConfigLarge struct {
	Price float64
}

func NewLargue(conf *ConfigLarge) (l *Large) {
	l = &Large{
		PriceProd: conf.Price,
	}
	return
}

type Large struct {
	PriceProd float64
}

func (l Large) Price() (price float64) {
	price = l.PriceProd + l.PriceProd*0.06 + 2500.0
	return
}
