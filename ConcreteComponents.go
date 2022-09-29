package Decorator

type HouseBlend struct {
	description string
}

func (h *HouseBlend) cost() float32 {
	return 0.85
}
func (h *HouseBlend) getDescription() string {
	return h.description
}
func newHouseBlend() *HouseBlend {
	houseBlend := new(HouseBlend)
	houseBlend.description = "house blend"
	return houseBlend
}

type DаrkRоаst struct {
	description string
}

func (d *DаrkRоаst) cost() float32 {
	return 2.33
}
func (d *DаrkRоаst) getDescription() string {
	return d.description
}
func newDarkRoast() *DаrkRоаst {
	darkRoast := new(DаrkRоаst)
	darkRoast.description = "dark roast"
	return darkRoast
}

type Espresso struct {
	description string
}

func (e *Espresso) cost() float32 {
	return 1.99
}
func (e *Espresso) getDescription() string {
	return e.description
}
func newEspresso() *Espresso {
	espresso := new(Espresso)
	espresso.description = "espresso"
	return espresso
}

type Decaf struct {
	description string
}

func (de *Decaf) cost() float32 {
	return 1.54
}
func (de *Decaf) getDescription() string {
	return de.description
}
func newDecaf() *Decaf {
	decaf := new(Decaf)
	decaf.description = "decaf"
	return decaf
}
