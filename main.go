package main

import "fmt"

func main() {
	appleFactory, _ := GetStoreFactory("Apple")
	samsungFactory, _ := GetStoreFactory("Samsung")

	applePhone := appleFactory.sellPhone()
	appleLaptop := appleFactory.sellLaptop()
	appleWatch := appleFactory.sellWatch()

	samsungPhone := samsungFactory.sellPhone()
	samsungLaptop := samsungFactory.sellLaptop()
	samsungWatch := samsungFactory.sellWatch()

	printPhoneDetails(applePhone)
	printLaptopDetails(appleLaptop)
	printWatchDetails(appleWatch)

	printPhoneDetails(samsungPhone)
	printLaptopDetails(samsungLaptop)
	printWatchDetails(samsungWatch)

}

func printPhoneDetails(p IPhone) {
	fmt.Printf("ram: %d", p.getRam())
	fmt.Println()
	fmt.Printf("color: %s", p.getColor())
	fmt.Println()
}
func printLaptopDetails(l ILaptop) {
	fmt.Printf("ram: %d", l.getRam())
	fmt.Println()
	fmt.Printf("color: %s", l.getColor())
	fmt.Println()
}
func printWatchDetails(w IWatch) {
	fmt.Printf("ram: %d", w.getRam())
	fmt.Println()
	fmt.Printf("color: %s", w.getColor())
	fmt.Println()
}
