package main

import (
	"flag"
	"fmt"

	"github.com/destinio/gomoji/data"
)

func main() {
	all := flag.String("all", "", "List all emojis")

	flag.Parse()

	if *all != "" {
		emojis, err := data.FindEmojisByTagAll(*all)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, emoji := range emojis {
			fmt.Println(emoji.Emoji, emoji.Unicode, emoji.Shortcode)
		}
	} else {
		if len(flag.Args()) == 0 {
			fmt.Println("Please provide a search term")
			return
		}
		input := flag.Args()[0]
		emojis, err := data.FindEmojisByTagOnlyOne(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, emoji := range emojis {
			fmt.Println(emoji.Emoji, emoji.Unicode, emoji.Shortcode)
		}
	}
}
