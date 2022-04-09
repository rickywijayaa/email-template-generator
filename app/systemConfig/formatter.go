package systemConfig

import "email-template-generator/entity"

type SystemConfigFormatter struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Value string `json:"value"`
}

func SystemConfigFormat(systemConfig entity.SystemConfig) SystemConfigFormatter {
	formatter := SystemConfigFormatter{
		ID:    systemConfig.ID,
		Code:  systemConfig.Code,
		Value: systemConfig.Value,
	}

	return formatter
}

func SystemConfigsFormat(systemConfig []entity.SystemConfig) []SystemConfigFormatter {
	formatter := []SystemConfigFormatter{}

	for _, config := range systemConfig {
		systemConfigFormat := SystemConfigFormat(config)
		formatter = append(formatter, systemConfigFormat)
	}

	return formatter
}
