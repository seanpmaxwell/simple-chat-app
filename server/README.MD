## About
Simple CRUD app with Login


#### Architecure
Layered: Routers (http) -> Services (business logic) -> Repos (Gorm queries) 


#### Tools
Go + Gin (web-server and hot-reloading) + Gorm (ORM) + Postgres (DB) + Godotenv (environment variable loading) + Wire (Depedency Injection) + Golang-JWT (Session Storage) + Bcrypt (password hashing)


#### Setup
- Created with: `go mod init server`
- For hot-reloading:
    - Add this to `export PATH=/home/your_name/go/bin:$PATH` to `.bashrc`.
    - Hot reloading docs: `https://github.com/codegangsta/gin`
- Add this to Vscode's `settings.json` file:
```JSON
    // so `wire.go` does not show linter warnings:
    "gopls.env": {
        "GOFLAGS": "-tags=wireinject"
    },
    // To turn off the linter when saving (get's annoying).
    "[go]": {
        "editor.formatOnSave": false
    }
```


#### Commands
- Start all commands with `sh run.sh` (i.e. `sh run.sh start`):
    - `start`: Run as is.
    - `hot`: Run with hot-reloading.
    - `build`: Compile.
    - `startb`: Run compiled code.
    - `full`: Build and run compiled code.
    - `clean`: Clean up go.mod
    - `wire`: Update dependency injections.
- Append the environment name at the end (i.e. `sh run.sh start staging`):
    - If you don't sepecify an environment, `development` is default.
    - Environment names must match file names under `./env`


#### Notes
- If you create a dependency (i.e. struct that's a router, repo, service, util, etc), you need to place it's dependencies (if it has any) as properties, create a `NewNameOfTheStruct"()` function for it, and place that function in `wire.go`. Then update the dependency injections with `sh run.sh wire`.
- Constraints cannot be altered after a table is created. You need to drop the table and recreate the constraint.
- For programs use structs w/ dependency injection to organize code, for libraries, use packages to organize code.
