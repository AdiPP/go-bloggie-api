module-new:
	touch domain/$$name.go && mkdir -p $$name/entities && mkdir -p $$name/delivery/http && mkdir -p $$name/delivery/grpc && mkdir -p $$name/repositories && mkdir -p $$name/usecases
	
run:
	go run main.go

dev:
	air

build:
	go build -o payment-gateway-service main.go

migrate-up:
	go run cmd/command/command.go migrate up

migrate-down:
	go run cmd/command/command.go migrate down

migrate-new:
	go run cmd/command/command.go migrate make -f=$$filename

test-test:
	echo "test"