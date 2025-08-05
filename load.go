package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TheBitDrifter/bappa/coldbrew"
	"github.com/TheBitDrifter/bappa/combat"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PreviewLoadSystem struct{}

func (PreviewLoadSystem) Run(cli coldbrew.LocalClient, scene coldbrew.Scene) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		load()
		scene.Reset()
		coldbrew.ForceSetTick(0)
	}

	return nil
}

func load() {
	animConfig := AnimationConfig{}
	if err := animConfig.Load(PREVIEW_JSON_PATH_ANIM); err != nil {
		log.Fatalf("Error loading animation JSON from '%s': %v", PREVIEW_JSON_PATH_ANIM, err)
	}
	PREVIEW_OFFSET = animConfig.GlobalOffset
	log.Println(PREVIEW_OFFSET, "PREVIEWOFFSET")
	log.Println(animConfig.GlobalOffset, "config")

	attackConfig := AttackConfig{}
	if err := attackConfig.Load(PREVIEW_JSON_PATH_ATTACK); err != nil {
		log.Fatalf("Error loading attack JSON from '%s': %v", PREVIEW_JSON_PATH_ATTACK, err)
	}

	// HURTBOXES YOLO BELOW:
	hurtboxDataBytes, err := os.ReadFile(PREVIEW_JSON_PATH_HURTBOX)
	if err != nil {
		log.Fatalf("Error reading hurtbox JSON file '%s': %v", PREVIEW_JSON_PATH_HURTBOX, err)
	}

	var tempBoxSlice []jsonHurtboxDef
	if err := json.Unmarshal(hurtboxDataBytes, &tempBoxSlice); err != nil {
		log.Fatalf("Error unmarshalling hurtbox JSON into temp slice: %v", err)
	}

	var finalHurtBoxArray combat.HurtBoxes

	if len(tempBoxSlice) > len(finalHurtBoxArray) {
		log.Fatalf(
			"Error: JSON defines %d hurtboxes, but the maximum allowed is %d",
			len(tempBoxSlice),
			len(finalHurtBoxArray),
		)
	}

	for i, boxDef := range tempBoxSlice {
		finalHurtBoxArray[i] = combat.NewHurtBox(boxDef.W, boxDef.H, boxDef.X, boxDef.Y)
	}

	PREVIEW_HURTBOXES = finalHurtBoxArray

	fmt.Printf("\nSuccessfully loaded and parsed hurtbox data from %s\n\n", PREVIEW_JSON_PATH_HURTBOX)
}
