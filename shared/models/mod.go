package models

type Mod struct {
	Name string
	Source Service
	Version string
	DownloadUrl string
	Hash string
	ProjectCode string
	VersionCode string
}