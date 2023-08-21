package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Emoji struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
	Category    string   `json:"category"`
	SubCategory string   `json:"subCategory"`
	Unicode     string   `json:"unicode"`
	Shortcode   string   `json:"shortcode"`
}

func GetEmojis() ([]Emoji, error) {
	homedir, err := os.UserHomeDir()

	filepath := fmt.Sprintf("%s/emojis.json", homedir)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	var emojis []Emoji

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&emojis)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return emojis, nil
}

func FindEmojisByTagAll(tag string) ([]Emoji, error) {
	emojis, err := GetEmojis()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var foundEmojis []Emoji

	for _, emoji := range emojis {
		for _, emojitag := range emoji.Tags {
			if emojitag == tag {
				foundEmojis = append(foundEmojis, emoji)
			}
		}
	}

	return foundEmojis, nil
}

func FindEmojisByTagOnlyOne(tag string) ([]Emoji, error) {
	emojis, err := GetEmojis()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var foundEmojis []Emoji

	for _, emoji := range emojis {
		uniCount := len(strings.Split(emoji.Unicode, " "))

		if uniCount == 1 {
			for _, emojitag := range emoji.Tags {
				if emojitag == tag {
					foundEmojis = append(foundEmojis, emoji)
				}
			}
		}
	}

	return foundEmojis, nil
}
