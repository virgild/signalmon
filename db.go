package main

import (
	"database/sql"
	"time"
)

func initdb(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	if db == nil {
		panic("Cannot init db")
	}

	sql := `
    CREATE TABLE IF NOT EXISTS readings (
      id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
      created_at integer(128),
      forward_freq1 float(128),
      forward_freq2 float(128),
      forward_freq3 float(128),
      forward_freq4 float(128),
      forward_freq5 float(128),
      forward_freq6 float(128),
      forward_freq7 float(128),
      forward_freq8 float(128),
      forward_power1 float(128),
      forward_power2 float(128),
      forward_power3 float(128),
      forward_power4 float(128),
      forward_power5 float(128),
      forward_power6 float(128),
      forward_power7 float(128),
      forward_power8 float(128),
      forward_snr1 float(128),
      forward_snr2 float(128),
      forward_snr3 float(128),
      forward_snr4 float(128),
      forward_snr5 float(128),
      forward_snr6 float(128),
      forward_snr7 float(128),
      forward_snr8 float(128),
      forward_ber1 float(128),
      forward_ber2 float(128),
      forward_ber3 float(128),
      forward_ber4 float(128),
      forward_ber5 float(128),
      forward_ber6 float(128),
      forward_ber7 float(128),
      forward_ber8 float(128),
      forward_mod1 float(128),
      forward_mod2 float(128),
      forward_mod3 float(128),
      forward_mod4 float(128),
      forward_mod5 float(128),
      forward_mod6 float(128),
      forward_mod7 float(128),
      forward_mod8 float(128),
      back_freq1 float(128),
      back_freq2 float(128),
      back_freq3 float(128),
      back_freq4 float(128),
      back_power1 float(128),
      back_power2 float(128),
      back_power3 float(128),
      back_power4 float(128),
      back_mod1 float(128),
      back_mod2 float(128),
      back_mod3 float(128),
      back_mod4 float(128)
    );
  `

	_, err = db.Exec(sql)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insertData(db *sql.DB, statsdata StatsData) error {
	sql := `
    INSERT INTO readings(
      created_at,
      forward_freq1, forward_freq2, forward_freq3, forward_freq4, forward_freq5, forward_freq6, forward_freq7, forward_freq8,
      forward_power1, forward_power2, forward_power3, forward_power4, forward_power5, forward_power6, forward_power7, forward_power8,
      forward_snr1, forward_snr2, forward_snr3, forward_snr4, forward_snr5, forward_snr6, forward_snr7, forward_snr8,
      forward_ber1, forward_ber2, forward_ber3, forward_ber4, forward_ber5, forward_ber6, forward_ber7, forward_ber8,
      forward_mod1, forward_mod2, forward_mod3, forward_mod4, forward_mod5, forward_mod6, forward_mod7, forward_mod8,
      back_freq1, back_freq2, back_freq3, back_freq4,
      back_power1, back_power2, back_power3, back_power4,
      back_mod1, back_mod2, back_mod3, back_mod4
    )
     VALUES(
      ?,
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
      ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
      ?, ?, ?, ?,
      ?, ?, ?, ?,
      ?, ?, ?, ?
    );
  `
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		time.Now().UTC(),
		statsdata.ForwardSignals[0].Frequency,
		statsdata.ForwardSignals[1].Frequency,
		statsdata.ForwardSignals[2].Frequency,
		statsdata.ForwardSignals[3].Frequency,
		statsdata.ForwardSignals[4].Frequency,
		statsdata.ForwardSignals[5].Frequency,
		statsdata.ForwardSignals[6].Frequency,
		statsdata.ForwardSignals[7].Frequency,
		statsdata.ForwardSignals[0].Power,
		statsdata.ForwardSignals[1].Power,
		statsdata.ForwardSignals[2].Power,
		statsdata.ForwardSignals[3].Power,
		statsdata.ForwardSignals[4].Power,
		statsdata.ForwardSignals[5].Power,
		statsdata.ForwardSignals[6].Power,
		statsdata.ForwardSignals[7].Power,
		statsdata.ForwardSignals[0].SNR,
		statsdata.ForwardSignals[1].SNR,
		statsdata.ForwardSignals[2].SNR,
		statsdata.ForwardSignals[3].SNR,
		statsdata.ForwardSignals[4].SNR,
		statsdata.ForwardSignals[5].SNR,
		statsdata.ForwardSignals[6].SNR,
		statsdata.ForwardSignals[7].SNR,
		statsdata.ForwardSignals[0].BER,
		statsdata.ForwardSignals[1].BER,
		statsdata.ForwardSignals[2].BER,
		statsdata.ForwardSignals[3].BER,
		statsdata.ForwardSignals[4].BER,
		statsdata.ForwardSignals[5].BER,
		statsdata.ForwardSignals[6].BER,
		statsdata.ForwardSignals[7].BER,
		statsdata.ForwardSignals[0].Modulation,
		statsdata.ForwardSignals[1].Modulation,
		statsdata.ForwardSignals[2].Modulation,
		statsdata.ForwardSignals[3].Modulation,
		statsdata.ForwardSignals[4].Modulation,
		statsdata.ForwardSignals[5].Modulation,
		statsdata.ForwardSignals[6].Modulation,
		statsdata.ForwardSignals[7].Modulation,
		statsdata.ReturnSignals[0].Frequency,
		statsdata.ReturnSignals[1].Frequency,
		statsdata.ReturnSignals[2].Frequency,
		statsdata.ReturnSignals[3].Frequency,
		statsdata.ReturnSignals[0].Power,
		statsdata.ReturnSignals[1].Power,
		statsdata.ReturnSignals[2].Power,
		statsdata.ReturnSignals[3].Power,
		statsdata.ReturnSignals[0].Modulation,
		statsdata.ReturnSignals[1].Modulation,
		statsdata.ReturnSignals[2].Modulation,
		statsdata.ReturnSignals[3].Modulation,
	)
	if err != nil {
		return err
	}

	return nil
}
