package env

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	Conf struct {
		MongoHost string
		MongoPort int
		MongoUid  string
		MongoPwd  string
	}
)

const (
	MONGO_UID  string = "MONGO_UID"
	MONGO_PWD  string = "MONGO_PWD"
	MONGO_HOST string = "MONGO_HOST"
	MONGO_PORT string = "MONGO_STRING"
)

func Parse() *Conf {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")
	viper.SetDefault(MONGO_HOST, "localhost")
	viper.SetDefault(MONGO_PORT, 27017)
	viper.SetDefault(MONGO_UID, "")
	viper.SetDefault(MONGO_PWD, "")

	conf := &Conf{}

	conf.MongoHost = viper.GetString(MONGO_HOST)
	conf.MongoPort = viper.GetInt(MONGO_PORT)
	conf.MongoUid = viper.GetString(MONGO_UID)
	conf.MongoPwd = viper.GetString(MONGO_PWD)

	return conf
}

func (conf *Conf) MongoConnStr() string {

	if conf.MongoUid == "" {
		return fmt.Sprintf("mongodb://%v:%v/",
			conf.MongoHost,
			conf.MongoPort)
	}

	return fmt.Sprintf("mongodb://%v:%v@%v:%v/",
		conf.MongoUid,
		conf.MongoPwd,
		conf.MongoHost,
		conf.MongoPort)
}
