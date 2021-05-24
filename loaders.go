package mcdata

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"
)

var (
	//go:embed minecraft-data/data/*
	embededMCData embed.FS
)

func LoadDataToStruct(edition, version, resource string, d interface{}) error {
	resourcePath := filepath.Join("minecraft-data/data", edition, version, fmt.Sprintf("%s.json", resource))
	resourceFile, err := embededMCData.Open(resourcePath)
	if err != nil {
		return err
	}
	defer resourceFile.Close()

	resourceStat, err := resourceFile.Stat()
	if err != nil {
		return err
	}
	data := make([]byte, resourceStat.Size())

	_, err = resourceFile.Read(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, d)
}
