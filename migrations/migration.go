package migration

import (
	"database/sql"
	"fmt"
)

func RunMigration(db *sql.DB) error {
	var extra string

	query := `
        SELECT EXTRA
        FROM information_schema.columns
        WHERE table_schema = DATABASE()
        AND table_name = 'category'
        AND column_name = 'id'
    `

	err := db.QueryRow(query).Scan(&extra)
	if err != nil {
		return fmt.Errorf("gagal cek informasi kolom: %v", err)
	}

	if extra != "auto_increment" {
		fmt.Println("[MIGRATION] Kolom ID belum auto_increment → menjalankan ALTER TABLE...")

		_, err := db.Exec(`
            ALTER TABLE category
            MODIFY id INT NOT NULL AUTO_INCREMENT
        `)
		if err != nil {
			return fmt.Errorf("gagal ALTER TABLE: %v", err)
		}

		fmt.Println("[MIGRATION] ALTER TABLE berhasil dijalankan")
	} else {
		fmt.Println("[MIGRATION] Kolom ID sudah auto_increment → skip")
	}

	return nil
}
