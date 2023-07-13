package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Welcome to key-value pair commands:")
	fmt.Println("- SET key value - Set a key-value pair")
	fmt.Println("- GET key - Retrieve the value for a given key")
	fmt.Println("- DELET key - Delete a key-value key")
	fmt.Println("- EXIT - Terminate the program")
	data := make(map[string]string)

	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()

		if len(userInput) == 0 {
			continue
		}

		words := strings.Split(userInput, " ")
		myCMD := words[0:1]
		strCMD := strings.Join(myCMD, " ")

		if strCMD == "EXIT" {
			fmt.Print("Good bye")
			os.Exit(0)
		}
		myKey := words[1:2]
		strKey := strings.Join(myKey, " ")

		if strCMD == "GET" {
			if len(data[strKey]) == 0 {
				fmt.Println("Key " + strKey + " not found")
			} else {
				fmt.Println("Value for Key " + strKey + " : " + data[strKey])
			}
			continue
		}
		if strCMD == "DELETE" {
			if len(data[strKey]) == 0 {
				fmt.Println("Key " + strKey + " not found")
			} else {
				delete(data, strKey)
				fmt.Println("Key " + strKey + " name deleted successfully")
			}
			continue
		}

		myValue := words[2:3]
		strValue := strings.Join(myValue, " ")

		if strCMD == "SET" {
			data[strKey] = strValue
			fmt.Println("Key " + strKey + " name set successfully")
		}

	}

}
