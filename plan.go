package main

import (
	"path/filepath"

	"github.com/TheBitDrifter/bappa/blueprint/client"
	"github.com/TheBitDrifter/bappa/blueprint/vector"
	"github.com/TheBitDrifter/bappa/combat"
	"github.com/TheBitDrifter/bappa/tteokbokki/spatial"
	"github.com/TheBitDrifter/bappa/warehouse"
)

func combatPreviewPlan(width, height int, sto warehouse.Storage) error {
	previewAnimArche, err := sto.NewOrExistingArchetype(
		spatial.Components.Position,
		client.Components.SpriteBundle,
		combat.Components.Attack,
		combat.Components.HurtBoxes,
	)
	if err != nil {
		return err
	}

	PREVIEW_SPRITE_SHEET := client.NewSpriteBundle().
		AddSprite(filepath.Base(SPRITE_SHEET_PATH), true).
		WithAnimations(PREVIEW_ANIMATION).
		WithOffset(vector.Two{X: PREVIEW_OFFSET.X, Y: PREVIEW_OFFSET.Y})

	_, err = previewAnimArche.GenerateAndReturnEntity(1,
		spatial.NewPosition(float64(width/2), float64(height/2)),
		PREVIEW_SPRITE_SHEET,
		PREVIEW_HURTBOXES,
		PREVIEW_HITBOX_ATTACK,
	)

	return err
}
