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

func field(t *testing.T, sch *schema.Schema, dbName string) *schema.Field {
	t.Helper()
	f, ok := sch.FieldsByDBName[dbName]
	if !ok || f.DBName == "" {
		t.Fatalf("[%s] ไม่พบคอลัมน์ %q", sch.Table, dbName)
	}
	return f
}

func hasUnique(f *schema.Field) bool {
	if f == nil {
		return false
	}
	if f.Unique || f.UniqueIndex != "" {
		return true
	}
	for k := range f.TagSettings {
		if k == "UNIQUEINDEX" || k == "UNIQUE" {
			return true
		}
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
	_, a := f.TagSettings["NOT NULL"]
	_, b := f.TagSettings["NOTNULL"]
	return a || b
}

func pkNames(sch *schema.Schema) []string {
	out := make([]string, 0, len(sch.PrimaryFields))
	for _, f := range sch.PrimaryFields {
		out = append(out, f.DBName)
	}
	return out
}

func contains(ss []string, want string) bool {
	for _, s := range ss {
		if s == want {
			return true
		}
	}
	return false
}
