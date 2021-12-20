// db is at the global scope to pass into our test functions
var db *sql.DB

func TestMain(m *testing.M) {
    pool, err := dockertest.NewPool("")
    if err != nil {
        log.Fatalf("Could not connect to docker: %s", err)
    }

    path, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    options := &dockertest.RunOptions{
        Repository: "postgres",
        Tag:        "12.3",
        Env: []string{
            "POSTGRES_USER=user",
            "POSTGRES_PASSWORD=secret",
            "listen_addresses = '*'",
        },
        ExposedPorts: []string{"5432"},
        PortBindings: map[docker.Port][]docker.PortBinding{
            "5432": {
                {HostIP: "0.0.0.0", HostPort: "5432"},
            },
        },
        Mounts: []string{fmt.Sprintf("%s/stub:/docker-entrypoint-initdb.d", path)},
    }

    resource, err := pool.RunWithOptions(options, func(config *docker.HostConfig) {
        config.AutoRemove = true
        config.RestartPolicy = docker.RestartPolicy{Name: "no"}
    })
    if err != nil {
        log.Fatalf("Could not start resource: %s", err)
    }

    err = resource.Expire(30)
    if err != nil {
        log.Fatalf("Could not expire resource: %s", err)
    }

    if err := pool.Retry(func() error {
        var err error
        db, err = sql.Open(
            "postgres",
            "postgres://user:secret@localhost:5432?sslmode=disable",
        )
        if err != nil {
            return err
        }
        return db.Ping()
    }); err != nil {
        log.Fatalf("Could not connect to database: %s", err)
    }

    code := m.Run()

    if err := pool.Purge(resource); err != nil {
        log.Fatalf("Could not purge resource: %s", err)
    }

    os.Exit(code)
}
