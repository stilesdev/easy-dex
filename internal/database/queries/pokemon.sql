-- name: GetPokemon :many
SELECT * FROM pokemon;

-- name: CreatePokemon :one
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
) RETURNING *;

