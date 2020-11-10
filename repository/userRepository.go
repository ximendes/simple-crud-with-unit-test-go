//go:generate mockgen -destination=../mocks/mock_userRepository.go -package=mocks go-project/repository UserRepositoryInterface
package repository

import (
    "database/sql"
    "fmt"
    "github.com/joho/godotenv" // package used to read the .env file
    "go-project/models"
    "log"
    "os" // used to read the environment variable

    _ "github.com/lib/pq" // postgres golang driver
)

type UserRepositoryInterface interface {
    InsertUser(user models.User) int64
    GetUser(ud int64) (models.User, error)
    GetAllUsers() ([]models.User, error)
    UpdateUser(id int64, user models.User) int64
    DeleteUser(id int64) int64
}

type UserRepository struct{}

func (userRepository UserRepository) InsertUser(user models.User) int64 {

    db := userRepository.createConnection()

    defer db.Close()

    sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

    var id int64

    err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", id)

    return id
}

func (userRepository UserRepository) GetUser(id int64) (models.User, error) {
    db := userRepository.createConnection()

    defer db.Close()

    var user models.User

    sqlStatement := `SELECT * FROM users WHERE userid=$1`

    row := db.QueryRow(sqlStatement, id)

    err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return user, nil
    case nil:
        return user, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    return user, err
}

func (userRepository UserRepository) GetAllUsers() ([]models.User, error) {
    db := UserRepository{}.createConnection()

    defer db.Close()

    var users []models.User

    sqlStatement := `SELECT * FROM users`

    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    defer rows.Close()

    for rows.Next() {
        var user models.User

        err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        users = append(users, user)

    }

    return users, err
}

func (userRepository UserRepository) UpdateUser(id int64, user models.User) int64 {

    db := UserRepository{}.createConnection()

    defer db.Close()

    sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

    res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

func (userRepository UserRepository) DeleteUser(id int64) int64 {

    db := UserRepository{}.createConnection()

    defer db.Close()

    sqlStatement := `DELETE FROM users WHERE userid=$1`

    res, err := db.Exec(sqlStatement, id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

func (userRepository UserRepository) createConnection() *sql.DB {
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    return db
}
