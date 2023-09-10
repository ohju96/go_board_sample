package db

import (
	"context"
	"fmt"
	"ginSample/config"
	"ginSample/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MySQL *ent.Client

func InitMySQL(toml *config.Config) {
	mySqlUrl := toml.Local.User + ":" + toml.Local.Password + "@" + toml.Local.Host + "/" + toml.Local.Db + "?parseTime=True"

	client, err := ent.Open(toml.Local.Dbms, mySqlUrl)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	MySQL = client

	fmt.Println("mysql connected successfully !!")
}
