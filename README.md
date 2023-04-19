# insomnia-custom-linting-service

1. Use Gin to create an HTTP server
2. Provide 3 APIs:
   1. `/add-custom-linting-rules`: allow admin users to upload a YAML file and add its content to a project
   2. `/update-custom-linting-rules`: allow admin users to upload a YAML file and use its content to update custom linting rules of a project
   3. `/get-custom-linting-rules`: return the custom linting rules of a project as a string to users who can access this project
3. Save linting rules as a YAML file locally

## TODOs

1. Need a service to get organization/project/user details from their own ids
2. Need a service to save organization/project/user details
