package app

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	UccBaseURL string `yaml:"ucc_url"`
	UccToken string `yaml:"ucc_token"`
	UccSitelistEndpoint string `yaml:"ucc_sitelist_endpoint"`
	UccReportEndpoint string `yaml:"ucc_report_endpoint"`
	UccStartJobEndpoint string `yaml:"ucc_start_job_endpoint"`
	UccEndJobEndpoint string `yaml:"ucc_end_job_endpoint"`
	GarieConfigPath string `yaml:"garie_config_path"`
	CronString string `yaml:"cron_string"`
}

func newConfig() *Config {
	return &Config{
		UccSitelistEndpoint: "/api/settings/lighthouse",
		UccReportEndpoint: "/api/checker/receiveresult",
		UccStartJobEndpoint: "/api/checker/started/lighthouse",
		UccEndJobEndpoint: "/api/checker/finished",
		GarieConfigPath: "./config/garie-config.json",
		CronString: "0 1 * * *",
	}
}

func Configure() (*Config, error) {
	c := newConfig()
	yamlFile, err := ioutil.ReadFile("config/exporter.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}