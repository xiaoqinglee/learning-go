package ch6

type BadIntSet []int //the slice is exposed to the outside world
type GoodIntSet struct {
	//the slice is invisible to the outside world
	intSet []int
}

type BadModel struct {
	internalStateField string
	//field that intruding store from outside world
	//might make the struct instance's internal states inconsistency
	ACertainField string
}
type GoodModel struct {
	internalStateField string
	aCertainField      string
}

func (g GoodModel) ACertainField() string { //Getter
	//do some update on internalStateField if needed
	return g.aCertainField
}
func (g GoodModel) SetACertainField(newValue string) { //Setter
	g.aCertainField = newValue
	//do some update on internalStateField
}
