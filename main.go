package main

import "fmt"

func main() {

	page, err := GetPage("1111", -1, false, true)
	if err != nil {
		fmt.Println(err)
	}
	content, err := page.GetSectionTreeList()
	// Get the content of the page
	//content, err := page.GetContent()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("This is the page content: %v\n", content)
}
