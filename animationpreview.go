package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TheBitDrifter/bappa/blueprint/vector"
)

type AnimationConfig struct {
	Active            string             `json:"active"`
	AnimationPreviews []AnimationPreview `json:"animations"`
	GlobalOffset      vector.Two         ` json:"globalOffset"`
}

type AnimationPreview struct {
	Name           string     `json:"name"`
	RowIndex       int        `json:"rowIndex"`
	FrameCount     int        `json:"frameCount"`
	FrameWidth     int        `json:"frameWidth"`
	FrameHeight    int        `json:"frameHeight"`
	Speed          int        `json:"speed"`
	FreezeFrame    int        `json:"freezeframe"`
	PositionOffset vector.Two `json:"offset"`
}

func (config *AnimationConfig) Load(path string) error {
	jsonDataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading JSON file '%s': %v", path, err)
	}

	err = json.Unmarshal(jsonDataBytes, config)
	if err != nil {
		return err
	}
	fmt.Printf("\nSuccessfully loaded data from %s\n\n", path)
	fmt.Printf("Active Animation: %s\n", config.Active)
	fmt.Printf("Animations (%d):\n", len(config.AnimationPreviews))

	for i, anim := range config.AnimationPreviews {
		fmt.Printf("  Animation %d:\n", i+1)
		fmt.Printf("    Name: %s\n", anim.Name)
		fmt.Printf("    RowIndex: %d\n", anim.RowIndex)
		fmt.Printf("    FrameCount: %d\n", anim.FrameCount)
		fmt.Printf("    FrameWidth: %d\n", anim.FrameWidth)
		fmt.Printf("    Speed: %d\n", anim.Speed)
		fmt.Printf("    FreezeFrame: %d\n", anim.FreezeFrame)

		if config.Active == anim.Name {
			log.Println("Matched", anim.Name)

			isFrozen := anim.FreezeFrame != 0

			if isFrozen {
				PREVIEW_ANIMATION.FrameCount = anim.FreezeFrame
				PREVIEW_ANIMATION.Freeze = true
			} else {
				PREVIEW_ANIMATION.FrameCount = anim.FrameCount
				PREVIEW_ANIMATION.Freeze = false
			}

			PREVIEW_ANIMATION.Name = anim.Name
			PREVIEW_ANIMATION.RowIndex = anim.RowIndex
			PREVIEW_ANIMATION.FrameWidth = anim.FrameWidth
			PREVIEW_ANIMATION.FrameHeight = anim.FrameHeight
			PREVIEW_ANIMATION.Speed = anim.Speed
			PREVIEW_ANIMATION.PositionOffset = anim.PositionOffset
			break
		}

	}
	return nil
}
