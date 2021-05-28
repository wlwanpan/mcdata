package mcdata

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"
)

const (
	EditionPC = "PC"
	EditionPE = "PE"

	SubmoduleDataPath = "minecraft-data/data"
)

var (
	//go:embed minecraft-data/data/*
	embededMCData embed.FS
)

type dataPath map[string]map[string]string

type dataPaths struct {
	PC dataPath `json:"pc"`
	PE dataPath `json:"pe"`
}

func (dp *dataPaths) getEditionedPath(e string) dataPath {
	if e == EditionPC {
		return dp.PC
	}
	return dp.PE
}

func (dp *dataPaths) getVersionedPaths(e, v string) (map[string]string, bool) {
	path := dp.getEditionedPath(e)
	vp, exist := path[v]
	return vp, exist
}

func (dp *dataPaths) getSupportedEditions(e string) []string {
	path := dp.getEditionedPath(e)
	editions := []string{}
	for s := range path {
		editions = append(editions, s)
	}
	return editions
}

func LoadDataPaths() (*dataPaths, error) {
	dataPathJsonPath := filepath.Join(SubmoduleDataPath, "dataPaths.json")
	dataPathsFile, err := embededMCData.Open(dataPathJsonPath)
	if err != nil {
		return nil, err
	}
	paths := &dataPaths{}
	jsonParser := json.NewDecoder(dataPathsFile)
	err = jsonParser.Decode(paths)
	return paths, err
}

func LoadDataToStruct(edition, version, resource string, d interface{}) error {
	resourcePath := filepath.Join(SubmoduleDataPath, edition, version, fmt.Sprintf("%s.json", resource))
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
