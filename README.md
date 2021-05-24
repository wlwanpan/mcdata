# mcdata

mcdata is a code generator tool for generating a package of go structs of minecraft data entities from [minecraft-data](https://github.com/PrismarineJS/minecraft-data).

Note: go 1.16 is required since this package relies on the [go embeded](https://golang.org/pkg/embed/) package. 


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

### effects.go

```go
package entities

type Effects []struct {
	DisplayName string `json:"displayName"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}
```

### Load your pre-generated go struct
```go
effects := &entities.Effects{}
err := mcdata.LoadDataToStruct("pc", "1.8", "effects", effects)
if err != nil {
	...
}

...
```
