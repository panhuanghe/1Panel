package dto

import "time"

type PostgresqlDBSearch struct {
	PageInfo
	Info     string `json:"info"`
	Database string `json:"database" validate:"required"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
}

type PostgresqlDBInfo struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	Name           string    `json:"name"`
	From           string    `json:"from"`
	PostgresqlName string    `json:"postgresqlName"`
	Format         string    `json:"format"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Permission     string    `json:"permission"`
	BackupCount    int       `json:"backupCount"`
	Description    string    `json:"description"`
}

type PostgresqlOption struct {
	ID       uint   `json:"id"`
	From     string `json:"from"`
	Type     string `json:"type"`
	Database string `json:"database"`
	Name     string `json:"name"`
}

type PostgresqlDBCreate struct {
	Name        string `json:"name" validate:"required"`
	From        string `json:"from" validate:"required,oneof=local remote"`
	Database    string `json:"database" validate:"required"`
	Format      string `json:"format"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Permission  string `json:"permission" validate:"required"`
	Description string `json:"description"`
}

type PostgresqlLoadDB struct {
	From     string `json:"from" validate:"required,oneof=local remote"`
	Type     string `json:"type" validate:"required,oneof=postgresql"`
	Database string `json:"database" validate:"required"`
}

type PostgresqlDBDeleteCheck struct {
	ID       uint   `json:"id" validate:"required"`
	Type     string `json:"type" validate:"required,oneof=postgresql"`
	Database string `json:"database" validate:"required"`
}

type PostgresqlDBDelete struct {
	ID           uint   `json:"id" validate:"required"`
	Type         string `json:"type" validate:"required,oneof=postgresql"`
	Database     string `json:"database" validate:"required"`
	ForceDelete  bool   `json:"forceDelete"`
	DeleteBackup bool   `json:"deleteBackup"`
}

type PostgresqlConfUpdateByFile struct {
	Type     string `json:"type" validate:"required,oneof=postgresql mariadb"`
	Database string `json:"database" validate:"required"`
	File     string `json:"file"`
}
