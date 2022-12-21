package preferences

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type PreferencesFile struct {
	config *viper.Viper
}

func NewPreferencesFile() *PreferencesFile {
	preferences := &PreferencesFile{
		config: loadPreferences(),
	}

	preferences.CreateDefaults()

	preferences.config.BindEnv("fyne_scale")

	fyneScale := preferences.config.GetFloat64("fyne_scale")

	os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))

	return preferences
}

func loadPreferences() *viper.Viper {
	config := viper.New()
	config.SetConfigName("preferences") // name of config file (without extension)
	config.SetConfigType("toml")        // REQUIRED if the config file does not have the extension in the name
	//config.AddConfigPath("")   // path to look for the config file in
	config.AddConfigPath(".")    // optionally look for config in the working directory
	err := config.ReadInConfig() // Find and read the config file

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			config.SafeWriteConfigAs("./preferences.toml")
		} else {
			panic(fmt.Errorf("fatal error save file: %w", err))
		}
	}

	return config
}

func (p *PreferencesFile) SavePreferences() {
	p.config.WriteConfig()
}

func (p *PreferencesFile) GetPreferenceInt(key string) int {
	return p.config.GetInt(key)
}

func (p *PreferencesFile) GetPreferenceFloat(key string) float64 {
	return p.config.GetFloat64(key)
}

func (p *PreferencesFile) GetPreferenceBool(key string) bool {
	return p.config.GetBool(key)
}

func (p *PreferencesFile) SetPreference(key string, value interface{}) {
	p.config.Set(key, value)
}

func (p *PreferencesFile) CreateDefaults() {
	p.config.SetDefault("Big_Keys", false)
	p.config.SetDefault("Big_Keys_Required", false)
	p.config.SetDefault("Bombs", true)
	p.config.SetDefault("Bosses", false)
	p.config.SetDefault("Bosses_Required", false)
	p.config.SetDefault("Bottle_Full", false)
	p.config.SetDefault("Chest_Count", false)
	p.config.SetDefault("Compasses", false)
	p.config.SetDefault("Fullscreen", false)
	p.config.SetDefault("Fyne_Scale", 1.000000)
	p.config.SetDefault("Global_Hotkeys", true)
	p.config.SetDefault("Goal", 0)
	p.config.SetDefault("Halfmagic", true)
	p.config.SetDefault("Heart_Pieces", true)
	p.config.SetDefault("Keys", false)
	p.config.SetDefault("Keys_Required", false)
	p.config.SetDefault("Mail", true)
	p.config.SetDefault("Maps", false)
	p.config.SetDefault("Progressive_Bows", true)
	p.config.SetDefault("Pseudo_Boots", false)
	p.config.SetDefault("Shield", true)
	p.config.SetDefault("Sword", true)
}

func (p *PreferencesFile) RestoreDefaults() {
	p.config.Set("Big_Keys", false)
	p.config.Set("Big_Keys_Required", false)
	p.config.Set("Bombs", true)
	p.config.Set("Bosses", false)
	p.config.Set("Bosses_Required", false)
	p.config.Set("Bottle_Full", false)
	p.config.Set("Chest_Count", false)
	p.config.Set("Compasses", false)
	p.config.Set("Fullscreen", false)
	p.config.Set("Fyne_Scale", 1.000000)
	p.config.Set("Global_Hotkeys", true)
	p.config.Set("Goal", 0)
	p.config.Set("Halfmagic", true)
	p.config.Set("Heart_Pieces", true)
	p.config.Set("Keys", false)
	p.config.Set("Keys_Required", false)
	p.config.Set("Mail", true)
	p.config.Set("Maps", false)
	p.config.Set("Progressive_Bows", true)
	p.config.Set("Pseudo_Boots", false)
	p.config.Set("Shield", true)
	p.config.Set("Sword", true)
}
