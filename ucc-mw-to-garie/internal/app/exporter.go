package app

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Sitelist []struct {
	WebSiteName string `json:"web-site-name"`
	Config      string `json:"config"`
}

type GarieConfig struct {
	Cron string `json:"cron"`
	Urls []GarieUrl `json:"urls"`
}
type GarieUrl struct {
	Url string `json:"url"`
	Plugins []GariePlugin `json:"plugins"`
}
type GariePlugin struct {
	Name string `json:"name"`
	Report bool `json:"report"`
}

var (
	s Sitelist
)
func StartProcessing(env *Config) (s Sitelist, checkId string){
	client := NewHTTPClient(env)
	checkId = strings.ReplaceAll(client.sendStartJobSignal(), "\"", "")
	s = client.getSitelist()
	return s, checkId
}

func StopProcessing(env *Config, checkId string) {
	client := NewHTTPClient(env)
	//client.sendResult(checkId, "{\"checker-title\": \"lighthouse\", \"message\": \"Lighthouse config updated\"}")
	client.sendEndJobSignal(checkId)
}

func newDefaultGarieConfig (cron string) (GarieConfig, []GariePlugin) {
	return GarieConfig{
		Cron: cron,
	},
	[]GariePlugin{
		{
			Name: "lighthouse",
			Report: true,
		},
	}
}

func GenerateGarieConfig(cron, garieConfigPath string, sitelist Sitelist ) {
	var gu []GarieUrl
	gcd, gpd := newDefaultGarieConfig(cron)
	for _, s := range sitelist {
		var u GarieUrl
		u.Url = s.WebSiteName
		u.Plugins = gpd
		gu = append(gu,u)
	}
	gcd.Urls = gu
	file, _ := json.MarshalIndent(gcd, "", " ")
	err := ioutil.WriteFile(garieConfigPath, file, 0644)
	HandleError(err)
}
