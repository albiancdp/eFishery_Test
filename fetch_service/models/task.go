package models

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type ResponseBody struct {
	Uuid         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUSD     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type ResponseBodyList struct {
	AreaProvinsi string  `json:"area_provinsi"`
	Year         string  `json:"year"`
	Month        string  `json:"month"`
	Week         string  `json:"week"`
	SizeMin      int     `json:"size_min"`
	SizeMax      int     `json:"size_max"`
	SizeMed      string  `json:"size_med"`
	SizeAvg      float32 `json:"size_avg"`
	PriceMin     int     `json:"price_min"`
	PriceMax     int     `json:"price_max"`
	PriceMed     string  `json:"price_med"`
	PriceAvg     float32 `json:"price_avg"`
}

// type ResponseBodyList struct {
// 	AreaProvinsi string     `json:"area_provinsi"`
// 	Year         string     `json:"year"`
// 	Month        string     `json:"month"`
// 	Week         string     `json:"week"`
// 	Size         SizeModel  `json:"size"`
// 	Price        PriceModel `json:"price"`
// }

// type SizeModel struct {
// 	SizeMin int     `json:"size_min"`
// 	SizeMax int     `json:"size_max"`
// 	SizeMed string  `json:"size_med"`
// 	SizeAvg float32 `json:"size_avg"`
// }

// type PriceModel struct {
// 	PriceMin int     `json:"price_min"`
// 	PriceMax int     `json:"price_max"`
// 	PriceMed string  `json:"price_med"`
// 	PriceAvg float32 `json:"price_avg"`
// }

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./task.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func Migration() error {
	sql := `CREATE TABLE IF NOT EXISTS tasks (
		uuid 			TEXT 		UNIQUE,
		komoditas 		TEXT,
		area_provinsi 	TEXT,
		area_kota 		TEXT,
		size 			INT,
		price 			INT,
		price_usd 		TEXT,
		tgl_parsed 		DATE,
		timestamp 		INT
	);`

	_, err := DB.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func CreateTask(task ResponseBody) (int64, error) {
	result, err := DB.Exec("INSERT INTO tasks(uuid, komoditas, area_provinsi, area_kota, size, price, price_usd, tgl_parsed, timestamp) VALUES (?, ?, ?, ?, CAST(? AS INTEGER), CAST(? AS INTEGER), ?, DATETIME(?), CAST(? AS INTEGER))",
		task.Uuid, task.Komoditas, task.AreaProvinsi, task.AreaKota, task.Size, task.Price, task.PriceUSD, strings.Replace(task.TglParsed, "+07:00", "", 1), task.Timestamp)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func GetTasks() ([]ResponseBodyList, error) {
	rows, err := DB.Query(`SELECT 
								area_provinsi, 
								STRFTIME('%Y', tgl_parsed) AS year,
								STRFTIME('%m', tgl_parsed) AS month,
								STRFTIME('%W', tgl_parsed) AS week,
								min(size) AS size_min,
								max(size) AS size_max,
								"Median Belum bisa" AS size_med,
								avg(size) AS size_avg,
								min(price) AS price_min,
								max(price) AS price_max,
								"Median Belum bisa" AS price_med,
								avg(price) AS price_avg
							FROM tasks 
							GROUP BY area_provinsi, week
							ORDER BY area_provinsi,week;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println(rows)
	tasks := []ResponseBodyList{}
	for rows.Next() {
		var task ResponseBodyList
		// if err := rows.Scan(&task.Uuid, &task.Komoditas, &task.AreaProvinsi, &task.AreaKota, &task.Size, &task.Price, &task.PriceUSD, &task.TglParsed, &task.Timestamp); err != nil {
		err := rows.Scan(&task.AreaProvinsi, &task.Year, &task.Month, &task.Week, &task.SizeMin, &task.SizeMax, &task.SizeMed, &task.SizeAvg, &task.PriceMin, &task.PriceMax, &task.PriceMed, &task.PriceAvg)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func DeleteTask() (int64, error) {
	result, err := DB.Exec("DELETE FROM tasks;")
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
