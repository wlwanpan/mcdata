package mcdata

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ChimeraCoder/gojson"
	"github.com/hashicorp/go-multierror"
)

var (
	ErrVersionNotSupported = errors.New("version provided not supported")
)

const (
	EditionPC = "PC"
	EditionPE = "PE"
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

func loadDataPaths() (*dataPaths, error) {
	curDir, _ := os.Getwd()
	pathsFile, err := os.Open(filepath.Join(curDir, "../minecraft-data/data/dataPaths.json"))
	if err != nil {
		return nil, err
	}

	paths := &dataPaths{}
	jsonParser := json.NewDecoder(pathsFile)
	err = jsonParser.Decode(paths)
	return paths, err
}

func GenerateGoStructs(edition, version, dest string) error {
	dataPaths, err := loadDataPaths()
	if err != nil {
		return err
	}
	versionedPath, exist := dataPaths.getVersionedPaths(edition, version)
	if !exist {
		return ErrVersionNotSupported
	}

	curDir, _ := os.Getwd()
	basepath := filepath.Join(curDir, "../minecraft-data/data")
	outpath := filepath.Join(curDir, dest)
	if err := os.Mkdir(outpath, 0777); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	var res error
	for datatype, datapath := range versionedPath {
		datafilepath := filepath.Join(basepath, datapath, datatype+".json")

		datafile, err := os.Open(datafilepath)
		if err != nil {
			// "Should" never happen, since dataPath definitions will always reflect a json file
			// within its respective version. Else an error should be filed to Prismarine:
			// https://github.com/PrismarineJS/minecraft-data/issues
			multierror.Append(res, err)
			continue
		}

		structName := strings.Title(datatype)
		genfile, err := gojson.Generate(datafile, gojson.ParseJson, structName, dest, []string{"json"}, false, true)
		if err != nil {
			multierror.Append(res, err)
			continue
		}

		structFile := filepath.Join(outpath, datatype+".go")
		log.Printf("Generating file %s", structFile)
		f, err := os.Create(structFile)
		if err != nil {
			multierror.Append(res, err)
			continue
		}
		defer f.Close()

		_, err = f.Write(genfile)
		if err != nil {
			multierror.Append(res, err)
		}
	}

	return res
}
