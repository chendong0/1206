package main

import "fmt"

func main()  {

getInfor()
Sex := "male"
fmt.Printf("%d",SexWeight)

}

func getInfor()  {
	var Name string
	fmt.Print("\n姓名:")
	fmt.Scanln(&Name)

	var Weight float64
	fmt.Print("体重(KG):")
	fmt.Scanln(&Weight)

	var Tall float64
	fmt.Print("身高(米):")
	fmt.Scanln(&Tall)

	var Age int
	fmt.Print("年龄:")
	fmt.Scanln(&Age)

	var SexWeight int
	var Sex string = "male"
	fmt.Println("性别(female/male):")
	fmt.Scanln(&Sex)

	if Sex == "male" { //==要区分理解=和==的区别和用法.
		SexWeight = 0
	} else {
		SexWeight = 1
	}


}
