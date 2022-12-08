package postgres

import (
	"fmt"
	"runtime"

	"auto_post/app/internal/adapters/port"
)

var _ port.Persister = (*SQLStore)(nil)

// UnitOfWork --
func (ths *SQLStore) UnitOfWork(fn func(persister port.Persister) error) (err error) {
	tx := ths.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()

			switch e := p.(type) {
			case runtime.Error:
				panic(e)
			case error:
				err = fmt.Errorf("panic err: %v", p)
				return
			default:
				panic(e)
			}
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	newStore := &SQLStore{
		db: tx,
	}
	err = fn(newStore)
	return err
}
