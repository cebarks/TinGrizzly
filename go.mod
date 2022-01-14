module github.com/cebarks/TinGrizzly

go 1.17

require (
	github.com/cebarks/spriteplus v0.5.2
	github.com/dusk125/pixelutils v1.1.0 // v1.1.0 breaks
	github.com/faiface/pixel v0.10.0
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20211213063430-748e38ca8aec
	github.com/kelindar/tile v1.2.0
	github.com/lrita/cmap v0.0.0-20200818170753-e987cd3dfa73
	github.com/panjf2000/ants/v2 v2.4.7
	github.com/pelletier/go-toml v1.9.4
	github.com/rs/zerolog v1.26.1
	github.com/snwfdhmp/errlog v0.0.0-20201130182740-aef7af651c46
	github.com/stretchr/testify v1.7.0
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/faiface/glhf v0.0.0-20211013000516-57b20770c369 // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/mathgl v1.0.0 // indirect
	github.com/kelindar/iostream v1.3.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pzsz/voronoi v0.0.0-20130609164533-4314be88c79f // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

// replace github.com/cebarks/spriteplus => ../spriteplus

replace github.com/dusk125/pixelutils v1.1.0 => github.com/dusk125/pixelutils v1.0.0
