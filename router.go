package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func addRules(context *gin.Context) {
	token := context.GetString("session-token")
	u, err := verifyUser(token)
	if err != nil {
		context.String(http.StatusBadRequest, "failed to verify user from session token: %v", err)
		return
	}
	if u.Role != AdminRole {
		context.String(http.StatusUnauthorized, "user %s is not an admin", u.Name)
		return
	}

	file, err := context.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, "failed to form file: %v", err)
		return
	}
	if err := context.SaveUploadedFile(file, customLintingRulesFile); err != nil {
		context.String(http.StatusBadRequest, "failed to save uploaded file: %v", err)
		return
	}
	bs, err := os.ReadFile(customLintingRulesFile)
	if err != nil {
		context.String(http.StatusInternalServerError, "failed to read file: %v", err)
		return
	}

	projects, err := ListAccessibleProjects(u)
	if err != nil {
		context.String(http.StatusBadRequest, "failed to list accessible project for user: %v", err)
		return
	}
	projectID := context.GetInt("project-id")
	var project Project
	var found bool
	for _, p := range projects {
		if p.ID == projectID {
			project = p
			found = true
			break
		}
	}
	if !found {
		context.String(http.StatusUnauthorized, "cannot access project with ID %d", projectID)
		return
	}

	if len(project.CustomLintingRules) != 0 {
		context.String(http.StatusForbidden, "cannot apply multiple rules to one project with ID %d", projectID)
		return
	}

	project.CustomLintingRules = bs // TODO: save it to the server

	context.String(http.StatusOK, "add custom linting rules to project %s successfully", project.Name)
}

func updateRules(context *gin.Context) {
	token := context.GetString("session-token")
	u, err := verifyUser(token)
	if err != nil {
		context.String(http.StatusBadRequest, "failed to verify user from session token: %v", err)
		return
	}
	if u.Role != AdminRole {
		context.String(http.StatusUnauthorized, "user %s is not an admin", u.Name)
		return
	}

	file, err := context.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, "failed to form file: %v", err)
		return
	}
	if err := context.SaveUploadedFile(file, customLintingRulesFile); err != nil {
		context.String(http.StatusBadRequest, "failed to save uploaded file: %v", err)
		return
	}
	bs, err := os.ReadFile(customLintingRulesFile)
	if err != nil {
		context.String(http.StatusInternalServerError, "failed to read file: %v", err)
		return
	}

	projects, err := ListAccessibleProjects(u)
	if err != nil {
		context.String(http.StatusBadRequest, "failed to list accessible project for user: %v", err)
		return
	}
	projectID := context.GetInt("project-id")
	var project Project
	var found bool
	for _, p := range projects {
		if p.ID == projectID {
			project = p
			found = true
			break
		}
	}
	if !found {
		context.String(http.StatusUnauthorized, "cannot access project with ID %d", projectID)
		return
	}

	if len(project.CustomLintingRules) == 0 {
		context.String(http.StatusForbidden, "no rules applied to one project with ID %d", projectID)
		return
	}

	project.CustomLintingRules = bs // TODO: save it to the server

	context.String(http.StatusOK, "update custom linting rules to project %s successfully", project.Name)
}

func getRules(context *gin.Context) {
	token := context.GetString("session-token")
	u, err := verifyUser(token)
	if err != nil {
		context.String(http.StatusBadRequest, "failed to verify user from session token: %v", err)
		return
	}

	projects, err := ListAccessibleProjects(u)
	if err != nil {
		context.String(http.StatusBadRequest, "failed to list accessible project for user: %v", err)
		return
	}

	projectID := context.GetInt("project-id")
	for _, p := range projects {
		if p.ID == projectID {
			context.JSON(http.StatusOK, gin.H{
				"custom-linting-rules": string(p.CustomLintingRules),
			})
			return
		}
	}

	context.String(http.StatusUnauthorized, "cannot access project whose ID is %d", projectID)
}
