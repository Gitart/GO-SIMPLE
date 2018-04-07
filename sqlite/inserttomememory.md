
## Insert to memory

```golang

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
```

```bach
lix@900X4C:~/Go/src$ go tool pprof --text ./test /tmp/profile102098478/mem.pprof
Adjusting heap profiles for 1-in-4096 sampling rate
Total: 0.0 MB
     0.0 100.0% 100.0%      0.0 100.0% runtime.allocm
     0.0   0.0% 100.0%      0.0 100.0% database/sql.(*DB).Exec
     0.0   0.0% 100.0%      0.0 100.0% database/sql.(*DB).conn
     0.0   0.0% 100.0%      0.0 100.0% database/sql.(*DB).exec
     0.0   0.0% 100.0%      0.0 100.0% github.com/mattn/go-sqlite3.(*SQLiteDriver).Open
     0.0   0.0% 100.0%      0.0 100.0% github.com/mattn/go-sqlite3._Cfunc_sqlite3_threadsafe
     0.0   0.0% 100.0%      0.0 100.0% main.main
     0.0   0.0% 100.0%      0.0 100.0% runtime.cgocall
     0.0   0.0% 100.0%      0.0 100.0% runtime.gosched0
     0.0   0.0% 100.0%      0.0 100.0% runtime.main
     0.0   0.0% 100.0%      0.0 100.0% runtime.newextram
```

## Links
https://stackoverflow.com/questions/26456253/sqlite-3-not-releasing-memory-in-golang  
https://github.com/mattn/go-sqlite3  



