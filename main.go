package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-vgo/robotgo"
	"net/http"
	"os"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	//var init
	var (
		warn string
	)

	//get the flags from the command line
	unsafe := flag.Bool("unsafe", false, "disable safeguards")
	flag.Parse()

	//check the unsafe flag
	if *unsafe {
		warn = "Warning: You're running this program without safeguard, stay alert! (to stop the program press ctrl+c in your terminal) [press enter to continue] "
	} else {
		warn = "Warning: This program is intended to stop if you write anything on your keyboard, if you want to disable this, please use the -unsafe flag [press enter to continue] "
	}

	//alert the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Warning: we're not responsible for any sanction that you may receive using this program")
	fmt.Print(warn)
	if _, err := reader.ReadString('\n'); err != nil {
		return
	}

	//create a keyboard listener

	//We're good to go, start the loop
	for {
		//get the quote
		quote, err := GetQuote()

		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		//display the quote
		quote.Write()
	}
}

func (quote *Quote) Write() {
	robotgo.TypeStr(fmt.Sprintf("> %s", quote.Content))
	robotgo.KeyPress("enter")
	robotgo.TypeStrDelay(fmt.Sprintf("-%s", quote.Author), 2)
	robotgo.KeyPress("enter")
}

func GetQuote() (Quote, error) {
	// Get the quote
	response, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		return Quote{}, err
	}

	// Decode the response
	var quote Quote
	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		return Quote{}, err
	}

	//return the quote
	return quote, nil
}
