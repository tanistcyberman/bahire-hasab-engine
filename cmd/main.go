package main

import (
	"fmt"

	"bahire-hasab-engine/internal/bahirehasab"
)

func main() {

	year := 2018

	fmt.Println("Year:", year)
	fmt.Println("Wenber:", bahirehasab.Wenber(year))
	fmt.Println("Abekte:", bahirehasab.Abekte(year))
	fmt.Println("Metqi:", bahirehasab.Metqi(year))

	fasika := bahirehasab.Fasika(year)

	fmt.Println("fashika:", fasika.Day , fasika.Month)


}