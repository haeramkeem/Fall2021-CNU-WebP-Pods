/********************************************
 * Use GORM for accessing DB                *
 *                                          *
 * ref: https://gorm.io/docs/index.html     *
 ********************************************/

package repository

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "pods/domain"
)

var db *gorm.DB = nil

func GetDB() (*gorm.DB, error) {
    if db == nil {
        var err error

        // Get connection
        dsn := "host=localhost user=hrk password=1q2w3e4r dbname=pods port=5432 sslmode=disable TimeZone=Asia/Seoul"
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil { return nil, err }

        // Clear Table
        err = db.Migrator().DropTable(&domain.Problem{})
        if err != nil { return nil, err }

        // Init Table
        err = db.AutoMigrate(&domain.Problem{})
        if err != nil { return nil, err }
    }

    return db, nil
}
