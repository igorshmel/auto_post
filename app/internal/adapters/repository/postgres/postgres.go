package postgres

import (
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/types"
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
	log logger.Logger
}

// NewPostgresRepository returns a memory repository instance
func NewPostgresRepository(cfg config.Config, log logger.Logger, migrate bool) (*SQLStore, error) {
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

	err = dbGorm.Migrator().DropTable("art_publish_counter")
	if err != nil {
		fmt.Println("Error Drop Table file")
	}
	err = dbGorm.Migrator().DropTable("cron_counter")
	if err != nil {
		fmt.Println("Error Drop Table file")
	}

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
		new(status.RecordStatusEnum),
		new(types.PublishTypeEnum),
	}

	for _, enum := range enums {
		if err := db.Exec(enum.MigrationSQL()).Error; err != nil {
			return err
		}
	}

	return db.AutoMigrate(
		&models.Manager{},
		&models.PublishCounter{},
	)
}
