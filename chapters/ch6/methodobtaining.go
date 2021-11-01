package ch6

import (
	"fmt"
)

type horseBody struct {
	Color string
	Size  string
}
type horse struct {
	horseBody
}

func (h *horse) GallopFast() {
	fmt.Printf("GallopFastAsHorse\n")
}
func (h *horse) EatGrass() {
	fmt.Printf("EatGrassAsHorse\n")
}

type donkeyBody struct {
	Color string
	Size  string
}
type donkey struct {
	donkeyBody
}

func (d *donkey) EasyToFarm() {
	fmt.Printf("EasyToFarmAsDonkey\n")
}
func (d *donkey) EatGrass() {
	fmt.Printf("EatGrassAsDonkey\n")
}

type Mule struct {
	horse
	donkey
}

func MethodAndFieldConflict() {
	mule := &Mule{
		horse: horse{
			horseBody{
				Color: "various",
				Size:  "big",
			},
		},
		donkey: donkey{
			donkeyBody{
				Color: "blackBackWhiteBelly",
				Size:  "small",
			},
		},
	}
	fmt.Printf("a Mule:\n")
	mule.GallopFast()
	mule.EasyToFarm()
	//mule.EatGrass() //Ambiguous reference 'EatGrass'
	mule.horse.EatGrass()
	mule.donkey.EatGrass()

	fmt.Printf("a Mule's body:\n")
	////Ambiguous reference 'Color' //Ambiguous reference 'Size'
	//fmt.Printf("color: %s size: %s\n", mule.Color, mule.Size)
	fmt.Printf("color: %s size: %s\n", mule.horse.Color, mule.horse.Size)
	fmt.Printf("color: %s size: %s\n", mule.horse.horseBody.Color, mule.horse.horseBody.Size)
	fmt.Printf("color: %s size: %s\n", mule.donkey.Color, mule.donkey.Size)
	fmt.Printf("color: %s size: %s\n", mule.donkey.donkeyBody.Color, mule.donkey.donkeyBody.Size)

}
