package storage

import (
"database/sql"
"fmt"
"os"
"path/filepath"
"sync"

_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
db   *sql.DB
path string
mu   sync.RWMutex
}

func NewSQLiteStorage(path string) (*SQLiteStorage, error) {
dir := filepath.Dir(path)
if err := os.MkdirAll(dir, 0755); err != nil {
return nil, fmt.Errorf("create storage dir: %w", err)
}

db, err := sql.Open("sqlite3", path)
if err != nil {
return nil, fmt.Errorf("open database: %w", err)
}

db.SetMaxOpenConns(1)

s := &SQLiteStorage{
db:   db,
path: path,
}

if err := s.migrate(); err != nil {
db.Close()
return nil, fmt.Errorf("run migrations: %w", err)
}

return s, nil
}

func (s *SQLiteStorage) migrate() error {
migrations := []string{
`CREATE TABLE IF NOT EXISTS extensions (
name TEXT PRIMARY KEY,
version TEXT NOT NULL,
enabled INTEGER DEFAULT 0,
config TEXT,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
)`,
`CREATE TABLE IF NOT EXISTS tasks (
id TEXT PRIMARY KEY,
parent_id TEXT,
title TEXT NOT NULL,
description TEXT,
status TEXT DEFAULT 'pending',
type TEXT,
priority INTEGER DEFAULT 0,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
completed_at DATETIME
)`,
`CREATE TABLE IF NOT EXISTS task_dependencies (
task_id TEXT,
depends_on TEXT,
PRIMARY KEY (task_id, depends_on)
)`,
`CREATE TABLE IF NOT EXISTS events (
id INTEGER PRIMARY KEY AUTOINCREMENT,
type TEXT NOT NULL,
source TEXT NOT NULL,
timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
data TEXT
)`,
`CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status)`,
`CREATE INDEX IF NOT EXISTS idx_events_type ON events(type)`,
}

for _, migration := range migrations {
if _, err := s.db.Exec(migration); err != nil {
return fmt.Errorf("migration failed: %w", err)
}
}

return nil
}

func (s *SQLiteStorage) Execute(query string, args ...interface{}) (sql.Result, error) {
s.mu.Lock()
defer s.mu.Unlock()
return s.db.Exec(query, args...)
}

func (s *SQLiteStorage) Query(query string, args ...interface{}) (*sql.Rows, error) {
s.mu.RLock()
defer s.mu.RUnlock()
return s.db.Query(query, args...)
}

func (s *SQLiteStorage) QueryRow(query string, args ...interface{}) *sql.Row {
s.mu.RLock()
defer s.mu.RUnlock()
return s.db.QueryRow(query, args...)
}

func (s *SQLiteStorage) Begin() (*sql.Tx, error) {
s.mu.Lock()
defer s.mu.Unlock()
return s.db.Begin()
}

func (s *SQLiteStorage) Close() error {
return s.db.Close()
}
