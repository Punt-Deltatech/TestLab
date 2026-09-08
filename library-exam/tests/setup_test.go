package tests

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

var cache sync.Map

func parse(t *testing.T, model interface{}) *schema.Schema {
	t.Helper()
	sch, err := schema.Parse(model, &cache, schema.NamingStrategy{})
	if err != nil {
		t.Fatalf("parse schema %T: %v", model, err)
	}
	return sch
}

func hasUnique(f *schema.Field) bool {
	if f == nil {
		return false
	}
	if f.Unique || f.UniqueIndex != "" {
		return true
	}
	if _, ok := f.TagSettings["UNIQUEINDEX"]; ok {
		return true
	}
	if _, ok := f.TagSettings["UNIQUE"]; ok {
		return true
	}
	return false
}

func isNotNull(f *schema.Field) bool {
	if f == nil {
		return false
	}
	if f.NotNull {
		return true
	}
	if _, ok := f.TagSettings["NOT NULL"]; ok {
		return true
	}
	if _, ok := f.TagSettings["NOTNULL"]; ok {
		return true
	}
	return false
}
