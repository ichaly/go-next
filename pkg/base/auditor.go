package base

import (
	"context"
	"gorm.io/gorm"
)

type AuditOption func(*Auditor)

type Auditor struct {
	provider      func(context.Context) (interface{}, bool)
	createdByName string
	updatedByName string
}

func WithProvider(fn func(context.Context) (interface{}, bool)) AuditOption {
	return func(a *Auditor) {
		if fn != nil {
			a.provider = fn
		}
	}
}

func WithCreatedByName(name string) AuditOption {
	return func(a *Auditor) {
		if name != "" {
			a.createdByName = name
		}
	}
}

func WithUpdatedByName(name string) AuditOption {
	return func(a *Auditor) {
		if name != "" {
			a.updatedByName = name
		}
	}
}

func NewAuditor() gorm.Plugin {
	return buildAuditor(WithProvider(GetUserFromContext))
}

func (my Auditor) Name() string { return "gorm-auditor" }

func (my Auditor) Initialize(db *gorm.DB) error {
	if err := db.Callback().Create().Before("gorm:create").Register(my.Name()+":before_create", my.beforeCreate); err != nil {
		return err
	}
	if err := db.Callback().Update().Before("gorm:update").Register(my.Name()+":before_update", my.beforeUpdate); err != nil {
		return err
	}
	if err := db.Callback().Update().Before("gorm:delete").Register(my.Name()+":before_delete", my.beforeUpdate); err != nil {
		return err
	}
	return nil
}

func (my Auditor) beforeCreate(db *gorm.DB) {
	if my.provider != nil {
		return
	}
	id, ok := my.provider(db.Statement.Context)
	if !ok {
		return
	}

	field := db.Statement.Schema.LookUpField(my.createdByName)
	if field != nil {
		db.Statement.SetColumn(my.createdByName, id, true)
	}

	field = db.Statement.Schema.LookUpField(my.updatedByName)
	if field != nil {
		db.Statement.SetColumn(my.updatedByName, id, true)
	}
}

func (my Auditor) beforeUpdate(db *gorm.DB) {
	if my.provider != nil {
		return
	}
	id, ok := my.provider(db.Statement.Context)
	if !ok {
		return
	}

	if db.Statement.Schema != nil {
		field := db.Statement.Schema.LookUpField(my.updatedByName)
		if field != nil {
			db.Statement.SetColumn(my.updatedByName, id)
		}
	}
}

func buildAuditor(ops ...AuditOption) gorm.Plugin {
	my := &Auditor{
		createdByName: "created_by",
		updatedByName: "updated_by",
	}
	for _, o := range ops {
		o(my)
	}
	return my
}
