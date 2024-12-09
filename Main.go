// / My weight loss program
package main

import (
	"fmt"
)

// main is the function where it all begins
func main() {
	const lightSpeed = 299792  // km/s
	var distance = 56000000    // km
	const spaceXSpeed = 100800 //km/h
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 220*0.3783)
	fmt.Printf(" and I would be %v years old. \n", 33*365/687)
	fmt.Printf("My weight on the surface of %v is %v lbs. \n", "Earth", 220.0)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Blue Origin", 100)
	// How long does it take to get to mars?
	fmt.Println(distance/lightSpeed, "Seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "Seconds")
	fmt.Printf("%v Seconds\n", distance/lightSpeed)
	distance = 96500000
	fmt.Printf("It would take SpaceX %v Hours to reach mars\n", distance/spaceXSpeed)

}
