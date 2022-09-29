package Decorator

type Milk struct {
	description string
	beverage    iBeverage
}

func newMilk(beverage iBeverage) *Milk {
	milk := new(Milk)
	milk.beverage = beverage
	return milk
}
func (m *Milk) cost() float32 {
	return m.beverage.cost() + 20
}
func (m *Milk) getDescription() string {
	return m.beverage.getDescription() + ", milk"
}

type Mоcha struct {
	description string
	beverage    iBeverage
}

func (mo *Mоcha) cost() float32 {
	return mo.beverage.cost() + 25
}
func (mo *Mоcha) getDescription() string {
	return mo.beverage.getDescription() + ", add mocha"
}
func newMocha(beverage iBeverage) *Mоcha {
	mоcha := new(Mocha)
	mоcha.beverage = beverage
	return mоcha
}

type Soy struct {
	description string
	beverage    iBeverage
}

func (s *Soy) cost() float32 {
	return s.beverage.cost() + 7
}
func (s *Soy) getDescription() string {
	return s.beverage.getDescription() + ", add soy"
}
func newSoy(beverage iBeverage) *Soy {
	soy := new(Soy)
	soy.beverage = beverage
	return soy
}

type Whip struct {
	beverage iBeverage
}

func (w *Whip) cost() float32 {
	return w.beverage.cost() + 10
}
func (w *Whip) getDescription() string {
	return w.beverage.getDescription() + ", add whip"
}
func newWhip(beverage iBeverage) *Whip {
	whip := new(Whip)
	whip.beverage = beverage
	return whip
}
