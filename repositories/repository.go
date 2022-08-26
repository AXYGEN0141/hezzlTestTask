package repositories

import (
	"database/sql"
	"fmt"
	pb "github.com/AXYGEN0141/hezzlTestTask/protoUsers"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "11534506"
	dbname   = "postgres"

	connectionString = "host=%s port=%d userStruct=%s password=%s dbname=%s sslmode=disable"

	dbInsertUserFormat = `INSERT INTO user_table ( "id", "name") VALUES ( %d ,'%s');`
	dbDeleteUserFormat = `DELETE FROM user_table WHERE id = %d;`
)

type Repo struct {
	DbStruct *sql.DB
}

func New() *Repo {
	connectionString := fmt.Sprintf(connectionString, host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	repo := Repo{
		DbStruct: db,
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &repo
}

func (repo *Repo) AddUser(user *pb.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbInsertRequest := fmt.Sprintf(dbInsertUserFormat, user.Id, user.Name)

	_, err := repo.DbStruct.ExecContext(
		ctx,
		dbInsertRequest,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repo) DeleteUser(id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbDeleteRequest := fmt.Sprintf(dbDeleteUserFormat, id)

	_, err := repo.DbStruct.ExecContext(
		ctx,
		dbDeleteRequest,
	)
	if err != nil {
		return err
	}
	return nil
}
