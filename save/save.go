package save

/*To Do:
-Add descriptions to functions
*/

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type SaveFile struct {
	config *viper.Viper
}

func NewSaveFile(saveFileDirectory string) *SaveFile {
	save := &SaveFile{
		config: loadState(saveFileDirectory),
	}

	return save
}

func loadState(saveFileDirectory string) *viper.Viper {
	config := viper.New()
	config.SetConfigName("save")
	config.SetConfigType("toml")
	config.AddConfigPath(saveFileDirectory)
	err := config.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			config.SafeWriteConfigAs(filepath.Join(saveFileDirectory, "save.toml"))
		} else {
			panic(fmt.Errorf("fatal error save file: %w", err))
		}
	}

	return config
}

func (s *SaveFile) SaveState() error {
	return s.config.WriteConfig()
}

func (s *SaveFile) GetSaveInt(key string) int {
	key = key + "_Int"
	return s.config.GetInt(key)
}

func (s *SaveFile) GetSaveBool(key string) bool {
	key = key + "_Bool"
	return s.config.GetBool(key)
}

func (s *SaveFile) SetSave(key string, value interface{}) {
	switch value.(type) {
	case bool:
		key = key + "_Bool"
	case int:
		key = key + "_Int"
	}

	s.config.Set(key, value)
}

func (s *SaveFile) SetDefault(key string, value interface{}) {
	switch value.(type) {
	case bool:
		s.config.SetDefault(key+"_Bool", value)
	case int:
		s.config.SetDefault(key+"_Int", value)
	}
}
