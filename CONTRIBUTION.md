
# Contribution

- [Folder Structure](#folder-structure)
  * [cmd](#cmd)
  * [internal](#internal)
  * [pkg](#pkg)
  * [config](#config)
  * [scripts](#scripts)
- [How to add new endpoint](#how-to-add-new-endpoint)
  * [Example](#example)
- [Contribution guidelines](#contribution-guidelines)

## Folder Structure

We'll use this [project structure](https://dev.to/jinxankit/go-project-structure-and-guidelines-4ccm)

```
online-judge-backend/
├── cmd
├── internal
├── pkg
├── config
└── scripts
```


### cmd
This folder contains the main application entry point files for the project

```
cmd/
└── main.go
```

### internal
This package holds the private library code used in your service, it is specific to the function of the service and not shared with other services. Most of our code will be in this folder.

Example of code that should be in this folder:

```
internal/
├── db
│   └── db_engine.go
├── routes
│   └── routes.go
└── server
    └── server.go
```



### pkg
This folder contains code which is OK for other services to consume, this may include API clients, or utility functions which may be handy for other projects but don’t justify their own project.

Example of code that should be in this folder (Requests and responses):
```
pkg/
├── requests
│   └── submission.go
└── responses
    └── submission.go
```
### config
This folder should'nt contain go code, it contains configuration files and environment variables.

Example:

```
config/
├── config.json
└── .env-example
```

### scripts
This folder should'nt contain go code, it contains any utility scripts or codes (dockefiles, .sh files)
```
scripts/
└── database
    └── populate.sh
```

## How to add new endpoint

1. Create file that contains handler function.
   if your method is post the file should be in `internal/handlers/post`
   else if your method is get the file should be in `internal/handlers/get`

2. Add path to end point, handler function, method in `LoadRoutes` function in `routes/routes.go`

   ```go
   router.HandleFunc("<path-to-endpoint>", <handler-function>).Methods("<method>")
   ```

3. Create request (if needed) in `pkg/requests`

4. Create response (if needed) in  `pkg/responses`

### Example

Assume we want to create new endpoint to submit problem with **POST** method.

1. Create `submit.go` file to contain our handler function in `internal/handlers/post`

   ```go
   package post
   
   // handler function
   func Submit(w http.ResponseWriter, r *http.Request) {
       
   }
   ```

2. Add path to submit endpoint, handler function, method in `LoadRoutes` function in `routes/routes.go`

   ```go
   package routes
   
   import (
   	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/get"
   	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/post"
   	"github.com/gorilla/mux"
   )
   func LoadRoutes()(* mux.Router){
   	router:=mux.NewRouter()
   	router.HandleFunc("/", get.Root).Methods("GET")	
   	
       /*
       	path to submit endpoint: "/submit/{problemID}"
       	handler function: post.Submit
       	method: POST
       */
   	router.HandleFunc("/submit/{problemID}", post.Submit).Methods("POST") // this is the added line
   	
   	return router
   }
   ```

3. Create `submit.go` in `pkg/requests`

   ```go
   package requests
   
   type SubmissionRequest struct {
   	OwnerID  int `json:"ownerID,omitempty"`
   	Language string `json:"language,omitempty"`
   	Code string `json:"code,omitempty"`
   }
   
   ```

4. Create `submit.go` in `pkg/responses`

   ```go
   package responses
   
   type SubmissionResponse struct {
   	Verdict string `json:"verdict,omitempty"`
   }
   ```

## Contribution guidelines

- Branch name of each pull request should be `OJ-#` where `#` corresponds to story id on jira, as example if I'm working on story with id = 1, branch name should be `OJ-1`.

- Each feature story on jira should have one commit corresponding to it. 

  If you have already made multiple commits before pushing your branch, squash the commits to one commit  then push, see [Squash](https://onlinejudge.atlassian.net/wiki/spaces/OJ/pages/33060/Squash).

- First line of commit message is the title of the commit, it should describe the goal of the commit

- The commit message should also include important modifications made in the commit

- Verbs in the commit message should be in imperative form, example: Add, Implement, Remove ..

  Example commit message of creating user profile page

  ```
  Implement User Profile Page
  
  * Create Dummy User data
  * Create Card Component
  ```

