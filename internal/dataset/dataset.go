package dataset

import "embed"

//go:embed data/data/pokemon-index.json
var PokemonIndexJSON []byte

//go:embed data/data/pokemon/*
var PokemonDataFS embed.FS
