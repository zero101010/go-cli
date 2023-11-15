package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	channel := flag.String("channelName", "Huncoding", "Nome do canal")
	flag.Parse()
	os.Mkdir(fmt.Sprintf("%s", *channel), 077)
	fmt.Print(*channel)

}
