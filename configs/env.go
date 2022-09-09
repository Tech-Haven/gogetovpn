package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		fmt.Println("Env loaded...")
	}

	if AuthSecret() == "" {
		log.Fatal("Error: AUTH_SECRET not set")
	}

	if OpenVpnInstallScriptDir() == "" {
		log.Fatal("Error: OPENVPN_INSTALL_SCRIPT_DIR not set")
	}
}

func AuthSecret() string {
	return os.Getenv("AUTH_SECRET")
}

func OpenVpnInstallScriptDir() string {
	return os.Getenv("OPENVPN_INSTALL_SCRIPT_DIR")
}
