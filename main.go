package main

import (
	"fmt"
	"github.com/XiaoqingLee/LearningGo/pkg/ch1"
	"github.com/biter777/countries"
)

func main() {

	fmt.Println(ch1.Hello())

	ch1.Echo1()

	countryChina := countries.China
	fmt.Printf("Country name in english: %v\n", countryChina)                   // Japan
	fmt.Printf("Country name in russian: %v\n", countryChina.StringRus())       // Япония
	fmt.Printf("Country ISO-3166 digit code: %d\n", countryChina)               // 392
	fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", countryChina.Alpha2())    // JP
	fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", countryChina.Alpha3())    // JPN
	fmt.Printf("Country IOC/NOC code: %v\n", countryChina.IOC())                // JPN
	fmt.Printf("Country FIFA code: %v\n", countryChina.FIFA())                  // JPN
	fmt.Printf("Country Capital: %v\n", countryChina.Capital())                 // Tokyo
	fmt.Printf("Country ITU-T E.164 call code: %v\n", countryChina.CallCodes()) // +81
	fmt.Printf("Country ccTLD domain: %v\n", countryChina.Domain())             // .jp
	fmt.Printf("Country UN M.49 region name: %v\n", countryChina.Region())      // Asia
	fmt.Printf("Country UN M.49 region code: %d\n", countryChina.Region())      // 142
	fmt.Printf("Country emoji/flag: %v\n\n", countryChina.Emoji())              // 🇯🇵

}
