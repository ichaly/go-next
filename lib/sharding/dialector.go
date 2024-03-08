package sharding

import "gorm.io/gorm"

func NewShardingDialector(d gorm.Dialector, s *Sharding) ShardingDialector {
	return ShardingDialector{
		Dialector: d,
		sharding:  s,
	}
}

type ShardingDialector struct {
	gorm.Dialector
	sharding *Sharding
}

func (my ShardingDialector) Migrator(db *gorm.DB) gorm.Migrator {
	m := my.Dialector.Migrator(db)
	return ShardingMigrator{
		Migrator:  m,
		sharding:  my.sharding,
		dialector: my.Dialector,
	}
}

type ShardingMigrator struct {
	gorm.Migrator
	sharding  *Sharding
	dialector gorm.Dialector
}

func (my ShardingMigrator) AutoMigrate(dst ...any) error {
	return my.Migrator.AutoMigrate(dst)
}

func (my ShardingMigrator) DropTable(dst ...any) error {
	return my.Migrator.DropTable(dst)
}
