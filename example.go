package main

var (
	kongOrg = Organization{
		ID:   1,
		Name: "Kong Organization",
		Users: []User{
			amyWang,
			adminUser,
		},
		Projects: []Project{
			uuidGeneratorPrj,
		},
	}
	uuidGeneratorPrj = Project{
		ID:   1,
		Name: "UUID Generator based on httpbin.org 1.3",
		Users: []User{
			amyWang,
			adminUser,
		},
	}
	adminUser = User{
		UUID: 0,
		Name: "I am your father",
		Role: AdminRole,
	}
	amyWang = User{
		UUID: 1,
		Name: "Amy Wang",
		Role: NonAdminRole,
	}
)
