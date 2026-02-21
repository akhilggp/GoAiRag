Check Go / environment
go version
go env GOPATH GOROOT
Create a new project (module)
mkdir myapp && cd myapp
go mod init myapp
Run / Build
go run . # run main package in current folder
go run ./cmd/api # run if main.go is under cmd/api

go build . # build current package
go build ./... # build all packages in module
go build -o bin/app ./cmd/api # build binary with output name
Test
go test ./...
go test -v ./...
go test ./... -run TestName
Dependencies (Go Modules)
go get <module>@latest # add/upgrade a dependency
go mod tidy # sync go.mod/go.sum with imports (add missing, remove unused)
go mod download # download all deps to local cache
go mod verify # verify deps checksums
go list -m all # list all modules (deps)
go mod why -m <module> # explain why a module is included
Format
gofmt -w . # format Go files (standard)
goimports (format + fix imports)

Install:

go install golang.org/x/tools/cmd/goimports@latest
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

Use:

goimports -w .

Docker Compose essentials
Start services
docker compose up -d
See status
docker compose ps
docker ps
Logs
docker compose logs -f
docker compose logs -f <service_name>
Stop without deleting (start later)
docker compose stop
docker compose start
Restart
docker compose restart
docker compose restart <service_name>
Remove containers (keeps named volumes by default)
docker compose down
Remove containers + delete volumes (data loss)
docker compose down -v
Pull/refresh images
docker compose pull
docker compose up -d --pull always
Clean unused stuff (optional)
docker system prune
docker volume ls
docker volume rm <volume_name>
