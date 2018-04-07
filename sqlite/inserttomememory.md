package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./test.db"); if err != nil {
        log.Fatal(err)
    }; defer db.Close()

    ddl := `
        PRAGMA automatic_index = ON;
        PRAGMA cache_size = 32768;
        PRAGMA cache_spill = OFF;
        PRAGMA foreign_keys = ON;
        PRAGMA journal_size_limit = 67110000;
        PRAGMA locking_mode = NORMAL;
        PRAGMA page_size = 4096;
        PRAGMA recursive_triggers = ON;
        PRAGMA secure_delete = ON;
        PRAGMA synchronous = NORMAL;
        PRAGMA temp_store = MEMORY;
        PRAGMA journal_mode = WAL;
        PRAGMA wal_autocheckpoint = 16384;

        CREATE TABLE IF NOT EXISTS "user" (
            "id" TEXT,
            "username" TEXT,
            "password" TEXT
        );

        CREATE UNIQUE INDEX IF NOT EXISTS "id" ON "user" ("id");
    `

    _, err = db.Exec(ddl); if err != nil {
        log.Fatal(err)
    }

    queries := map[string]*sql.Stmt{}

    queries["user"], _ = db.Prepare(`INSERT OR REPLACE INTO "user" VALUES (?, ?, ?);`); if err != nil {
        log.Fatal(err)
    }; defer queries["user"].Close()

    tx, err := db.Begin(); if err != nil {
        log.Fatal(err)
    }

    for i := 0; i < 10000000; i++ {
        user := map[string]string{
            "id": string(i),
            "username": "foo",
            "password": "bar",
        }

        _, err := tx.Stmt(queries["user"]).Exec(user["id"], user["username"], user["password"]); if err != nil {
            log.Fatal(err)
        }

        if i % 32768 == 0 {
            tx.Commit()
            db.Exec(`PRAGMA shrink_memory;`)

            tx, err = db.Begin(); if err != nil {
                log.Fatal(err)
            }

            fmt.Println(i)
        }
    }

    tx.Commit()
}
