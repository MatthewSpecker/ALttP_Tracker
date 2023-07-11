package preferences

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"

	"github.com/spf13/viper"
)

type PreferencesFile struct {
	config *viper.Viper
}

func NewPreferencesFile(preferencesFileDirectory string, mainWindow fyne.Window) *PreferencesFile {
	preferences := &PreferencesFile{
		config: loadPreferences(preferencesFileDirectory),
	}

	preferences.CreateDefaults(mainWindow)

	preferences.config.BindEnv("fyne_scale")

	fyneScale := preferences.config.GetFloat64("fyne_scale")

	os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))

	return preferences
}

func loadPreferences(preferencesFileDirectory string) *viper.Viper {
	config := viper.New()
	config.SetConfigName("preferences")
	config.SetConfigType("toml")
	config.AddConfigPath(preferencesFileDirectory)
	err := config.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			config.SafeWriteConfigAs(filepath.Join(preferencesFileDirectory, "preferences.toml"))
		} else {
			panic(fmt.Errorf("fatal error preferences file: %w", err))
		}
	}

	return config
}

func (p *PreferencesFile) SavePreferences() error {
	return p.config.WriteConfig()
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

func (p *PreferencesFile) GetWindowSize() fyne.Size {
	width := float32(p.config.GetFloat64("MainWindowSizeWidth"))
	height := float32(p.config.GetFloat64("MainWindowSizeHeight"))

	mainWindowSize := fyne.NewSize(width, height)

	return mainWindowSize
}

func (p *PreferencesFile) SetPreference(key string, value interface{}) error {
	switch value.(type) {
	case bool, int, float64:
		p.config.Set(key, value)
		return nil
	default:
		return fmt.Errorf("%T is not an acceptable type to SetPreference. Must be string, bool/int/float64", value)
	}
}

func (p *PreferencesFile) SetWindowSize(mainWindow fyne.Window) {
	p.config.Set("MainWindowSizeWidth", p.GetWindowSizeWidth(mainWindow))
	p.config.Set("MainWindowSizeHeight", p.GetWindowSizeHeight(mainWindow))
}

func (p *PreferencesFile) GetWindowSizeWidth(mainWindow fyne.Window) float64 {
	windowSize := mainWindow.Content().Size()
	width, _ := windowSize.Components()
	return float64(width)
}

func (p *PreferencesFile) GetWindowSizeHeight(mainWindow fyne.Window) float64 {
	windowSize := mainWindow.Content().Size()
	_, height := windowSize.Components()
	return float64(height)
}

func (p *PreferencesFile) CreateDefaults(mainWindow fyne.Window) {
	p.config.SetDefault("Big_Keys", false)
	p.config.SetDefault("Bombs", true)
	p.config.SetDefault("Bosses", false)
	p.config.SetDefault("Bottle_Full", true)
	p.config.SetDefault("Chest_Count", true)
	p.config.SetDefault("Compasses", false)
	p.config.SetDefault("Fullscreen", false)
	p.config.SetDefault("Fyne_Scale", 1.000000)
	p.config.SetDefault("Global_Hotkeys", true)
	p.config.SetDefault("Halfmagic", true)
	p.config.SetDefault("Heart_Pieces", true)
	p.config.SetDefault("Keys", false)
	p.config.SetDefault("Mail", true)
	p.config.SetDefault("MainWindowSizeWidth", p.GetWindowSizeWidth(mainWindow))
	p.config.SetDefault("MainWindowSizeHeight", p.GetWindowSizeHeight(mainWindow))
	p.config.SetDefault("Maps", false)
	p.config.SetDefault("Shield", true)
	p.config.SetDefault("Sword", true)
}

func (p *PreferencesFile) RestoreDefaults(mainWindow fyne.Window) {
	p.config.Set("Big_Keys", false)
	p.config.Set("Bombs", true)
	p.config.Set("Bosses", false)
	p.config.Set("Bottle_Full", true)
	p.config.Set("Chest_Count", true)
	p.config.Set("Compasses", false)
	p.config.Set("Fullscreen", false)
	p.config.Set("Fyne_Scale", 1.000000)
	p.config.Set("Global_Hotkeys", true)
	p.config.Set("Halfmagic", true)
	p.config.Set("Heart_Pieces", true)
	p.config.Set("Keys", false)
	p.config.Set("Mail", true)
	p.config.Set("MainWindowSizeWidth", p.GetWindowSizeWidth(mainWindow))
	p.config.Set("MainWindowSizeHeight", p.GetWindowSizeHeight(mainWindow))
	p.config.Set("Maps", false)
	p.config.Set("Shield", true)
	p.config.Set("Sword", true)
}
