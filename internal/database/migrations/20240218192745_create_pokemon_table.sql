-- migrate:up
CREATE TABLE pokemon (
    id VARCHAR NOT NULL,
    nid VARCHAR NOT NULL,
    dex_num INTEGER NOT NULL,
    form_id VARCHAR, -- nullable
    name VARCHAR NOT NULL,
    form_name VARCHAR, -- nullable
    region VARCHAR NOT NULL,
    generation INTEGER NOT NULL,
    type1 VARCHAR NOT NULL,
    type2 VARCHAR, -- nullable
    tera_type VARCHAR, -- nullable
    color VARCHAR NOT NULL,
    primary_ability VARCHAR NOT NULL,
    secondary_ability VARCHAR, -- nullable
    hidden_ability VARCHAR, -- nullable
    is_default BOOL NOT NULL,
    is_form BOOL NOT NULL,
    is_legendary BOOL NOT NULL,
    is_mythical BOOL NOT NULL,
    is_ultra_beast BOOL NOT NULL,
    ultra_beast_code VARCHAR, -- nullable
    is_paradox BOOL DEFAULT FALSE NOT NULL,
    paradox_species VARCHAR[], -- nullable
    is_convergent BOOL DEFAULT FALSE NOT NULL,
    convergent_species VARCHAR[], -- nullable
    is_special_ability_form BOOL NOT NULL,
    is_cosmetic_form BOOL NOT NULL,
    is_female_form BOOL NOT NULL,
    has_gender_differences BOOL NOT NULL,
    is_battle_only_form BOOL NOT NULL,
    is_switchable_form BOOL NOT NULL,
    is_fusion BOOL NOT NULL,
    fused_with VARCHAR[], -- nullable
    is_mega BOOL NOT NULL,
    is_primal BOOL NOT NULL,
    is_gmax BOOL NOT NULL,
    is_regional BOOL NOT NULL,
    can_gmax BOOL NOT NULL,
    can_dynamax BOOL NOT NULL,
    can_be_alpha BOOL NOT NULL,
    debut_in VARCHAR NOT NULL,
    obtainable_in VARCHAR[] NOT NULL,
    version_exclusive_in VARCHAR[] NOT NULL,
    event_only_in VARCHAR[] NOT NULL,
    storable_in VARCHAR[] NOT NULL,
    registrable_in VARCHAR[] NOT NULL,
    shiny_released BOOL NOT NULL,
    shiny_base VARCHAR, -- nullable
    base_stats JSON NOT NULL, --hp,atk,def,spa,spd,spe
    weight INTEGER NOT NULL,
    height INTEGER NOT NULL,
    male_rate REAL NOT NULL,
    female_rate REAL NOT NULL,
    base_species VARCHAR, -- nullable
    base_forms VARCHAR[], -- nullable
    forms VARCHAR[], -- nullable
    family VARCHAR, -- nullable
    evolves_from JSON, -- nullable
    refs JSON NOT NULL
);

-- migrate:down
DROP TABLE pokemon;

