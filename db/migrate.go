package db

// func MigrateDB() {
// 	conf := config.GetConfig()

// 	ConnectDB()
// 	migrations := &migrate.FileMigrationSource{
// 		Dir: conf.MigrationSource,
// 	}

// 	_, err := migrate.Exec(Db, "postgres", migrations, migrate.Up)
// 	if err != nil {
// 		panic(err)
// 	}
// 	slog.Info("Successfully migrate database")
// }
