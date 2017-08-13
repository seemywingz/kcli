package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/martinlindhe/imgcat/lib"

	mtg "github.com/MagicTheGathering/mtg-sdk-go"
	gt "github.com/seemywingz/gtils"
)

const filePath = "/tmp/mtg/"

// Mtg :
func Mtg() {
	cards, err := mtg.NewQuery().Where(mtg.CardName, options.Mtg.Name).Where(mtg.CardSet, options.Mtg.Set).Where(mtg.CardSetName, options.Mtg.SetName).All()
	gt.EoE("Failed to get Cards", err)

	gt.Mkdir(filePath)

	for i := range cards {
		card := cards[i]
		id := fmt.Sprint(card.MultiverseId)
		fileName := filepath.Join(filePath + string(card.Set) + "." + card.Name + "." + id + ".png")
		if card.ImageUrl != "" { // there is a remote image
			if _, err := os.Stat(fileName); os.IsNotExist(err) { // the image is not cached
				gt.DownloadImage(card.ImageUrl, fileName)
				if options.Verbose {
					fmt.Println("Downloading Image URL:", card.ImageUrl)
					fmt.Println("            File Path:", fileName)
				}
			}
			imgcat.CatFile(fileName, os.Stdout)
		}
	}

}

// err := os.Remove(filePath)
// gt.EoE("Error Deleting Temp Dir", err)
