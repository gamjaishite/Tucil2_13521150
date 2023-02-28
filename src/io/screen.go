package io

import "fmt"

// Initial screen display
func InitScreen() {
	welcomeString := "Closest Pair N-D Points"
	fmt.Printf("%s %*s\n", "+++", len(welcomeString)+4, "+++")
	fmt.Printf("+++ %s +++\n", welcomeString)
	fmt.Printf("%s %*s\n\n", "+++", len(welcomeString)+4, "+++")
}
