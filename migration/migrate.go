package migration

import (
  "github.com/golang-migrate/migrate/v4"
  _ "github.com/golang-migrate/migrate/v4/database/postgres"
  _ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp() error {
  
  m, err := migrate.New("file://migrations", "postgres://...")
  if err != nil {
    return err
  }
  
  if err := m.Up(); err && err != migrate.ErrNoChange {
    return err
  }
  
  return nil
}

func MigrateDown() error {
  m, err := migrate.New("file://migrations", "postgres://...")
  if err := m.Down(); err && err != migrate.ErrNoChange {
    return err
  }
  
  return nil
}