package postgres

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/config"
	status "auto_post/app/pkg/vars/statuses"
	log "git.fintechru.org/dfa/dfa_lib/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MigrateEnum - интерфейс для перечислений (enum)
type MigrateEnum interface {
	MigrationSQL() string
}

// SQLStore fulfills the Extractor and Persister document interfaces
type SQLStore struct {
	db  *gorm.DB
	log log.Logger
}

// NewPostgresRepository returns a memory repository instance
func NewPostgresRepository(cfg config.Config, log log.Logger, migrate bool) (*SQLStore, error) {
	dbGorm, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.CreateDSN(),
	}))
	if err != nil {
		return nil, err
	}

	postgresDb, err := dbGorm.DB()
	if err != nil {
		return nil, err
	}

	postgresDb.SetMaxIdleConns(10)
	postgresDb.SetMaxOpenConns(100)

	if migrate {
		if err = migrateData(dbGorm); err != nil {
			return nil, err
		}
	}

	return &SQLStore{
		db:  dbGorm,
		log: log,
	}, nil
}

func migrateData(db *gorm.DB) error {
	// Создаём перечисления в БД перед миграциями основных моделей
	enums := []MigrateEnum{
		new(status.ParseImageStatusEnum),
	}

	for _, enum := range enums {
		if err := db.Exec(enum.MigrationSQL()).Error; err != nil {
			return err
		}
	}

	return db.AutoMigrate(
		&models.ParseImage{},
	)
}