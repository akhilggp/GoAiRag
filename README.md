# Go + Docker Compose Cheat Sheet

A quick reference for the Go commands and Docker Compose commands you’ll use most often.

---

## Go

### Check Go / Environment

```bash
go version
go env GOPATH GOROOT
```

### Create a New Project (Module)

```bash
mkdir myapp && cd myapp
go mod init myapp
```

### Run / Build

```bash
go run .          # Run main package in current folder
go run ./cmd/api  # Run if main.go is under cmd/api
```

```bash
go build .                    # Build current package
go build ./...                # Build all packages in module
go build -o bin/app ./cmd/api # Build binary with output name
```

### Test

```bash
go test ./...
go test -v ./...
go test ./... -run TestName
```

### Dependencies (Go Modules)

```bash
go get <module>@latest   # Add/upgrade a dependency
go mod tidy              # Sync go.mod/go.sum with imports (add missing, remove unused)
go mod download          # Download all deps to local cache
go mod verify            # Verify deps checksums
go list -m all           # List all modules (deps)
go mod why -m <module>   # Explain why a module is included
```

### Format

```bash
gofmt -w .  # Format Go files (standard)
```

### goimports (Format + Fix Imports)

#### Install

```bash
go install golang.org/x/tools/cmd/goimports@latest
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

#### Use

```bash
goimports -w .
```

---

## Docker Compose

### Start Services

```bash
docker compose up -d
```

### See Status

```bash
docker compose ps
docker ps
```

### Logs

```bash
docker compose logs -f
docker compose logs -f <service_name>
```

### Stop Without Deleting (Start Later)

```bash
docker compose stop
docker compose start
```

### Restart

```bash
docker compose restart
docker compose restart <service_name>
```

### Remove Containers (Keeps Named Volumes by Default)

```bash
docker compose down
```

### Remove Containers + Delete Volumes (Data Loss)

```bash
docker compose down -v
```

### Pull / Refresh Images

```bash
docker compose pull
docker compose up -d --pull always
```

### Clean Unused Stuff (Optional)

```bash
docker system prune
docker volume ls
docker volume rm <volume_name>
```
