package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stilesdev/easy-dex/internal/commands"
	"github.com/stilesdev/easy-dex/internal/database"
	"github.com/urfave/cli/v2"
)


func main() {
    godotenv.Load()

    db, err := database.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalln(err)
    }
    defer db.Connection.Close()

    err = db.Connection.Ping(db.Context)
    if err != nil {
        log.Fatalln(err)
    }

    app := &cli.App{
        Name: "easy-dex",
        Usage: "easy-dex",
        Commands: []*cli.Command{
            commands.Dataset(db),
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatalln(err)
    }
}
