package inventory

import (
	"fmt"

	"tracker/preferences"
	"tracker/save"
	"tracker/tappable_icons"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

type InventoryIcons struct {
	bowTapIcon                     *tappable_icons.TappableIcon
	bowNonProgressiveTapIcon       *tappable_icons.TappableIcon
	blueBoomerangTapIcon           *tappable_icons.TappableIcon
	redBoomerangTapIcon            *tappable_icons.TappableIcon
	hookshotTapIcon                *tappable_icons.TappableIcon
	bombTapIcon                    *tappable_icons.TappableIcon
	mushroomTapIcon                *tappable_icons.TappableIconWithIcon
	magicPowderTapIcon             *tappable_icons.TappableIconWithIcon
	fireRodTapIcon                 *tappable_icons.TappableIcon
	iceRodTapIcon                  *tappable_icons.TappableIcon
	bombosTapIcon                  *tappable_icons.TappableIconWithString
	etherTapIcon                   *tappable_icons.TappableIconWithString
	quakeTapIcon                   *tappable_icons.TappableIconWithString
	lampTapIcon                    *tappable_icons.TappableIcon
	hammerTapIcon                  *tappable_icons.TappableIcon
	fluteTapIcon                   *tappable_icons.TappableIconWithIcon
	shovelTapIcon                  *tappable_icons.TappableIconWithIcon
	bugNetTapIcon                  *tappable_icons.TappableIcon
	bookOfMudoraTapIcon            *tappable_icons.TappableIcon
	bottleTotalTapIcon             *tappable_icons.TappableIconWithCenteredNum
	bottle1TapIcon                 *tappable_icons.TappableIconCycled
	bottle2TapIcon                 *tappable_icons.TappableIconCycled
	bottle3TapIcon                 *tappable_icons.TappableIconCycled
	bottle4TapIcon                 *tappable_icons.TappableIconCycled
	caneOfSomariaTapIcon           *tappable_icons.TappableIcon
	caneOfByrnaTapIcon             *tappable_icons.TappableIcon
	magicCapeTapIcon               *tappable_icons.TappableIcon
	magicMirrorTapIcon             *tappable_icons.TappableIcon
	pseudoPegasusBootsTapIcon      *tappable_icons.TappableIcon
	pegasusBootsTapIcon            *tappable_icons.TappableIcon
	glovesTapIcon                  *tappable_icons.TappableIcon
	flippersTapIcon                *tappable_icons.TappableIcon
	moonPearlTapIcon               *tappable_icons.TappableIcon
	swordTapIcon                   *tappable_icons.TappableIconVariedSize
	shieldTapIcon                  *tappable_icons.TappableIconVariedSize
	mailTapIcon                    *tappable_icons.TappableIcon
	halfMagicTapIcon               *tappable_icons.TappableIcon
	heartPieceTapIcon              *tappable_icons.TappableIconCycled
	ganonGoalTapIcon               *tappable_icons.TappableIconWithBottomCenteredString
	pedestalGoalTapIcon            *tappable_icons.TappableIcon
	triforceGoalTapIcon            *tappable_icons.TappableIcon
	ganonTowerGoalTapIcon          *tappable_icons.TappableIcon
	preferencesFile                *preferences.PreferencesFile
	saveFile                       *save.SaveFile
	undoRedo				*undo_redo.UndoRedoStacks
	bowTapContainer                *fyne.Container
	bowNonProgressiveTapContainer  *fyne.Container
	blueBoomerangTapContainer      *fyne.Container
	redBoomerangTapContainer       *fyne.Container
	hookshotTapContainer           *fyne.Container
	bombTapContainer               *fyne.Container
	mushroomTapContainer           *fyne.Container
	magicPowderTapContainer        *fyne.Container
	fireRodTapContainer            *fyne.Container
	iceRodTapContainer             *fyne.Container
	bombosTapContainer             *fyne.Container
	etherTapContainer              *fyne.Container
	quakeTapContainer              *fyne.Container
	lampTapContainer               *fyne.Container
	hammerTapContainer             *fyne.Container
	fluteTapContainer              *fyne.Container
	shovelTapContainer             *fyne.Container
	bugNetTapContainer             *fyne.Container
	bookOfMudoraTapContainer       *fyne.Container
	bottleTotalTapContainer        *fyne.Container
	bottle1TapContainer            *fyne.Container
	bottle2TapContainer            *fyne.Container
	bottle3TapContainer            *fyne.Container
	bottle4TapContainer            *fyne.Container
	caneOfSomariaTapContainer      *fyne.Container
	caneOfByrnaTapContainer        *fyne.Container
	magicCapeTapContainer          *fyne.Container
	magicMirrorTapContainer        *fyne.Container
	pseudoPegasusBootsTapContainer *fyne.Container
	pegasusBootsTapContainer       *fyne.Container
	glovesTapContainer             *fyne.Container
	flippersTapContainer           *fyne.Container
	moonPearlTapContainer          *fyne.Container
	swordTapContainer              *fyne.Container
	shieldTapContainer             *fyne.Container
	mailTapContainer               *fyne.Container
	halfMagicTapContainer          *fyne.Container
	heartPieceTapContainer         *fyne.Container
	ganonGoalTapContainer          *fyne.Container
	pedestalGoalTapContainer       *fyne.Container
	triforceGoalTapContainer       *fyne.Container
	ganonTowerGoalTapContainer     *fyne.Container
	itemGrid                       *fyne.Container
}

func NewInventoryIcons(undoRedo *undo_redo.UndoRedoStacks, preferencesConfig *preferences.PreferencesFile, saveConfig *save.SaveFile) (*InventoryIcons, error) {
	var err error
	inventory := &InventoryIcons{
		preferencesFile: preferencesConfig,
		saveFile:        saveConfig,
		undoRedo:	undoRedo,
	}
	inventory.bowTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceBowGrayPng, resourceBowPng, resourceBowSilversPng}, 15, undoRedo, saveConfig, "Bow")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bowTapIcon: %w", err))
	}
	inventory.bowNonProgressiveTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceBowGrayPng, resourceBowPng, resourceSilversPng, resourceBowSilversPng}, 15, undoRedo, saveConfig, "Non-Progressive Bow")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bowNonProgressiveTapIcon: %w", err))
	}
	inventory.blueBoomerangTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceBlueBoomerangGrayPng, resourceBlueBoomerangPng}, 12, undoRedo, saveConfig, "Blue Boomerang")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making blueBoomerangTapIcon: %w", err))
	}
	inventory.redBoomerangTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceRedBoomerangGrayPng, resourceRedBoomerangPng}, 12, undoRedo, saveConfig, "Red Boomerang")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making redBoomerangTapIcon: %w", err))
	}
	inventory.hookshotTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceHookshotGrayPng, resourceHookshotPng}, 16, undoRedo, saveConfig, "Hookshot")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making hookshotTapIcon: %w", err))
	}
	inventory.bombTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceBombGrayPng, resourceBombPng}, 14, undoRedo, saveConfig, "Bomb")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bombTapIcon: %w", err))
	}
	inventory.mushroomTapIcon, err = tappable_icons.NewTappableIconWithIcon([]fyne.Resource{resourceMushroomGrayPng, resourceMushroomPng}, []fyne.Resource{resourceWitchPng, resourceWitchApprenticePng}, 16, undoRedo, saveConfig, "Mushroom")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making mushroomTapIcon: %w", err))
	}
	inventory.magicPowderTapIcon, err = tappable_icons.NewTappableIconWithIcon([]fyne.Resource{resourceMagicPowderGrayPng, resourceMagicPowderPng}, []fyne.Resource{resourceBatPng}, 16, undoRedo, saveConfig, "Magic Powder")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making magicPowderTapIcon: %w", err))
	}
	inventory.fireRodTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceFireRodGrayPng, resourceFireRodPng}, 16, undoRedo, saveConfig, "Fire Rod")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making fireRodTapIcon: %w", err))
	}
	inventory.iceRodTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceIceRodGrayPng, resourceIceRodPng}, 16, undoRedo, saveConfig, "Ice Rod")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making iceRodTapIcon: %w", err))
	}
	inventory.bombosTapIcon, err = tappable_icons.NewTappableIconWithString([]fyne.Resource{resourceBombosGrayPng, resourceBombosPng}, []string{"", "MM", "TR", "BOTH"}, 16, undoRedo, saveConfig, "Bombos")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bombosTapIcon: %w", err))
	}
	inventory.etherTapIcon, err = tappable_icons.NewTappableIconWithString([]fyne.Resource{resourceEtherGrayPng, resourceEtherPng}, []string{"", "MM", "TR", "BOTH"}, 16, undoRedo, saveConfig, "Ether")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making etherTapIcon: %w", err))
	}
	inventory.quakeTapIcon, err = tappable_icons.NewTappableIconWithString([]fyne.Resource{resourceQuakeGrayPng, resourceQuakePng}, []string{"", "MM", "TR", "BOTH"}, 16, undoRedo, saveConfig, "Quake")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making quakeTapIcon: %w", err))
	}
	inventory.lampTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceLampGrayPng, resourceLampPng}, 16, undoRedo, saveConfig, "Lamp")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making lampTapIcon: %w", err))
	}
	inventory.hammerTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceHammerGrayPng, resourceHammerPng}, 15, undoRedo, saveConfig, "Hammer")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making hammerTapIcon: %w", err))
	}
	inventory.fluteTapIcon, err = tappable_icons.NewTappableIconWithIcon([]fyne.Resource{resourceFluteGrayPng, resourceFlutePng}, []fyne.Resource{resourceBirdPng}, 14, undoRedo, saveConfig, "Flute")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making fluteTapIcon: %w", err))
	}
	inventory.shovelTapIcon, err = tappable_icons.NewTappableIconWithIcon([]fyne.Resource{resourceShovelGrayPng, resourceShovelPng}, []fyne.Resource{resourceFluteBoyPng}, 16, undoRedo, saveConfig, "Shovel")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making shovelTapIcon: %w", err))
	}
	inventory.bugNetTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceBugCatchingNetGrayPng, resourceBugCatchingNetPng}, 16, undoRedo, saveConfig, "Bug-Net")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bugNetTapIcon: %w", err))
	}
	inventory.bookOfMudoraTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceBookofMudoraGrayPng, resourceBookofMudoraPng}, 15, undoRedo, saveConfig, "Book of Mudora")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bookOfMudoraTapIcon: %w", err))
	}
	inventory.bottleTotalTapIcon, err = tappable_icons.NewTappableIconWithCenteredNum([]fyne.Resource{resourceEmptyBottleGrayPng, resourceEmptyBottlePng}, 4, 16, undoRedo, saveConfig, "Bottle Total")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bottleTapIcon: %w", err))
	}
	inventory.bottle1TapIcon, err = tappable_icons.NewTappableIconCycled([]fyne.Resource{resourceEmptyBottleGrayPng, resourceEmptyBottlePng, resourceBottlewithBeePng, resourceBottlewithFairyPng, resourceBottlewithGreenPotionPng, resourceBottlewithRedPotionPng, resourceBottlewithBluePotionPng}, 16, undoRedo, saveConfig, "Bottle 1")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bottle1TapIcon: %w", err))
	}
	inventory.bottle2TapIcon, err = tappable_icons.NewTappableIconCycled([]fyne.Resource{resourceEmptyBottleGrayPng, resourceEmptyBottlePng, resourceBottlewithBeePng, resourceBottlewithFairyPng, resourceBottlewithGreenPotionPng, resourceBottlewithRedPotionPng, resourceBottlewithBluePotionPng}, 16, undoRedo, saveConfig, "Bottle 2")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bottle2TapIcon: %w", err))
	}
	inventory.bottle3TapIcon, err = tappable_icons.NewTappableIconCycled([]fyne.Resource{resourceEmptyBottleGrayPng, resourceEmptyBottlePng, resourceBottlewithBeePng, resourceBottlewithFairyPng, resourceBottlewithGreenPotionPng, resourceBottlewithRedPotionPng, resourceBottlewithBluePotionPng}, 16, undoRedo, saveConfig, "Bottle 3")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bottle3TapIcon: %w", err))
	}
	inventory.bottle4TapIcon, err = tappable_icons.NewTappableIconCycled([]fyne.Resource{resourceEmptyBottleGrayPng, resourceEmptyBottlePng, resourceBottlewithBeePng, resourceBottlewithFairyPng, resourceBottlewithGreenPotionPng, resourceBottlewithRedPotionPng, resourceBottlewithBluePotionPng}, 16, undoRedo, saveConfig, "Bottle 4")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making bottle4TapIcon: %w", err))
	}
	inventory.caneOfSomariaTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceCaneofSomariaGrayPng, resourceCaneofSomariaPng}, 16, undoRedo, saveConfig, "Cane of Somaria")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making caneOfSomariaTapIcon: %w", err))
	}
	inventory.caneOfByrnaTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceCaneofByrnaGrayPng, resourceCaneofByrnaPng}, 16, undoRedo, saveConfig, "Cane of Byrna")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making caneOfByrnaTapIcon: %w", err))
	}
	inventory.magicCapeTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceMagicCapeGrayPng, resourceMagicCapePng}, 16, undoRedo, saveConfig, "Magic Cape")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making magicCapeTapIcon: %w", err))
	}
	inventory.magicMirrorTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceMagicMirrorGrayPng, resourceMagicMirrorPng}, 15, undoRedo, saveConfig, "Magic Mirror")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making magicMirrorTapIcon: %w", err))
	}
	inventory.pseudoPegasusBootsTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourcePseudoPegasusBootsPng, resourcePegasusBootsPng}, 15, undoRedo, saveConfig, "Pseudo Boots")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making pseudoPegasusBootsTapIcon: %w", err))
	}
	inventory.pegasusBootsTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourcePegasusBootsGrayPng, resourcePegasusBootsPng}, 15, undoRedo, saveConfig, "Pegasus Boots")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making pegasusBootsTapIcon: %w", err))
	}
	inventory.glovesTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourcePowerGlovesGrayPng, resourcePowerGlovesPng, resourceTitanMittsPng}, 16, undoRedo, saveConfig, "Gloves")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making glovesTapIcon: %w", err))
	}
	inventory.flippersTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceFlippersGrayPng, resourceFlippersPng}, 16, undoRedo, saveConfig, "Flippers")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making flippersTapIcon: %w", err))
	}
	inventory.moonPearlTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceMoonPearlGrayPng, resourceMoonPearlPng}, 12, undoRedo, saveConfig, "Moon Pearl")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making moonPearlTapIcon: %w", err))
	}
	inventory.swordTapIcon, err = tappable_icons.NewTappableIconVariedSize([]fyne.Resource{resourceFighterSSwordGrayPng, resourceFighterSSwordPng, resourceMasterSwordPng, resourceTemperedSwordPng, resourceGoldenSwordPng}, []float32{13, 13, 16, 16, 16}, undoRedo, saveConfig, "Sword")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making swordTapIcon: %w", err))
	}
	inventory.shieldTapIcon, err = tappable_icons.NewTappableIconVariedSize([]fyne.Resource{resourceFighterSShieldGrayPng, resourceFighterSShieldPng, resourceFireShieldPng, resourceMirrorShieldPng}, []float32{10, 10, 12, 16}, undoRedo, saveConfig, "Shield")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making shieldTapIcon: %w", err))
	}
	inventory.mailTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceGreenMailPng, resourceBlueMailPng, resourceRedMailPng}, 16, undoRedo, saveConfig, "Mail")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making mailTapIcon: %w", err))
	}
	inventory.halfMagicTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceHalfMagicGrayPng, resourceHalfMagicPng}, 16, undoRedo, saveConfig, "Half-Magic")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making halfMagicTapIcon: %w", err))
	}
	inventory.heartPieceTapIcon, err = tappable_icons.NewTappableIconCycled([]fyne.Resource{resourceHeartPieceGrayPng, resourceHeartPiece1Png, resourceHeartPiece2Png, resourceHeartPiece3Png}, 14, undoRedo, saveConfig, "Heart Pieces")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making heartPieceTapIcon: %w", err))
	}
	inventory.ganonGoalTapIcon, err = tappable_icons.NewTappableIconWithBottomCenteredString([]fyne.Resource{resourceGanonGrayPng, resourceGanonPng}, 16, undoRedo, saveConfig, "Ganon Goal")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making ganonGoalTapIcon: %w", err))
	}
	inventory.pedestalGoalTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourcePedestalGrayPng, resourcePedestalPng}, 16, undoRedo, saveConfig, "Pedestal Goal")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making pedestalGoalTapIcon: %w", err))
	}
	inventory.triforceGoalTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceTriforceGrayPng, resourceTriforcePng}, 16, undoRedo, saveConfig, "Triforce Goal")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making triforceGoalTapIcon: %w", err))
	}
	inventory.ganonTowerGoalTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceGanonTowerGrayPng, resourceGanonTowerPng}, 20, undoRedo, saveConfig, "Ganon Tower Goal")
	if err != nil {
		return nil, (fmt.Errorf("Encountered error making ganonTowerGoalTapIcon: %w", err))
	}

	return inventory, nil
}

func (i *InventoryIcons) Layout() *fyne.Container {
	i.bowTapContainer = i.bowTapIcon.Layout()
	i.bowNonProgressiveTapContainer = i.bowNonProgressiveTapIcon.Layout()
	i.blueBoomerangTapContainer = i.blueBoomerangTapIcon.Layout()
	i.redBoomerangTapContainer = i.redBoomerangTapIcon.Layout()
	i.hookshotTapContainer = i.hookshotTapIcon.Layout()
	i.bombTapContainer = i.bombTapIcon.Layout()
	i.mushroomTapContainer = i.mushroomTapIcon.Layout()
	i.magicPowderTapContainer = i.magicPowderTapIcon.Layout()
	i.fireRodTapContainer = i.fireRodTapIcon.Layout()
	i.iceRodTapContainer = i.iceRodTapIcon.Layout()
	i.bombosTapContainer = i.bombosTapIcon.Layout()
	i.etherTapContainer = i.etherTapIcon.Layout()
	i.quakeTapContainer = i.quakeTapIcon.Layout()
	i.lampTapContainer = i.lampTapIcon.Layout()
	i.hammerTapContainer = i.hammerTapIcon.Layout()
	i.fluteTapContainer = i.fluteTapIcon.Layout()
	i.shovelTapContainer = i.shovelTapIcon.Layout()
	i.bugNetTapContainer = i.bugNetTapIcon.Layout()
	i.bookOfMudoraTapContainer = i.bookOfMudoraTapIcon.Layout()
	i.bottleTotalTapContainer = i.bottleTotalTapIcon.Layout()
	i.bottle1TapContainer = i.bottle1TapIcon.Layout()
	i.bottle2TapContainer = i.bottle2TapIcon.Layout()
	i.bottle3TapContainer = i.bottle3TapIcon.Layout()
	i.bottle4TapContainer = i.bottle4TapIcon.Layout()
	i.caneOfSomariaTapContainer = i.caneOfSomariaTapIcon.Layout()
	i.caneOfByrnaTapContainer = i.caneOfByrnaTapIcon.Layout()
	i.magicCapeTapContainer = i.magicCapeTapIcon.Layout()
	i.magicMirrorTapContainer = i.magicMirrorTapIcon.Layout()
	i.pseudoPegasusBootsTapContainer = i.pseudoPegasusBootsTapIcon.Layout()
	i.pegasusBootsTapContainer = i.pegasusBootsTapIcon.Layout()
	i.glovesTapContainer = i.glovesTapIcon.Layout()
	i.flippersTapContainer = i.flippersTapIcon.Layout()
	i.moonPearlTapContainer = i.moonPearlTapIcon.Layout()
	i.swordTapContainer = i.swordTapIcon.Layout()
	i.shieldTapContainer = i.shieldTapIcon.Layout()
	i.mailTapContainer = i.mailTapIcon.Layout()
	i.halfMagicTapContainer = i.halfMagicTapIcon.Layout()
	i.heartPieceTapContainer = i.heartPieceTapIcon.Layout()
	i.ganonGoalTapContainer = i.ganonGoalTapIcon.Layout()
	i.pedestalGoalTapContainer = i.pedestalGoalTapIcon.Layout()
	i.triforceGoalTapContainer = i.triforceGoalTapIcon.Layout()
	i.ganonTowerGoalTapContainer = i.ganonTowerGoalTapIcon.Layout()

	boomContainer1 := container.NewWithoutLayout(i.blueBoomerangTapContainer, i.redBoomerangTapContainer)
	boomContainer2 := container.New(layout.NewCenterLayout(), boomContainer1)
	boomSize := i.blueBoomerangTapContainer.Size()
	blueBoomChangePosition := fyne.NewPos(-(boomSize.Width+theme.Padding()*2)/4, 0)
	redBoomChangePosition := fyne.NewPos((boomSize.Width+theme.Padding()*2)/4, 0)
	i.blueBoomerangTapContainer.Move(blueBoomChangePosition)
	i.redBoomerangTapContainer.Move(redBoomChangePosition)

	i.itemGrid = container.New(layout.NewGridLayout(4), i.bowTapContainer, i.bowNonProgressiveTapContainer, boomContainer2, i.hookshotTapContainer,
		i.bombTapContainer, i.mushroomTapContainer, i.magicPowderTapContainer, i.fireRodTapContainer, i.iceRodTapContainer, i.bombosTapContainer,
		i.etherTapContainer, i.quakeTapContainer, i.lampTapContainer, i.hammerTapContainer, i.fluteTapContainer, i.shovelTapContainer, i.bugNetTapContainer,
		i.bookOfMudoraTapContainer, i.bottleTotalTapContainer, i.bottle1TapContainer, i.bottle2TapContainer, i.bottle3TapContainer, i.bottle4TapContainer,
		i.caneOfSomariaTapContainer, i.caneOfByrnaTapContainer, i.magicCapeTapContainer, i.magicMirrorTapContainer, i.pseudoPegasusBootsTapContainer,
		i.pegasusBootsTapContainer, i.glovesTapContainer, i.flippersTapContainer, i.moonPearlTapContainer, i.swordTapContainer, i.shieldTapContainer,
		i.mailTapContainer, i.halfMagicTapContainer, i.heartPieceTapContainer, i.ganonGoalTapContainer, i.pedestalGoalTapContainer, i.triforceGoalTapContainer,
		i.ganonTowerGoalTapContainer)

	i.CreateSaveDefaults()
	i.SaveUpdate()
	i.PreferencesUpdate()

	return i.itemGrid
}

func (i *InventoryIcons) GetItemGrid() *fyne.Container {
	return i.itemGrid
}

func (i *InventoryIcons) Keys(result int) {
	if result == 0 {
		i.bowTapIcon.Keyed()
		i.bowNonProgressiveTapIcon.Keyed()
	} else if result == 1 {
		i.blueBoomerangTapIcon.Keyed()
	} else if result == 2 {
		i.redBoomerangTapIcon.Keyed()
	} else if result == 3 {
		i.hookshotTapIcon.Keyed()
	} else if result == 4 {
		i.bombTapIcon.Keyed()
	} else if result == 5 {
		i.mushroomTapIcon.Keyed()
	} else if result == 6 {
		i.magicPowderTapIcon.Keyed()
	} else if result == 7 {
		i.fireRodTapIcon.Keyed()
	} else if result == 8 {
		i.iceRodTapIcon.Keyed()
	} else if result == 9 {
		i.bombosTapIcon.Keyed()
	} else if result == 10 {
		i.etherTapIcon.Keyed()
	} else if result == 11 {
		i.quakeTapIcon.Keyed()
	} else if result == 12 {
		i.lampTapIcon.Keyed()
	} else if result == 13 {
		i.hammerTapIcon.Keyed()
	} else if result == 14 {
		i.fluteTapIcon.Keyed()
	} else if result == 15 {
		i.shovelTapIcon.Keyed()
	} else if result == 16 {
		i.bugNetTapIcon.Keyed()
	} else if result == 17 {
		i.bookOfMudoraTapIcon.Keyed()
	} else if result == 18 {
		i.bottleTotalTapIcon.Keyed()
	} else if result == 19 {
		i.bottle1TapIcon.Keyed()
	} else if result == 20 {
		i.bottle2TapIcon.Keyed()
	} else if result == 21 {
		i.bottle3TapIcon.Keyed()
	} else if result == 22 {
		i.bottle4TapIcon.Keyed()
	} else if result == 23 {
		i.caneOfSomariaTapIcon.Keyed()
	} else if result == 24 {
		i.caneOfByrnaTapIcon.Keyed()
	} else if result == 25 {
		i.magicCapeTapIcon.Keyed()
	} else if result == 26 {
		i.magicMirrorTapIcon.Keyed()
	} else if result == 27 {
		i.pegasusBootsTapIcon.Keyed()
		i.pseudoPegasusBootsTapIcon.Keyed()
	} else if result == 28 {
		i.glovesTapIcon.Keyed()
	} else if result == 29 {
		i.flippersTapIcon.Keyed()
	} else if result == 30 {
		i.moonPearlTapIcon.Keyed()
	} else if result == 31 {
		i.swordTapIcon.Keyed()
	} else if result == 32 {
		i.shieldTapIcon.Keyed()
	} else if result == 33 {
		i.mailTapIcon.Keyed()
	} else if result == 34 {
		i.halfMagicTapIcon.Keyed()
	} else if result == 35 {
		i.heartPieceTapIcon.Keyed()
	} else if result == 36 {
		i.ganonGoalTapIcon.Keyed()
	} else if result == 37 {
		i.pedestalGoalTapIcon.Keyed()
	} else if result == 38 {
		i.triforceGoalTapIcon.Keyed()
	} else if result == 39 {
		i.ganonTowerGoalTapIcon.Keyed()
	} else if result == 40 {
		i.undoRedo.Undo()
	} else if result == 41 {
		i.undoRedo.Redo()
	}
}

func (i *InventoryIcons) SaveUpdate() {
	i.bowTapIcon.Update()
	i.bowNonProgressiveTapIcon.Update()
	i.blueBoomerangTapIcon.Update()
	i.redBoomerangTapIcon.Update()
	i.hookshotTapIcon.Update()
	i.bombTapIcon.Update()
	i.mushroomTapIcon.Update()
	i.magicPowderTapIcon.Update()
	i.fireRodTapIcon.Update()
	i.iceRodTapIcon.Update()
	i.bombosTapIcon.Update()
	i.etherTapIcon.Update()
	i.quakeTapIcon.Update()
	i.lampTapIcon.Update()
	i.hammerTapIcon.Update()
	i.fluteTapIcon.Update()
	i.shovelTapIcon.Update()
	i.bugNetTapIcon.Update()
	i.bookOfMudoraTapIcon.Update()
	i.bottleTotalTapIcon.Update()
	i.bottle1TapIcon.Update()
	i.bottle2TapIcon.Update()
	i.bottle3TapIcon.Update()
	i.bottle4TapIcon.Update()
	i.caneOfSomariaTapIcon.Update()
	i.caneOfByrnaTapIcon.Update()
	i.magicCapeTapIcon.Update()
	i.magicMirrorTapIcon.Update()
	i.pseudoPegasusBootsTapIcon.Update()
	i.pegasusBootsTapIcon.Update()
	i.glovesTapIcon.Update()
	i.flippersTapIcon.Update()
	i.moonPearlTapIcon.Update()
	i.swordTapIcon.Update()
	i.shieldTapIcon.Update()
	i.mailTapIcon.Update()
	i.halfMagicTapIcon.Update()
	i.heartPieceTapIcon.Update()
	i.ganonGoalTapIcon.Update()
	i.pedestalGoalTapIcon.Update()
	i.triforceGoalTapIcon.Update()
	i.ganonTowerGoalTapIcon.Update()
}

func (i *InventoryIcons) UpdateGanonGoal(value int) {
	i.saveFile.SetSave("Ganon Goal_Current", value)
	i.ganonGoalTapIcon.Update()
}

func (i *InventoryIcons) PreferencesUpdate() {
	if i.preferencesFile.GetPreferenceBool("Bombs") {
		i.bombTapContainer.Show()
	} else {
		i.bombTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("Bottle_Full") {
		i.bottle1TapContainer.Show()
		i.bottle2TapContainer.Show()
		i.bottle3TapContainer.Show()
		i.bottle4TapContainer.Show()
		i.bottleTotalTapContainer.Hide()
	} else {
		i.bottleTotalTapContainer.Show()
		i.bottle1TapContainer.Hide()
		i.bottle2TapContainer.Hide()
		i.bottle3TapContainer.Hide()
		i.bottle4TapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("HalfMagic") {
		i.halfMagicTapContainer.Show()
	} else {
		i.halfMagicTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("Heart_Pieces") {
		i.heartPieceTapContainer.Show()
	} else {
		i.heartPieceTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("Mail") {
		i.mailTapContainer.Show()
	} else {
		i.mailTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("Shield") {
		i.shieldTapContainer.Show()
	} else {
		i.shieldTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("Sword") {
		i.swordTapContainer.Show()
	} else {
		i.swordTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("Pseudo_Boots") {
		i.pseudoPegasusBootsTapContainer.Show()
		i.pegasusBootsTapContainer.Hide()
	} else {
		i.pegasusBootsTapContainer.Show()
		i.pseudoPegasusBootsTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceBool("Progressive_Bows") {
		i.bowTapContainer.Show()
		i.bowNonProgressiveTapContainer.Hide()
	} else {
		i.bowNonProgressiveTapContainer.Show()
		i.bowTapContainer.Hide()
	}
	if i.preferencesFile.GetPreferenceInt("Goal") == 0 {
		i.ganonGoalTapContainer.Show()
		i.pedestalGoalTapContainer.Hide()
		i.triforceGoalTapContainer.Hide()
	} else if i.preferencesFile.GetPreferenceInt("Goal") == 1 {
		i.ganonGoalTapContainer.Hide()
		i.pedestalGoalTapContainer.Show()
		i.triforceGoalTapContainer.Hide()
	} else if i.preferencesFile.GetPreferenceInt("Goal") == 2 {
		i.ganonGoalTapContainer.Hide()
		i.pedestalGoalTapContainer.Hide()
		i.triforceGoalTapContainer.Show()
	} else {
		i.preferencesFile.SetPreference("Goal", 0)
		i.ganonGoalTapContainer.Show()
		i.pedestalGoalTapContainer.Hide()
		i.triforceGoalTapContainer.Hide()
	}

	i.itemGrid.Refresh()
}

func (i *InventoryIcons) CreateSaveDefaults() {
	i.bowTapIcon.SetSaveDefaults()
	i.bowNonProgressiveTapIcon.SetSaveDefaults()
	i.blueBoomerangTapIcon.SetSaveDefaults()
	i.redBoomerangTapIcon.SetSaveDefaults()
	i.hookshotTapIcon.SetSaveDefaults()
	i.bombTapIcon.SetSaveDefaults()
	i.mushroomTapIcon.SetSaveDefaults()
	i.magicPowderTapIcon.SetSaveDefaults()
	i.fireRodTapIcon.SetSaveDefaults()
	i.iceRodTapIcon.SetSaveDefaults()
	i.bombosTapIcon.SetSaveDefaults()
	i.etherTapIcon.SetSaveDefaults()
	i.quakeTapIcon.SetSaveDefaults()
	i.lampTapIcon.SetSaveDefaults()
	i.hammerTapIcon.SetSaveDefaults()
	i.fluteTapIcon.SetSaveDefaults()
	i.shovelTapIcon.SetSaveDefaults()
	i.bugNetTapIcon.SetSaveDefaults()
	i.bookOfMudoraTapIcon.SetSaveDefaults()
	i.bottleTotalTapIcon.SetSaveDefaults()
	i.bottle1TapIcon.SetSaveDefaults()
	i.bottle2TapIcon.SetSaveDefaults()
	i.bottle3TapIcon.SetSaveDefaults()
	i.bottle4TapIcon.SetSaveDefaults()
	i.caneOfSomariaTapIcon.SetSaveDefaults()
	i.caneOfByrnaTapIcon.SetSaveDefaults()
	i.magicCapeTapIcon.SetSaveDefaults()
	i.magicMirrorTapIcon.SetSaveDefaults()
	i.pseudoPegasusBootsTapIcon.SetSaveDefaults()
	i.pegasusBootsTapIcon.SetSaveDefaults()
	i.glovesTapIcon.SetSaveDefaults()
	i.flippersTapIcon.SetSaveDefaults()
	i.moonPearlTapIcon.SetSaveDefaults()
	i.swordTapIcon.SetSaveDefaults()
	i.shieldTapIcon.SetSaveDefaults()
	i.mailTapIcon.SetSaveDefaults()
	i.halfMagicTapIcon.SetSaveDefaults()
	i.heartPieceTapIcon.SetSaveDefaults()
	i.ganonGoalTapIcon.SetSaveDefaults()
	i.pedestalGoalTapIcon.SetSaveDefaults()
	i.triforceGoalTapIcon.SetSaveDefaults()
	i.ganonTowerGoalTapIcon.SetSaveDefaults()
}

func (i *InventoryIcons) RestoreDefaults() {
	i.bowTapIcon.GetSaveDefaults()
	i.bowNonProgressiveTapIcon.GetSaveDefaults()
	i.blueBoomerangTapIcon.GetSaveDefaults()
	i.redBoomerangTapIcon.GetSaveDefaults()
	i.hookshotTapIcon.GetSaveDefaults()
	i.bombTapIcon.GetSaveDefaults()
	i.mushroomTapIcon.GetSaveDefaults()
	i.magicPowderTapIcon.GetSaveDefaults()
	i.fireRodTapIcon.GetSaveDefaults()
	i.iceRodTapIcon.GetSaveDefaults()
	i.bombosTapIcon.GetSaveDefaults()
	i.etherTapIcon.GetSaveDefaults()
	i.quakeTapIcon.GetSaveDefaults()
	i.lampTapIcon.GetSaveDefaults()
	i.hammerTapIcon.GetSaveDefaults()
	i.fluteTapIcon.GetSaveDefaults()
	i.shovelTapIcon.GetSaveDefaults()
	i.bugNetTapIcon.GetSaveDefaults()
	i.bookOfMudoraTapIcon.GetSaveDefaults()
	i.bottleTotalTapIcon.GetSaveDefaults()
	i.bottle1TapIcon.GetSaveDefaults()
	i.bottle2TapIcon.GetSaveDefaults()
	i.bottle3TapIcon.GetSaveDefaults()
	i.bottle4TapIcon.GetSaveDefaults()
	i.caneOfSomariaTapIcon.GetSaveDefaults()
	i.caneOfByrnaTapIcon.GetSaveDefaults()
	i.magicCapeTapIcon.GetSaveDefaults()
	i.magicMirrorTapIcon.GetSaveDefaults()
	i.pseudoPegasusBootsTapIcon.GetSaveDefaults()
	i.pegasusBootsTapIcon.GetSaveDefaults()
	i.glovesTapIcon.GetSaveDefaults()
	i.flippersTapIcon.GetSaveDefaults()
	i.moonPearlTapIcon.GetSaveDefaults()
	i.swordTapIcon.GetSaveDefaults()
	i.shieldTapIcon.GetSaveDefaults()
	i.mailTapIcon.GetSaveDefaults()
	i.halfMagicTapIcon.GetSaveDefaults()
	i.heartPieceTapIcon.GetSaveDefaults()
	i.ganonGoalTapIcon.GetSaveDefaults()
	i.pedestalGoalTapIcon.GetSaveDefaults()
	i.triforceGoalTapIcon.GetSaveDefaults()
	i.ganonTowerGoalTapIcon.GetSaveDefaults()
}
