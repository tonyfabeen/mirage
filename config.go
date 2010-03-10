package config

import(
	"io/ioutil"
	"fmt"
	"json"
)

const (
	CONFIG_FILE_NAME = "config"
)


type Config struct {
	Port string
	Channels []string
}

func (config *Config) SetChannels(channels []string){
	config.Channels = channels
}


func (config *Config) Parse() bool{
	bytes, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err == nil {
		ok, errToken := json.Unmarshal(string(bytes), config)
		if !ok {
			fmt.Printf("Error on Parse Config File : %s\n", errToken)
		}else {
			return true
		}
	}
	return false
}
