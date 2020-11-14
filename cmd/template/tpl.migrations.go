package template

var (
	MigrationInitTpl = map[string]string{}

	MigrationAddTpl string
)

func init() {
	MigrationAddTpl = `// Generate By Template.
package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "{{ . }}",
		Migrate: func(tx *gorm.DB) error {
			// set tx default options
			tx = tx.Debug().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")

			// Use gorm auto migrate
			//if err := tx.AutoMigrate(&Category{}).Error; err != nil {
			//	return err
			//}

			// Use sql migrate (use gorm better)
			//if err := tx.execute("sql").Error; err != nil {
			//	return err
			//}	

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("categories").Error
		},
	})
}
`

	MigrationInitTpl["migrations.go"] = `// Generate By Template.
package migrations

import (
	"sort"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var Migrations []*gormigrate.Migration

func DoMigrate(db *gorm.DB) error {
	sort.Slice(Migrations, func(i, j int) bool {
		n := Migrations
		return n[i].ID < n[j].ID
	})

	gormigrate.DefaultOptions.IDColumnSize = 50
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		Migrations,
	)

	return m.Migrate()
}`
}
