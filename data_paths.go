package mcdata

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
