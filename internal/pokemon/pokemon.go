package pokemon

// https://www.jsdocs.io/package/@supeffective/dataset#pokemonSchema

type IndexPokemon struct {
    Id string
    Region string
    Name string
    Nid string
    IsForm bool
}

type PokemonAbilities struct {
    Primary   string           `json:"primary"`
    Secondary string           `json:"secondary"` // nullable
    Hidden    string           `json:"hidden"`    // nullable
}

type PokemonStats struct {
    Hp  int                    `json:"hp"`
    Atk int                    `json:"atk"`
    Def int                    `json:"def"`
    Spa int                    `json:"spa"`
    Spd int                    `json:"spd"`
    Spe int                    `json:"spe"`
}

type EvolvesFrom struct {
    Pokemon string             `json:"pokemon"`   // nullable
    Level int                  `json:"level"`     // nullable
    Item string                `json:"item"`      // nullable
    Move string                `json:"move"`      // nullable
    Type string                `json:"type"`      // nullable
    Region string              `json:"region"`    // nullable
    Ability string             `json:"ability"`   // nullable
    Condition string           `json:"condition"` // nullable
}

type ExternalRefs struct {
    Smogon string              `json:"smogon"`
    Showdown string            `json:"showdown"`
    ShowdownName string        `json:"showdownName"`
    Serebii string             `json:"serebii"`
    Bulbapedia string          `json:"bulbapedia"`
}

type PokemonInput struct {
    Id string
    Nid string
    DexNum int32
    FormId string                                 // nullable
    Name string
    FormName string                               // nullable
    Region string
    Generation int32
    Type1 string
    Type2 string                                  // nullable
    TeraType string                               // nullable
    Color string
    Abilities PokemonAbilities
    IsDefault bool
    IsForm bool
    IsLegendary bool
    IsMythical bool
    IsUltraBeast bool
    UltraBeastCode string                         // nullable
    IsParadox bool                                // nullable
    ParadoxSpecies []string                       // nullable
    IsConvergent bool                             // nullable
    ConvergentSpecies []string                    // nullable
    IsSpecialAbilityForm bool
    IsCosmeticForm bool
    IsFemaleForm bool
    HasGenderDifferences bool
    IsBattleOnlyForm bool
    IsSwitchableForm bool
    IsFusion bool
    FusedWith []string                            // nullable
    IsMega bool
    IsPrimal bool
    IsGmax bool
    IsRegional bool
    CanGmax bool
    CanDynamax bool
    CanBeAlpha bool
    DebutIn string
    ObtainableIn []string
    VersionExclusiveIn []string
    EventOnlyIn []string
    StorableIn []string
    RegistrableIn []string
    ShinyReleased bool
    ShinyBase string                              // nullable
    BaseStats PokemonStats
    Weight int32
    Height int32
    MaleRate float32
    FemaleRate float32
    BaseSpecies string                            // nullable
    BaseForms []string                            // nullable
    Forms []string                                // nullable
    Family string                                 // nullable
    EvolvesFrom *EvolvesFrom                      // nullable
    Refs ExternalRefs
}
