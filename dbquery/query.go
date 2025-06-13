package dbquery

import (
	"database/sql"
	"fmt"
	"maWeb/model"
)

func GetTrackTimes(db *sql.DB) ([]model.TrackTime, error) {
	rows, err := db.Query("SELECT name, time FROM trackTimes")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var data []model.TrackTime
	for rows.Next() {
		var t model.TrackTime
		if err := rows.Scan(&t.Name, &t.Time); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		data = append(data, t)
	}
	return data, nil
}

func GetStar(db *sql.DB) ([]model.TrackTime, error) {
	rows, err := db.Query("SELECT name, club, time FROM times")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var data []model.TrackTime
	for rows.Next() {
		var t model.TrackTime
		if err := rows.Scan(&t.Name, &t.Club, &t.Time); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		data = append(data, t)
	}
	return data, nil
}

func PrintStar(db *sql.DB) error {
	data, err := GetStar(db)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}

	fmt.Println("name\tclub\ttime")
	for _, item := range data {
		fmt.Printf("%s\t%d\t%s\n", item.Name, item.Club, item.Time)
	}
	println()
	return nil
}
