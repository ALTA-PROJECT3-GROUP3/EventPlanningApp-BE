package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	JWTKEY                 string
	CloudinaryName         string
	CloudinaryApiKey       string
	CloudinaryApiScret     string
	CloudinaryUploadFolder string
	MidtransServerKey      string
	MidtransClientKey      string
)

type AppConfig struct {
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBPORT     int
	DBNAME     string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DB_USER"); found {
		app.DBUSER = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PASS"); found {
		app.DBPASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_HOST"); found {
		app.DBHOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DB_NAME"); found {
		app.DBNAME = val
		isRead = false
	}

	if val, found := os.LookupEnv("JWT_KEY"); found {
		JWTKEY = val
		isRead = false
	}

	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		CloudinaryName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		CloudinaryApiKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		CloudinaryApiScret = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		CloudinaryUploadFolder = val
		isRead = false
	}
	if val, found := os.LookupEnv("MIDTRANS_SERVERKEY"); found {
		MidtransServerKey = val
		isRead = false
	}

	if val, found := os.LookupEnv("MIDTRANS_CLIENTKEY"); found {
		MidtransClientKey = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.DBUSER = viper.Get("DB_USER").(string)
		app.DBPASSWORD = viper.Get("DB_PASS").(string)
		app.DBHOST = viper.Get("DB_HOST").(string)
		app.DBPORT, _ = strconv.Atoi(viper.Get("DB_PORT").(string))
		app.DBNAME = viper.Get("DB_NAME").(string)

		JWTKEY = viper.Get("JWT_KEY").(string)

		CloudinaryName = viper.Get("CLOUDINARY_CLOUD_NAME").(string)
		CloudinaryApiKey = viper.Get("CLOUDINARY_API_KEY").(string)
		CloudinaryApiScret = viper.Get("CLOUDINARY_API_SECRET").(string)
		CloudinaryUploadFolder = viper.Get("CLOUDINARY_UPLOAD_FOLDER").(string)

		MidtransServerKey = viper.Get("MIDTRANS_SERVERKEY").(string)
		MidtransClientKey = viper.Get("MIDTRANS_CLIENTKEY").(string)

	}
	return &app
}
