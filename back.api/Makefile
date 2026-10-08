.SUFFIXES:

CMD_PATH := ./cmd
LOCAL_COMPOSE_FILE := docker-compose.local.yaml
COMPOSE_FILE := docker-compose.yaml
DB_CONTAINER := dasom-postgres
MIGRATION_DUMP := migrated_postgres_dump.sql


## test 관련 #############################################################################

# mock 생성
mock-gen:
	@echo "Installing counterfeiter for mock generation..."
	go install github.com/maxbrunsfeld/counterfeiter/v6@v6.12.2

	@echo "Generating mocks with counterfeiter..."
	go generate ./...

# 테스트 실행 및 커버리지 리포트 생성
test:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out




## docs 관련 #############################################################################
# Swagger 문서 생성
swag:
	@echo "Generating Swagger docs..."
	swag init -g $(CMD_PATH)/main.go -o ./docs
	

## Local test 관련 #######################################################################
# local 테스트용 run 명령 (.env.local 파일에서 환경변수 로드)
local-run: swag
	@if [ -f .env.local ]; then \
		echo "Loading .env.local and starting server..."; \
		export $$(grep -v '^#' .env.local | xargs) && go run $(CMD_PATH); \
	else \
		echo ".env.local file not found!"; \
		exit 1; \
	fi

# Local DB만 실행
local-db-up:
	@echo "Starting local database..."
	docker compose -f $(LOCAL_COMPOSE_FILE) up -d postgres


# Local DB만 종료
local-db-down:
	@echo "Stopping local database..."
	docker compose -f $(LOCAL_COMPOSE_FILE) stop postgres

# Local DB + volume 삭제 (전체 스택 down -v)
local-db-reset:
	@echo "Resetting local database..."
	docker compose -f $(LOCAL_COMPOSE_FILE) down -v

# DB 로그 확인
local-db-logs:
	docker compose -f $(LOCAL_COMPOSE_FILE) logs -f

# Local DB Migration 실행
local-db-migrate:
	@echo "Running local database migrations..."
	docker exec -i dasom-postgres-local psql -U testuser -d dasomweb -c "TRUNCATE TABLE users, posts, comments, files, images RESTART IDENTITY CASCADE;"
	docker exec -i dasom-postgres-local psql -U testuser -d dasomweb < migrated_postgres_dump.sql

	################# 배포 관련 #############################################################################
# 배포용 run 명령 (환경변수 로드)

# DB 실행
db-up:
	@echo "Starting database..."
	docker compose -f $(COMPOSE_FILE) up -d

# DB 종료
db-down:
	@echo "Stopping database..."
	docker compose -f $(COMPOSE_FILE) down

# DB + volume 삭제
db-reset:
	@echo "Resetting database..."
	docker compose -f $(COMPOSE_FILE) down -v

# DB 로그 확인
db-logs:
	docker compose -f $(COMPOSE_FILE) logs -f

# 레거시 데이터 이관 덤프 적용 (GORM AutoMigrate로 스키마가 먼저 만들어져 있어야 함)
db-migrate:
	@if [ ! -f .env ]; then \
		echo ".env file not found!"; \
		exit 1; \
	fi
	@if [ ! -f $(MIGRATION_DUMP) ]; then \
		echo "$(MIGRATION_DUMP) not found!"; \
		exit 1; \
	fi
	@export $$(grep -v '^#' .env | xargs) && \
	echo "Applying $(MIGRATION_DUMP) to $$DB_NAME..." && \
	docker exec -i $(DB_CONTAINER) psql -U $$DB_USER -d $$DB_NAME < $(MIGRATION_DUMP)

# 서버 실행
run: swag
	@if [ -f .env ]; then \
		echo "Loading .env and starting server..."; \
		export $$(grep -v '^#' .env | xargs) && go run $(CMD_PATH); \
	else \
		echo ".env file not found!"; \
		exit 1; \
	fi

# 외부 의존성 패키지 설치
deps:
	@echo "Installing dependencies..."
	go mod tidy

	@echo "Installing Gin framework..."
	go get github.com/gin-gonic/gin

	@echo "Installing database packages..."
	get gorm.io/gorm
	go get github.com/lib/pq
	go get gorm.io/driver/postgres
	
	@echo "Installing JWT package..."
	go get github.com/golang-jwt/jwt/v5

	@echo "Installing Config related packages..."
	go get github.com/go-viper/mapstructure/v2
	go get github.com/joho/godotenv
	go get github.com/spf13/viper