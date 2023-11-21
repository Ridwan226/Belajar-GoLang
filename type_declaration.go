package main

import "fmt"


func main() {
	type NoKTP string
	
	var ktpRdw NoKTP = "123123123"
	
	
	var contoh string = "456456"
	var contohKtp NoKTP = NoKTP(contoh)
	
	fmt.Println(ktpRdw)
	fmt.Println(contohKtp)
	
	
}