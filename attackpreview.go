package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TheBitDrifter/bappa/combat"
	"github.com/TheBitDrifter/bappa/tteokbokki/spatial"
)

type AttackConfig struct {
	Active         string          `json:"active"`
	AttackPreviews []AttackPreview `json:"attacks"`
}

type AttackPreview struct {
	Name        string         `json:"name"`
	Speed       int            `json:"speed"`
	Boxes       [][]PreviewBox `json:"boxes"`
	FreezeFrame int            `json:"freezeFrame"`
	Length      int            `json:"length"`
}

type PreviewBox struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

func (config *AttackConfig) Load(path string) error {
	jsonDataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading JSON file '%s': %v", path, err)
	}

	err = json.Unmarshal(jsonDataBytes, config)
	if err != nil {
		return err
	}
	fmt.Printf("\nSuccessfully loaded data from %s\n\n", path)
	fmt.Printf("Active Attack: %s\n", config.Active)
	fmt.Printf("Attack Frames Length (%d):\n", len(config.AttackPreviews))

	for i, attack := range config.AttackPreviews {
		fmt.Printf("  Attack %d:\n", i+1)
		fmt.Printf("    Name: %s\n", attack.Name)

		if config.Active == attack.Name {
			PREVIEW_HITBOX_ATTACK.Speed = attack.Speed
			PREVIEW_HITBOX_ATTACK.Length = len(attack.Boxes)

			if attack.FreezeFrame != 0 {
				boxes := [][]PreviewBox{
					attack.Boxes[attack.FreezeFrame-1],
				}
				attack.Boxes = boxes
				PREVIEW_HITBOX_ATTACK.Length = 1
				PREVIEW_HITBOX_ATTACK.Boxes = [15]combat.HitBoxes{}
			}

			for i, boxes := range attack.Boxes {
				for l, box := range boxes {
					PREVIEW_HITBOX_ATTACK.Boxes[i][l] = combat.HitBox(spatial.NewRectangle(float64(box.W), float64(box.H)))
					PREVIEW_HITBOX_ATTACK.BoxesPositionOffsets[i][l].X = float64(box.X)
					PREVIEW_HITBOX_ATTACK.BoxesPositionOffsets[i][l].Y = float64(box.Y)
				}
			}
			break
		}

	}

	return nil
}
