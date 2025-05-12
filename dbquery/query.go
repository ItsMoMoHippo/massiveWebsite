package dbquery

import (
	"database/sql"
	"fmt"
)

type TrackTime struct {
	Name string
	Time string
}

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

func GetTrackTimes(db *sql.DB) ([]TrackTime, error) {
	rows, err := db.Query("SELECT name, time FROM trackTimes")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var data []TrackTime
	for rows.Next() {
		var t TrackTime
		if err := rows.Scan(&t.Name, &t.Time); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		data = append(data, t)
	}
	return data, nil
}
