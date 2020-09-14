package util

import (
	"log"
)

//const DbConfig string = "kushan:asdf@1234@tcp(34.211.94.118:3306)/otpUsers?charset=utf8&parseTime=True&loc=Local"
var DBUSER, DBPW, DBHOST, DBPORT, DBNAME, HSM_HOST, HSM_PORT, HARSHI_VAULT_ROOT_TOKEN, PUSH_URL, HSM_URL_PREFIX string

var DbConfig, HashiCorpUrl, HarshiVaultCaClientUrl, HashiVaultRootToken, ZsPushServiceUrl string

func init() {

	DbConfig = "userName:password@tcp(DBHOST:DBPORT)/DBNAME?charset=utf8&parseTime=True&loc=Local"
	HashiCorpUrl = ""
	HarshiVaultCaClientUrl = ""
	HashiVaultRootToken = ""
	ZsPushServiceUrl = ""

	log.Println("DbConfig : ", DbConfig)
	log.Println("HashiCorpUrl : ", HashiCorpUrl)
	log.Println("HarshiVaultCaClientUrl : ", HarshiVaultCaClientUrl)
	log.Println("HashiVaultRootToken : ", HashiVaultRootToken)
	log.Println("ZsPushServiceUrl : ", ZsPushServiceUrl)
}
