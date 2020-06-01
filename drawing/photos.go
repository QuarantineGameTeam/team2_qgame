package main

import (
	"os"
)

func main() {
	CreateCastlePhoto()
}

//CreateCastlePhoto function which draws castle
func CreateCastlePhoto() {

	// Read image from file that already exists
	existingImageFile, err := os.Open("castle.jpg")
	if err != nil {
		// Handle error
	}

	defer existingImageFile.Close()
}

//CreatEnemyPhoto  function which draws enemy
func CreatEnemyPhoto() {

	// Read image from file that already exists
	existingImageFile, err := os.Open("enemy.png")
	if err != nil {
		// Handle error
	}

	defer existingImageFile.Close()
}

//CreatFactoryPhoto  function which draws factory
func CreatFactoryPhoto() {

	// Read image from file that already exists
	existingImageFile, err := os.Open("candy_factory.png")
	if err != nil {
		// Handle error
	}

	defer existingImageFile.Close()
}

//CreatChestPhoto  function which draws chest
func CreatChestPhoto() {

	// Read image from file that already exists
	existingImageFile, err := os.Open("chest.png")
	if err != nil {
		// Handle error
	}

	defer existingImageFile.Close()
}
