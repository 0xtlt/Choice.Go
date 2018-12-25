package main

import (
	"./makeChoice"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main(){
	var myMoney uint16

	reader := bufio.NewReader(os.Stdin)
	myMoney = AskmyMoney(reader)

	Choices := AskObjects(reader, []makeChoice.Choice{})

	oneStep := makeChoice.MakeChoice(Choices, myMoney)
	secStep := makeChoice.WhatPurchase(oneStep, myMoney)

	fmt.Println("With "+strconv.Itoa(int(myMoney))+"€")
	secStep.Say()

	fmt.Print("Press enter to exit")
	leave, _ := reader.ReadString('\n')
	fmt.Println(leave)
}

func AskmyMoney(e *bufio.Reader) uint16 {
	fmt.Print("How much do you have to spend : ")
	text, _ := e.ReadString('\n')
	tmp, err := strconv.Atoi(text[0:(len(text) - 2)])
	if err != nil {
		fmt.Println("sorry, I encountered an error :")
		fmt.Println("it's not a number")
		return AskmyMoney(e)
	} else {
		return uint16(tmp)
	}
}

func AskInt(e *bufio.Reader) uint16 {
	text, _ := e.ReadString('\n')
	tmp, err := strconv.Atoi(RemoveRN(text))
	if err != nil {
		fmt.Println("sorry, I encountered an error :")
		fmt.Println("it's not a number")
		fmt.Print("Write a number : ")
		return AskInt(e)
	} else {
		return uint16(tmp)
	}
}

func AskInt1And10(e *bufio.Reader) uint16 {
	text, _ := e.ReadString('\n')
	tmp, err := strconv.Atoi(RemoveRN(text))
	if err != nil {
		fmt.Println("sorry, I encountered an error :")
		fmt.Println("it's not a number")
		fmt.Print("Retry : ")
		return AskInt1And10(e)
	} else if tmp <= 10 && tmp >= 1 {
		return uint16(tmp)
	} else {
		fmt.Println("sorry, I encountered an error (0-10) : ")
		return AskInt1And10(e)
	}
}

func AskObjects(e *bufio.Reader, last []makeChoice.Choice) []makeChoice.Choice{
	var Choices []makeChoice.Choice
	var tmp = makeChoice.Choice{
		Name: "",
		Price: 0,
		Desire: 0,
		MakeMoney: false,
		SimplifyLife: false,
		Transportable: false,
		Gadget: false,
		Fun: false,
	}
	fmt.Println("------")

	fmt.Print("Name of the object : ")
	objectName, _ := e.ReadString('\n')
	tmp.Name = RemoveRN(objectName)

	fmt.Print("Price of the object : ")
	tmp.Price = AskInt(e)

	fmt.Print("On a scale of 1 to 10, write down your desire to have it, 10 being the most desired :")
	tmp.Desire = uint8(AskInt1And10(e))

	fmt.Print("It helps you earn money [Y/n] : ")
	makeMoney, _ := e.ReadString('\n')
	makeMoney = RemoveRN(makeMoney)
	if makeMoney == "Y" || makeMoney == "y" || makeMoney == "1" {
		tmp.MakeMoney = true
	}

	fmt.Print("It simplifies your life [Y/n] : ")
	simply, _ := e.ReadString('\n')
	simply = RemoveRN(simply)
	if simply == "Y" || simply == "y" || simply == "1" {
		tmp.SimplifyLife = true
	}

	fmt.Print("It's transportable [Y/n] : ")
	transportable, _ := e.ReadString('\n')
	transportable = RemoveRN(transportable)
	if transportable == "Y" || transportable == "y" || transportable == "1" {
		tmp.Transportable = true
	}

	fmt.Print("It's a Gadget [Y/n] : ")
	gadget, _ := e.ReadString('\n')
	gadget = RemoveRN(gadget)
	if gadget == "Y" || gadget == "y" || gadget == "1" {
		tmp.Gadget = true
	}

	fmt.Print("He can entertain you [Y/n] : ")
	fun, _ := e.ReadString('\n')
	fun = RemoveRN(fun)
	if fun == "Y" || fun == "y" || fun == "1" {
		tmp.Fun = true
	}

	Choices = append(Choices, tmp)
	for _, element := range last {
		Choices = append(Choices, element)
	}

	fmt.Print("Do you have another object to add [Y/n] : ")
	another, _ := e.ReadString('\n')
	another = RemoveRN(another)
	if another == "Y" || another == "y" || another == "1" {
		return AskObjects(e, Choices)
	} else {
		fmt.Println("-------")
		return Choices
	}
}

func RemoveRN(e string) string{
	re := regexp.MustCompile(`\r?\n`)
	result := re.ReplaceAllString(e, "")

	return result
}