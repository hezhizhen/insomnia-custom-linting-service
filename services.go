package main

// existing external services

// an authentication service to verify the identity and role of a user via session token
func verifyUser(token string) (u User, err error) {
	return User{}, nil
}

// a service to retrieve a list of project ids that an authenticated user has access to.
func listProjectsForUser(u User) (ids []int, err error) {
	return nil, nil
}

func ListAccessibleProjects(u User) (projects []Project, err error) {
	// a service to retrieve a list of project ids that an authenticated user has access to
	ids, err := listProjectsForUser(u)
	if err != nil {
		return nil, err
	}

	m := convertIntSlicesToMap(ids)
	for _, p := range u.Organization.Projects {
		if !m[p.ID] {
			continue
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func convertIntSlicesToMap(nums []int) map[int]bool {
	ret := make(map[int]bool, len(nums))
	for _, num := range nums {
		ret[num] = true
	}
	return ret
}
