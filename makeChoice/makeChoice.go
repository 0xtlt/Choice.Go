package makeChoice

import (
	"fmt"
	"sort"
	"strconv"
)

type Choice struct {
	Name string
	Price uint16
	Desire uint8
	MakeMoney bool
	SimplifyLife bool
	Transportable bool
	Gadget bool
	Fun bool
}

type Valuable struct {
	Choice Choice
	Points uint8
}

type ByPoints []Valuable
func (a ByPoints) Len() int           { return len(a) }
func (a ByPoints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool { return a[i].Points > a[j].Points }

type Result struct {
	Valuable []Valuable
	MoneyRemaining uint16
}

func MakeChoice(e[] Choice, money uint16) []Valuable{
	MoneyAvailable := money

	var Val []Valuable

	for _, element := range e {
		if element.Price < MoneyAvailable {

			var tmpVal = Valuable{
				Choice: element,
				Points: 0x00,
			}

			if element.MakeMoney {
				tmpVal.Points += 5
			}

			if element.SimplifyLife {
				tmpVal.Points += 4
			}

			if element.Transportable {
				tmpVal.Points += 3
			}

			if element.Gadget {
				tmpVal.Points += 2
			}

			if element.Fun {
				tmpVal.Points++
			}

			tmpVal.Points *= element.Desire


			Val = append(Val, tmpVal)
		}
	}

	sort.Sort(ByPoints(Val))
	return Val
}

func WhatPurchase(e []Valuable, money uint16) Result{
	var tmp []Valuable
	tmpMoney := money

	for _, element := range e {
		if tmpMoney > element.Choice.Price {
			tmpMoney -= element.Choice.Price
			tmp = append(tmp, element)
		}
	}

	return Result{
		Valuable: tmp,
		MoneyRemaining: tmpMoney,
	}
}

func (e Result) Say(){
	fmt.Println("Results")
	for index, element := range e.Valuable {
		fmt.Println("The "+strconv.Itoa(index)+` - "`+element.Choice.Name+`" at `+strconv.Itoa(int(element.Choice.Price))+`€ | `+strconv.Itoa(int(element.Points))+"pts")
	}
	fmt.Println("At the end, you will have "+strconv.Itoa(int(e.MoneyRemaining))+"€ left")
}