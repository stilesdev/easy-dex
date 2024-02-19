package commands

import (
	"encoding/json"
	"fmt"
	"io/fs"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stilesdev/easy-dex/internal/database"
	"github.com/stilesdev/easy-dex/internal/database/sqlc"
	"github.com/stilesdev/easy-dex/internal/dataset"
	"github.com/stilesdev/easy-dex/internal/pokemon"
	"github.com/urfave/cli/v2"
)

func Dataset(db *database.Connection) *cli.Command {
    return &cli.Command{
        Name: "dataset",
        Aliases: []string{"ds"},
        Usage: "manage the source dataset for the app",
        UseShortOptionHandling: true,
        Subcommands: []*cli.Command{
            {
                Name: "import",
                Aliases: []string{"i"},
                Usage: "import the source dataset from json files",
                Action: func(cCtx *cli.Context) error {
                    index, err := readPokemonIndex()
                    if err != nil {
                        return err
                    }

                    _, err = db.Connection.Exec(db.Context, "TRUNCATE TABLE pokemon")
                    if err != nil {
                        return err
                    }

                    count := 0
                    for _, i := range index {
                        parsedPokemon, err := readPokemonInput(i.Region, i.Id)
                        if err != nil {
                            return err
                        }
                        _, err = insertPokemon(db, parsedPokemon)
                        if err != nil {
                            fmt.Println("Error inserting:", parsedPokemon.Id)
                            return err
                        }

                        //fmt.Printf("Inserted: %+v\n", pokemon.ID)
                        count++
                    }

                    fmt.Println("Inserted pokemon:", count)

                    return nil
                },
            },
            {
                Name: "list",
                Usage: "list all pokemon entries in the database",
                Action: func(cCtx *cli.Context) error {
                    queries := sqlc.New(db.Connection)

                    pokemonList, err := queries.GetPokemon(db.Context)
                    if err != nil {
                        return err
                    }

                    fmt.Println("Pokemon List:", pokemonList)

                    return nil
                },
            },
            {
                Name: "truncate",
                Usage: "delete all pokemon entries in the database",
                Action: func(cCtx *cli.Context) error {
                    _, err := db.Connection.Exec(db.Context, "TRUNCATE TABLE pokemon")

                    return err
                },
            },
        },
    }
}

func readPokemonIndex() ([]pokemon.IndexPokemon, error) {
    data := dataset.PokemonIndexJSON

    var index []pokemon.IndexPokemon
    err := json.Unmarshal(data, &index)
    if err != nil {
        return nil, err
    }

    return index, nil
}

func readPokemonInput(region string, id string) (pokemon.PokemonInput, error) {
    file := fmt.Sprintf("data/data/pokemon/%s/%s.json", region, id)

    data, err := fs.ReadFile(dataset.PokemonDataFS, file)
    if err != nil {
        return pokemon.PokemonInput{}, err
    }

    var pokemon pokemon.PokemonInput
    err = json.Unmarshal(data, &pokemon)

    return pokemon, err
}

func nullableString(val string) pgtype.Text {
    return pgtype.Text{String: val, Valid: len(val) > 0}
}

func insertPokemon(db *database.Connection, input pokemon.PokemonInput) (sqlc.Pokemon, error) {
    queries := sqlc.New(db.Connection)

    insertedPokemon, err := queries.CreatePokemon(db.Context, sqlc.CreatePokemonParams{
        ID                  : input.Id                  ,
        Nid                 : input.Nid                 ,
        DexNum              : input.DexNum              ,
        FormID              : nullableString(input.FormId),
        Name                : input.Name                ,
        FormName            : nullableString(input.FormName),
        Region              : input.Region              ,
        Generation          : input.Generation          ,
        Type1               : input.Type1               ,
        Type2               : nullableString(input.Type2),
        TeraType            : nullableString(input.TeraType),
        Color               : input.Color               ,
        PrimaryAbility      : input.Abilities.Primary   ,
        SecondaryAbility    : nullableString(input.Abilities.Secondary),
        HiddenAbility       : nullableString(input.Abilities.Hidden),
        IsDefault           : input.IsDefault           ,
        IsForm              : input.IsForm              ,
        IsLegendary         : input.IsLegendary         ,
        IsMythical          : input.IsMythical          ,
        IsUltraBeast        : input.IsUltraBeast        ,
        UltraBeastCode      : nullableString(input.UltraBeastCode),
        IsParadox           : input.IsParadox           ,
        ParadoxSpecies      : input.ParadoxSpecies      ,
        IsConvergent        : input.IsConvergent        ,
        ConvergentSpecies   : input.ConvergentSpecies   ,
        IsSpecialAbilityForm: input.IsSpecialAbilityForm,
        IsCosmeticForm      : input.IsCosmeticForm      ,
        IsFemaleForm        : input.IsFemaleForm        ,
        HasGenderDifferences: input.HasGenderDifferences,
        IsBattleOnlyForm    : input.IsBattleOnlyForm    ,
        IsSwitchableForm    : input.IsSwitchableForm    ,
        IsFusion            : input.IsFusion            ,
        FusedWith           : input.FusedWith           ,
        IsMega              : input.IsMega              ,
        IsPrimal            : input.IsPrimal            ,
        IsGmax              : input.IsGmax              ,
        IsRegional          : input.IsRegional          ,
        CanGmax             : input.CanGmax             ,
        CanDynamax          : input.CanDynamax          ,
        CanBeAlpha          : input.CanBeAlpha          ,
        DebutIn             : input.DebutIn             ,
        ObtainableIn        : input.ObtainableIn        ,
        VersionExclusiveIn  : input.VersionExclusiveIn  ,
        EventOnlyIn         : input.EventOnlyIn         ,
        StorableIn          : input.StorableIn          ,
        RegistrableIn       : input.RegistrableIn       ,
        ShinyReleased       : input.ShinyReleased       ,
        ShinyBase           : nullableString(input.ShinyBase),
        BaseStats           : input.BaseStats           ,
        Weight              : input.Weight              ,
        Height              : input.Height              ,
        MaleRate            : input.MaleRate            ,
        FemaleRate          : input.FemaleRate          ,
        BaseSpecies         : nullableString(input.BaseSpecies),
        BaseForms           : input.BaseForms           ,
        Forms               : input.Forms               ,
        Family              : nullableString(input.Family),
        EvolvesFrom         : input.EvolvesFrom         ,
        Refs                : input.Refs                ,
    })

    if err != nil {
        return sqlc.Pokemon{}, err
    }

    return insertedPokemon, nil
}
