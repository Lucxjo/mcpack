package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/lucxjo/mcpack/shared/models"
)

func GetMods(dbName string) []models.Mod {
	db, err := sql.Open("sqlite3", dbName)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt := Prepare(db, "SELECT name, source, version, download_url, hash, project_code, version_code FROM mods")

	rows, err := stmt.Query()

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var mods []models.Mod

	for rows.Next() {
		var mod models.Mod

		err := rows.Scan(&mod.Name, &mod.Source, &mod.Version, &mod.DownloadUrl, &mod.Hash, &mod.ProjectCode, &mod.VersionCode)

		if err != nil {
			panic(err)
		}

		mods = append(mods, mod)
	}

	return mods
}