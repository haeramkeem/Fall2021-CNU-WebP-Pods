/********************************************
 * Use GORM for accessing DB                *
 *                                          *
 * ref: https://gorm.io/docs/index.html     *
 ********************************************/

package repository

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  . "pods/domain"
  "errors"
  e "pods/modules/errors"
)

var db *gorm.DB = nil

/**
 * Get connection with DB
 * @return *gorm.DB: connection with DB
 * @return error: generated error if exists
 */
func getDB() (*gorm.DB, error) {
    if db == nil {
        var err error

        // Get connection
        dsn := "host=localhost user=hrk password=1q2w3e4r dbname=pods port=5432 sslmode=disable TimeZone=Asia/Seoul"
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil { return nil, err }

        // Clear Table
        err = db.Migrator().DropTable(&ProblemDB{})
        if err != nil { return nil, err }

        // Init Table
        err = db.AutoMigrate(&ProblemDB{})
        if err != nil { return nil, err }
    }

    return db, nil
}

/**
 * Create problem record
 * @params problem *domain.ProblemDB: target problem structure
 * @return error: generated error if exists
 */
func CreateProblem(problem *ProblemDB) error {
    db, err := getDB()
    if err := e.Check(err); err != nil {
        return err
    }

    res := db.Create(problem)
    if err := e.Check(res.Error); err != nil {
        return err
    }

    return nil
}

/**
 * Select one problem from DB
 *      one or less problem must be selected by combination of date & category
 * @params date string: serving date of problem
 * @params category uint: category of problem
 * @return *domain.ProblemDB: selected record
 * @return error: generated error if exists
 */
func SelectProblem(date string, category uint) (*ProblemDB, error) {
    db, err := getDB()
    if err := e.Check(err); err != nil {
        return nil, err
    }

    var problem ProblemDB
    res := db.Where("date = ? AND category = ?", date, category).First(&problem)

    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        return nil, nil
    } else if err := e.Check(res.Error); err != nil {
        return nil, err
    }

    return &problem, nil
}
