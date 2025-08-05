package main

import (
	"embed"
	"flag"
	"log"
	"path/filepath"

	"github.com/TheBitDrifter/bappa/blueprint"
	"github.com/TheBitDrifter/bappa/blueprint/client"
	"github.com/TheBitDrifter/bappa/blueprint/vector"
	"github.com/TheBitDrifter/bappa/coldbrew"
	"github.com/TheBitDrifter/bappa/coldbrew/coldbrew_rendersystems"
	"github.com/TheBitDrifter/bappa/coldbrew/combat_rendersystems"
	"github.com/TheBitDrifter/bappa/combat"
)

var assets embed.FS

const (
	RESOLUTION_X       = 640
	RESOLUTION_Y       = 360
	MAX_SPRITES_CACHED = 100
	MAX_SCENES_CACHED  = 1
	MAX_SOUNDS_CACHED  = 1
)

const (
	SCENE_NAME   = "combatpreview"
	SCENE_WIDTH  = RESOLUTION_X
	SCENE_HEIGHT = RESOLUTION_Y
)

var (
	PREVIEW_JSON_PATH_ANIM    = "animation_preview.json"
	PREVIEW_JSON_PATH_ATTACK  = "attack_preview.json"
	PREVIEW_JSON_PATH_HURTBOX = "hurtbox_preview.json" // New path for hurtboxes
	SPRITE_SHEET_PATH         = "spritesheets/example_sheet.png"

	PREVIEW_ANIMATION     = client.AnimationData{}
	PREVIEW_HITBOX_ATTACK = combat.Attack{}
	PREVIEW_HURTBOXES     = combat.HurtBoxes{}
	PREVIEW_OFFSET        = vector.Two{}
)

func main() {
	flag.StringVar(&PREVIEW_JSON_PATH_ANIM, "jsonanim", PREVIEW_JSON_PATH_ANIM, "Path to the animation preview JSON file.")
	flag.StringVar(&PREVIEW_JSON_PATH_ATTACK, "jsonattacks", PREVIEW_JSON_PATH_ATTACK, "Path to the attack preview JSON file.")
	flag.StringVar(&PREVIEW_JSON_PATH_HURTBOX, "jsonhurtboxes", PREVIEW_JSON_PATH_HURTBOX, "Path to the hurtbox preview JSON file.")
	flag.StringVar(&SPRITE_SHEET_PATH, "sheet", SPRITE_SHEET_PATH, "Path to the sprite sheet PNG file.")
	flag.Parse()

	load()

	log.Println("Configuration after flag parsing and load():")
	log.Printf("  Animation Preview JSON: %s\n", PREVIEW_JSON_PATH_ANIM)
	log.Printf("  Attack Preview JSON:    %s\n", PREVIEW_JSON_PATH_ATTACK)
	log.Printf("  Hurtbox Preview JSON:   %s\n", PREVIEW_JSON_PATH_HURTBOX)
	log.Printf("  Sprite Sheet:           %s\n", SPRITE_SHEET_PATH)

	client := coldbrew.NewClient(
		RESOLUTION_X,
		RESOLUTION_Y,
		MAX_SPRITES_CACHED,
		MAX_SOUNDS_CACHED,
		MAX_SCENES_CACHED,
		assets,
	)

	client.SetTitle("combat-preview")
	client.SetResizable(true)
	client.SetLocalAssetPath(filepath.Dir(SPRITE_SHEET_PATH) + "/")

	err := client.RegisterScene(
		SCENE_NAME,
		SCENE_WIDTH,
		SCENE_HEIGHT,
		combatPreviewPlan,
		[]coldbrew.RenderSystem{
			combat_rendersystems.HitBoxRenderSystem{},
			combat_rendersystems.HurtBoxRenderSystem{},
		},
		[]coldbrew.ClientSystem{
			PreviewLoadSystem{},
		},
		[]blueprint.CoreSystem{},
	)
	if err != nil {
		log.Fatalf("Failed to register %v — %s", err, SCENE_NAME)
	}

	client.RegisterGlobalRenderSystem(
		coldbrew_rendersystems.GlobalRenderer{},
		&coldbrew_rendersystems.DebugRenderer{},
	)

	_, err = client.ActivateCamera()
	if err != nil {
		log.Fatalf("Failed to activate camera: %v", err)
	}

	if err := client.Start(); err != nil {
		log.Fatalf("Client exited with error: %v", err)
	}
}
