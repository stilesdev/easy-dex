// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: pokemon.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pokemon "github.com/stilesdev/easy-dex/internal/pokemon"
)

const createPokemon = `-- name: CreatePokemon :one
INSERT INTO pokemon (
    id,
    nid,
    dex_num,
    form_id,
    name,
    form_name,
    region,
    generation,
    type1,
    type2,
    tera_type,
    color,
    primary_ability,
    secondary_ability,
    hidden_ability,
    is_default,
    is_form,
    is_legendary,
    is_mythical,
    is_ultra_beast,
    ultra_beast_code,
    is_paradox,
    paradox_species,
    is_convergent,
    convergent_species,
    is_special_ability_form,
    is_cosmetic_form,
    is_female_form,
    has_gender_differences,
    is_battle_only_form,
    is_switchable_form,
    is_fusion,
    fused_with,
    is_mega,
    is_primal,
    is_gmax,
    is_regional,
    can_gmax,
    can_dynamax,
    can_be_alpha,
    debut_in,
    obtainable_in,
    version_exclusive_in,
    event_only_in,
    storable_in,
    registrable_in,
    shiny_released,
    shiny_base,
    base_stats,
    weight,
    height,
    male_rate,
    female_rate,
    base_species,
    base_forms,
    forms,
    family,
    evolves_from,
    refs
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17,
    $18,
    $19,
    $20,
    $21,
    $22,
    $23,
    $24,
    $25,
    $26,
    $27,
    $28,
    $29,
    $30,
    $31,
    $32,
    $33,
    $34,
    $35,
    $36,
    $37,
    $38,
    $39,
    $40,
    $41,
    $42,
    $43,
    $44,
    $45,
    $46,
    $47,
    $48,
    $49,
    $50,
    $51,
    $52,
    $53,
    $54,
    $55,
    $56,
    $57,
    $58,
    $59
) RETURNING id, nid, dex_num, form_id, name, form_name, region, generation, type1, type2, tera_type, color, primary_ability, secondary_ability, hidden_ability, is_default, is_form, is_legendary, is_mythical, is_ultra_beast, ultra_beast_code, is_paradox, paradox_species, is_convergent, convergent_species, is_special_ability_form, is_cosmetic_form, is_female_form, has_gender_differences, is_battle_only_form, is_switchable_form, is_fusion, fused_with, is_mega, is_primal, is_gmax, is_regional, can_gmax, can_dynamax, can_be_alpha, debut_in, obtainable_in, version_exclusive_in, event_only_in, storable_in, registrable_in, shiny_released, shiny_base, base_stats, weight, height, male_rate, female_rate, base_species, base_forms, forms, family, evolves_from, refs
`

type CreatePokemonParams struct {
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

func (q *Queries) CreatePokemon(ctx context.Context, arg CreatePokemonParams) (Pokemon, error) {
	row := q.db.QueryRow(ctx, createPokemon,
		arg.ID,
		arg.Nid,
		arg.DexNum,
		arg.FormID,
		arg.Name,
		arg.FormName,
		arg.Region,
		arg.Generation,
		arg.Type1,
		arg.Type2,
		arg.TeraType,
		arg.Color,
		arg.PrimaryAbility,
		arg.SecondaryAbility,
		arg.HiddenAbility,
		arg.IsDefault,
		arg.IsForm,
		arg.IsLegendary,
		arg.IsMythical,
		arg.IsUltraBeast,
		arg.UltraBeastCode,
		arg.IsParadox,
		arg.ParadoxSpecies,
		arg.IsConvergent,
		arg.ConvergentSpecies,
		arg.IsSpecialAbilityForm,
		arg.IsCosmeticForm,
		arg.IsFemaleForm,
		arg.HasGenderDifferences,
		arg.IsBattleOnlyForm,
		arg.IsSwitchableForm,
		arg.IsFusion,
		arg.FusedWith,
		arg.IsMega,
		arg.IsPrimal,
		arg.IsGmax,
		arg.IsRegional,
		arg.CanGmax,
		arg.CanDynamax,
		arg.CanBeAlpha,
		arg.DebutIn,
		arg.ObtainableIn,
		arg.VersionExclusiveIn,
		arg.EventOnlyIn,
		arg.StorableIn,
		arg.RegistrableIn,
		arg.ShinyReleased,
		arg.ShinyBase,
		arg.BaseStats,
		arg.Weight,
		arg.Height,
		arg.MaleRate,
		arg.FemaleRate,
		arg.BaseSpecies,
		arg.BaseForms,
		arg.Forms,
		arg.Family,
		arg.EvolvesFrom,
		arg.Refs,
	)
	var i Pokemon
	err := row.Scan(
		&i.ID,
		&i.Nid,
		&i.DexNum,
		&i.FormID,
		&i.Name,
		&i.FormName,
		&i.Region,
		&i.Generation,
		&i.Type1,
		&i.Type2,
		&i.TeraType,
		&i.Color,
		&i.PrimaryAbility,
		&i.SecondaryAbility,
		&i.HiddenAbility,
		&i.IsDefault,
		&i.IsForm,
		&i.IsLegendary,
		&i.IsMythical,
		&i.IsUltraBeast,
		&i.UltraBeastCode,
		&i.IsParadox,
		&i.ParadoxSpecies,
		&i.IsConvergent,
		&i.ConvergentSpecies,
		&i.IsSpecialAbilityForm,
		&i.IsCosmeticForm,
		&i.IsFemaleForm,
		&i.HasGenderDifferences,
		&i.IsBattleOnlyForm,
		&i.IsSwitchableForm,
		&i.IsFusion,
		&i.FusedWith,
		&i.IsMega,
		&i.IsPrimal,
		&i.IsGmax,
		&i.IsRegional,
		&i.CanGmax,
		&i.CanDynamax,
		&i.CanBeAlpha,
		&i.DebutIn,
		&i.ObtainableIn,
		&i.VersionExclusiveIn,
		&i.EventOnlyIn,
		&i.StorableIn,
		&i.RegistrableIn,
		&i.ShinyReleased,
		&i.ShinyBase,
		&i.BaseStats,
		&i.Weight,
		&i.Height,
		&i.MaleRate,
		&i.FemaleRate,
		&i.BaseSpecies,
		&i.BaseForms,
		&i.Forms,
		&i.Family,
		&i.EvolvesFrom,
		&i.Refs,
	)
	return i, err
}

const getPokemon = `-- name: GetPokemon :many
SELECT id, nid, dex_num, form_id, name, form_name, region, generation, type1, type2, tera_type, color, primary_ability, secondary_ability, hidden_ability, is_default, is_form, is_legendary, is_mythical, is_ultra_beast, ultra_beast_code, is_paradox, paradox_species, is_convergent, convergent_species, is_special_ability_form, is_cosmetic_form, is_female_form, has_gender_differences, is_battle_only_form, is_switchable_form, is_fusion, fused_with, is_mega, is_primal, is_gmax, is_regional, can_gmax, can_dynamax, can_be_alpha, debut_in, obtainable_in, version_exclusive_in, event_only_in, storable_in, registrable_in, shiny_released, shiny_base, base_stats, weight, height, male_rate, female_rate, base_species, base_forms, forms, family, evolves_from, refs FROM pokemon
`

func (q *Queries) GetPokemon(ctx context.Context) ([]Pokemon, error) {
	rows, err := q.db.Query(ctx, getPokemon)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pokemon
	for rows.Next() {
		var i Pokemon
		if err := rows.Scan(
			&i.ID,
			&i.Nid,
			&i.DexNum,
			&i.FormID,
			&i.Name,
			&i.FormName,
			&i.Region,
			&i.Generation,
			&i.Type1,
			&i.Type2,
			&i.TeraType,
			&i.Color,
			&i.PrimaryAbility,
			&i.SecondaryAbility,
			&i.HiddenAbility,
			&i.IsDefault,
			&i.IsForm,
			&i.IsLegendary,
			&i.IsMythical,
			&i.IsUltraBeast,
			&i.UltraBeastCode,
			&i.IsParadox,
			&i.ParadoxSpecies,
			&i.IsConvergent,
			&i.ConvergentSpecies,
			&i.IsSpecialAbilityForm,
			&i.IsCosmeticForm,
			&i.IsFemaleForm,
			&i.HasGenderDifferences,
			&i.IsBattleOnlyForm,
			&i.IsSwitchableForm,
			&i.IsFusion,
			&i.FusedWith,
			&i.IsMega,
			&i.IsPrimal,
			&i.IsGmax,
			&i.IsRegional,
			&i.CanGmax,
			&i.CanDynamax,
			&i.CanBeAlpha,
			&i.DebutIn,
			&i.ObtainableIn,
			&i.VersionExclusiveIn,
			&i.EventOnlyIn,
			&i.StorableIn,
			&i.RegistrableIn,
			&i.ShinyReleased,
			&i.ShinyBase,
			&i.BaseStats,
			&i.Weight,
			&i.Height,
			&i.MaleRate,
			&i.FemaleRate,
			&i.BaseSpecies,
			&i.BaseForms,
			&i.Forms,
			&i.Family,
			&i.EvolvesFrom,
			&i.Refs,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
