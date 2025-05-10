package dbquery

import (
	"database/sql"
	"fmt"
)

func PrintAll(db *sql.DB) error {
	rows, err := db.Query("SELECT name, time FROM trackTimes")
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var name, time string
	fmt.Println("name\ttime")
	for rows.Next() {
		if err := rows.Scan(&name, &time); err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}
		fmt.Printf("%s\t%s\n", name, time)
	}
	return rows.Err()
}
