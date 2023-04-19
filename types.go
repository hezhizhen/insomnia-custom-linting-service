package main

type Organization struct {
	ID       int
	Name     string
	Users    []User // users in this organization
	Projects []Project
}

type Project struct {
	ID                 int
	Name               string
	Users              []User // users can access this project
	CustomLintingRules []byte
}

type User struct {
	UUID         int
	Name         string
	Role         Role
	Organization *Organization
}

type Role int

const (
	AdminRole Role = iota
	NonAdminRole
)
