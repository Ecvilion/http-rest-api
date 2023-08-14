package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Инициализация, открытие хранилища + отлов ошибок
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Отключение от базы данных при завершении работы сервера
func (s *Store) Close() {
	s.db.Close()
}

// Нашими репозиториями не должны пользоваться в обход хранилища store.User().Create()
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
