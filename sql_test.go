package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExectSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scr := "INSERT INTO customer(id, name) VALUES('joko', 'sembung')"
	_, err := db.ExecContext(ctx, scr)

	if err != nil {
		panic(err)
	}

	fmt.Println("insert successfully new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scr := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, scr)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("name :", name)
	}
	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scr := "SELECT id, name, email, balance, rating, birth_date, maried, created_at FROM customer"
	rows, err := db.QueryContext(ctx, scr)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date, created_at time.Time
		var maried bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &maried, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("name :", name)
		if email.Valid {
			fmt.Println("email :", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		fmt.Println("birth_date :", birth_date)
		fmt.Println("maried :", maried)
		fmt.Println("created_at :", created_at)
	}

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	scr := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, scr)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("berhasil login, anda adalah :", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestSqlParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	sclQuery := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sclQuery, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("berhasil login, anda adalah :", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "test@email.com"
	comment := "hello selamat belajar"

	scr := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, scr, email, comment)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("berhasil insert data baru, id :", id)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	scr := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, scr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "test" + strconv.Itoa(i) + "@gmail.com"
		commet := "ini komen ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, commet)
		if err != nil {
			panic(err)
		}
		lastInserId, _ := result.LastInsertId()
		fmt.Println("komen id :", lastInserId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	scr := "INSERT INTO comments(email, comment) VALUES(?,?)"

	// do transaction
	for i := 0; i < 10; i++ {
		email := "test" + strconv.Itoa(i) + "@gmail.com"
		commet := "ini komen ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, scr, email, commet)
		if err != nil {
			panic(err)
		}
		lastInserId, _ := result.LastInsertId()
		fmt.Println("komen id :", lastInserId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
