package main

import (
	"fmt"
)

func main() {
	//arrays are not generally used in golang, instead slices are used
	var bballgoatts [10]string
	bballgoatts[0] = "Michael Jordan"
	bballgoatts[1] = "Larry Bird"
	bballgoatts[2] = "LeBron James"
	bballgoatts[3] = "Kareem Abdul-Jabbar"
	bballgoatts[4] = "Kobe Bryant"
	bballgoatts[5] = "Tim Duncan"
	bballgoatts[6] = "Big Daddy Shaq"
	bballgoatts[7] = "Jerry West"
	bballgoatts[8] = "Manute Ginobli"

	for i := 0; i < len(bballgoatts); i++ {
		fmt.Println(bballgoatts[i])
	}
	var greatespg[10] string;
	greatespg[0] = "Stephen curry";
	greatespg[1] = "John stockton";
	greatespg[2] = "Magic Johnson";
	greatespg[2] = "Isaiah Thomas";
	greatespg[3] = "Chris paul";
	greatespg[4] = "Oscar Robertson";
	greatespg[5] = "steve nash";
	greatespg[6] = "Jason kidd";
	greatespg[7] = "Walt Frazier";
	greatespg[8] = "Russel Westbrook";
	for i := 0; i < len(greatespg); i++ {
		fmt.Println(greatespg[i]);
	}
}
