package main

import (
	"bufio"
	"fmt"
	"os"

	"baner_maker/convertor"
	"baner_maker/resize"
)

func main() {
	fmt.Println("		WH40K BANNER MAKER		")
	err := os.MkdirAll("images/output", os.ModePerm)
	if err != nil {
		fmt.Printf("Error while creating folders: %v\n", err)
	}
	err = os.MkdirAll("images/input", os.ModePerm)
	if err != nil {
		fmt.Printf("Error while creating folders: %v\n", err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n\nWhat would you like to make?\n\n1 - for banners\n2 - for badges\n0 - for break")
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Can not read")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		switch num {
		case 1:
			fmt.Println("Please write file name. Make sure it located in 'images/old/filename'. Do not forget the type of file(.jpg, .png, etc...)")
			if scanner.Scan() {
				fileName := scanner.Text()
				img := resize.Image_resize(fileName, 64, 96)
				if img == nil {
					fmt.Println("Can't transform image")
					continue
				}
				convertor.Tga_converter(img, fileName)
				continue
			}
		case 2:
			fmt.Println("Please write file name. Make sure it located in 'images/old/filename'. Do not forget the type of file(.jpg, .png, etc...)")
			if scanner.Scan() {
				fileName := scanner.Text()
				img := resize.Image_resize(fileName, 64, 64)
				if img == nil {
					fmt.Println("Can't transform image")
					continue
				}
				convertor.Tga_converter(img, fileName)
				continue
			}
		case 0:
			fmt.Println("Thank for using this program")
			return
		default:
			fmt.Println("Please, choose the correct option")
		}
	}

}
