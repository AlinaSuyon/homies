package main

import (
	"awesomeProject1/apartment"
	"awesomeProject1/family"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func main() {
	db := setupDB()
	defer db.Close()

	createTable(db)
	insertFamilyData(db)
	printFamilyMembers(db)

	familyMembers := family.FamilyMembers()

	// Print family members individually
	for _, member := range familyMembers {
		fmt.Printf("Name: %s, Age: %d, Role: %s, Weight: %.2f\n", member.Name, member.Age, member.Role, member.Weight)
	}

	apartmentInstance := apartment.NewApartment()
	apartmentInstance.PrintDescription()

	for {
		time.Sleep(10 * time.Second)
	}
}
func setupDB() *sql.DB {
	connectionString := "host=db user=alinauser password=1356 dbname=alinadatabase sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func createTable(db *sql.DB) {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS family_members (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			age INT,
			role VARCHAR(100),
			weight FLOAT
		)
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Таблица 'family_members' создана успешно.")
}

func insertFamilyData(db *sql.DB) {
	familyMembers := family.FamilyMembers()

	for _, member := range familyMembers {
		insertDataQuery := `
			INSERT INTO family_members (name, age, role, weight) VALUES
				($1, $2, $3, $4)
		`

		_, err := db.Exec(insertDataQuery, member.Name, member.Age, member.Role, member.Weight)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Данные из файла Family вставлены успешно.")
}

func printFamilyMembers(db *sql.DB) {
	selectDataQuery := "SELECT * FROM family_members"
	rows, err := db.Query(selectDataQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nСемья:")
	for rows.Next() {
		var id, age int
		var name, role string
		var weight float64

		err := rows.Scan(&id, &name, &age, &role, &weight)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Имя: %s, Возраст: %d лет, Роль: %s, Вес: %.2f кг\n", id, name, age, role, weight)
	}
}
