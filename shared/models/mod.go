package models

type Mod struct {
	// Name of the mod.
	Name string `json:"name"`
	// Source of the mod. Currently this is always either Modrinth (0) or CurseForge (1).
	Source Service `json:"source"`
	// Version of the mod.
	Version string `json:"version"`
	// Download URL of the current version of the mod.
	DownloadUrl string `json:"download_url"`
	// Hash of the current version of the mod.
	Hash string `json:"hash"`
	// Project code of the mod for the source
	ProjectCode string `json:"project_code"`
	// Version code of the mod for the source
	VersionCode string `json:"version_code"`
	// Whether the mod is to be installed on the client, server, or both.
	Sides Side `json:"sides"`
}