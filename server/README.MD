#### Setup
- Created with: `go mod init server`
- For hot-reloading:
    - Add this to `export PATH=/home/your_name/go/bin:$PATH` to `.bashrc`.
    - Hot reloading docs: `https://github.com/codegangsta/gin`


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
- If you add a dependency struct (i.e. router, repo, service, util, etc), you need to place it's dependencies (if it has any) as properties, create a `NewNameOfTheStruct"()` function for it, and place that function in `wire.go`. Then update the dependency injections with `sh run.sh wire`.
- Constraints cannot be altered after a table is created. You need to drop the table and recreate the constraint.


#### Todo
