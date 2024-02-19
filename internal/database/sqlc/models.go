// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
	pokemon "github.com/stilesdev/easy-dex/internal/pokemon"
)

type Pokemon struct {
	ID                   string
	Nid                  string
	DexNum               int32
	FormID               pgtype.Text
	Name                 string
	FormName             pgtype.Text
	Region               string
	Generation           int32
	Type1                string
	Type2                pgtype.Text
	TeraType             pgtype.Text
	Color                string
	PrimaryAbility       string
	SecondaryAbility     pgtype.Text
	HiddenAbility        pgtype.Text
	IsDefault            bool
	IsForm               bool
	IsLegendary          bool
	IsMythical           bool
	IsUltraBeast         bool
	UltraBeastCode       pgtype.Text
	IsParadox            bool
	ParadoxSpecies       []string
	IsConvergent         bool
	ConvergentSpecies    []string
	IsSpecialAbilityForm bool
	IsCosmeticForm       bool
	IsFemaleForm         bool
	HasGenderDifferences bool
	IsBattleOnlyForm     bool
	IsSwitchableForm     bool
	IsFusion             bool
	FusedWith            []string
	IsMega               bool
	IsPrimal             bool
	IsGmax               bool
	IsRegional           bool
	CanGmax              bool
	CanDynamax           bool
	CanBeAlpha           bool
	DebutIn              string
	ObtainableIn         []string
	VersionExclusiveIn   []string
	EventOnlyIn          []string
	StorableIn           []string
	RegistrableIn        []string
	ShinyReleased        bool
	ShinyBase            pgtype.Text
	BaseStats            pokemon.PokemonStats
	Weight               int32
	Height               int32
	MaleRate             float32
	FemaleRate           float32
	BaseSpecies          pgtype.Text
	BaseForms            []string
	Forms                []string
	Family               pgtype.Text
	EvolvesFrom          *pokemon.EvolvesFrom
	Refs                 pokemon.ExternalRefs
}
