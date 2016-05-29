package main

import (
	"database/sql"
	"fmt"
	"time"
)

var init_sql string = `
    CREATE TABLE IF NOT EXISTS readings (
      id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
      reading_date integer(128) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS channel_readings (
      id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
      reading_id integer(128) NOT NULL,
      type integer(128) NOT NULL,
      channel_number integer(128) NOT NULL,
      frequency float(128) NOT NULL,
      power float(128) NOT NULL,
      snr float(128),
      ber float(128),
      modulation integer(128) NOT NULL
    );

    CREATE INDEX IF NOT EXISTS idx_reading ON channel_readings (reading_id ASC);
  `

var fsignal_sql string = `
    INSERT INTO channel_readings (
      reading_id,
      type,
      channel_number,
      frequency,
      power,
      snr,
      ber,
      modulation
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?);
  `
var rsignal_sql string = `
    INSERT INTO channel_readings (
      reading_id,
      type,
      channel_number,
      frequency,
      power,
      modulation
    ) VALUES (?, ?, ?, ?, ?, ?);
  `

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	if db == nil {
		panic(fmt.Errorf("No database after opening"))
	}

	_, err = db.Exec(init_sql)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertData(db *sql.DB, signalsdata *SignalsData) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	res1, err := tx.Exec("INSERT INTO readings (reading_date) VALUES (?)", time.Now())
	if err != nil {
		return err
	}

	reading_id, err := res1.LastInsertId()
	if err != nil {
		return err
	}

	fsignal_stmt, err := tx.Prepare(fsignal_sql)
	if err != nil {
		panic(err)
	}

	rsignal_stmt, err := tx.Prepare(rsignal_sql)
	if err != nil {
		panic(err)
	}

	for _, fsignal := range signalsdata.ForwardSignals {
		_, err := fsignal_stmt.Exec(
			reading_id,
			1,
			fsignal.Number(),
			fsignal.Frequency(),
			fsignal.Power(),
			fsignal.SNR(),
			fsignal.BER(),
			fsignal.Modulation(),
		)
		if err != nil {
			panic(err)
		}
	}

	for _, rsignal := range signalsdata.ReturnSignals {
		_, err := rsignal_stmt.Exec(
			reading_id,
			2,
			rsignal.Number(),
			rsignal.Frequency(),
			rsignal.Power(),
			rsignal.Modulation(),
		)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	return err
}
