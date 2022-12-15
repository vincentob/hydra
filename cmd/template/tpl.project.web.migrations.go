package template

var (
	MigrationInitTpl = map[string]string{}

	MigrationAddTpl string
)

func init() {
	MigrationAddTpl = `// Generate By Template.
package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "{{ . }}",
		Migrate: func(tx *gorm.DB) error {
			// set tx default options
			tx = tx.Debug().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")

			// Use gorm auto migrate
			//if err := tx.AutoMigrate(&Category{}); err != nil {
			//	return err
			//}

			// Use sql migrate (use gorm better)
			//if err := tx.execute("sql"); err != nil {
			//	return err
			//}	

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("xxxxx")
		},
	})
}
`

	MigrationInitTpl["migrations.go"] = `// Generate By Template.
package migrations

import (
	"sort"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations []*gormigrate.Migration

func DoMigrate(db *gorm.DB) error {
	sort.Slice(Migrations, func(i, j int) bool {
		n := Migrations
		return n[i].ID < n[j].ID
	})

	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		Migrations,
	)

	return m.Migrate()
}
`
}
