package env

import "github.com/spf13/viper"

type (
 	Conf struct {
 		MongoHost string
 		MongoPort int
	}
)

const (
	MONGO_HOST string = "MONGO_HOST"
	MONGO_PORT string = "MONGO_STRING"
)

func Parse() *Conf {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")
	viper.SetDefault(MONGO_HOST, "localhost")
	viper.SetDefault(MONGO_PORT, 27017)

	conf := &Conf{}

	conf.MongoHost = viper.GetString(MONGO_HOST)
	conf.MongoPort = viper.GetInt(MONGO_PORT)

	return conf
}
