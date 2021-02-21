# mcdata

mcdata is a code generator tool for generating a package of go structs of minecraft data entities from [minecraft-data](https://github.com/PrismarineJS/minecraft-data).


## Usage

```bash
go run cmd/main.go gen --edition PC --version 1.8 --output out/entities
```

Note: the generated package name will be the output dir name. For example, running the above command will generate a package `"entities"` with the following files:

```
- src
    /out
        /entities
            /materials.go
            /items.go
            /recipes.go
            /version.go
            /foods.go
            /particles.go
            /instruments.go
            /protocol.go
            /biomes.go
            /enchantments.go
            /entities.go
            /blocks.go
            /blockCollisionShapes.go
            /language.go
            /mapIcons.go
            /effects.go
            /windows.go
```

### materials.go

```go
package entities

type Materials struct {
	Dirt struct {
		Two56 int64 `json:"256"`
		Two69 int64 `json:"269"`
		Two73 int64 `json:"273"`
		Two77 int64 `json:"277"`
		Two84 int64 `json:"284"`
	} `json:"dirt"`
	Leaves struct {
		Two67   float64 `json:"267"`
		Two68   float64 `json:"268"`
		Two72   float64 `json:"272"`
		Two76   float64 `json:"276"`
		Two83   float64 `json:"283"`
		Three59 int64   `json:"359"`
	} `json:"leaves"`
	Melon struct {
		Two67 float64 `json:"267"`
		Two68 float64 `json:"268"`
		Two72 float64 `json:"272"`
		Two76 float64 `json:"276"`
		Two83 float64 `json:"283"`
	} `json:"melon"`
	Plant struct {
		Two58 int64   `json:"258"`
		Two67 float64 `json:"267"`
		Two68 float64 `json:"268"`
		Two71 int64   `json:"271"`
		Two72 float64 `json:"272"`
		Two75 int64   `json:"275"`
		Two76 float64 `json:"276"`
		Two79 int64   `json:"279"`
		Two83 float64 `json:"283"`
		Two86 int64   `json:"286"`
	} `json:"plant"`
	Rock struct {
		Two57 int64 `json:"257"`
		Two70 int64 `json:"270"`
		Two74 int64 `json:"274"`
		Two78 int64 `json:"278"`
		Two85 int64 `json:"285"`
	} `json:"rock"`
	Web struct {
		Two67   int64 `json:"267"`
		Two68   int64 `json:"268"`
		Two72   int64 `json:"272"`
		Two76   int64 `json:"276"`
		Two83   int64 `json:"283"`
		Three59 int64 `json:"359"`
	} `json:"web"`
	Wood struct {
		Two58 int64 `json:"258"`
		Two71 int64 `json:"271"`
		Two75 int64 `json:"275"`
		Two79 int64 `json:"279"`
		Two86 int64 `json:"286"`
	} `json:"wood"`
	Wool struct {
		Three59 float64 `json:"359"`
	} `json:"wool"`
}
```
