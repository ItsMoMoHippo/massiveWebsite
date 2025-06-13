package util

import (
	"database/sql"
	"fmt"
	"maWeb/dbquery"
	"maWeb/model"
)

func PrintCache(cache []model.TrackTime) {
	println("name\tclub\ttime")
	for _, unit := range cache {
		fmt.Printf("%s\t%d\t%s\n", unit.Name, unit.Club, unit.Time)
	}
	println()
}

func FilterCache(cache *[]model.TrackTime, clubID int) *[]model.TrackTime {
	var filtered []model.TrackTime
	for _, entry := range *cache {
		if entry.Club == clubID {
			filtered = append(filtered, entry)
		}
	}
	return &filtered
}

func RefreshCache(db *sql.DB) ([]model.TrackTime, error) {
	newCache, err := dbquery.GetStar(db)
	if err != nil {
		return nil, fmt.Errorf("cache refresh failed: %w", err)
	}
	return newCache, nil
}
