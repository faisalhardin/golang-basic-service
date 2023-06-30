consumer:
	@echo "===== RUNNING CONSUMER SERVICES ====="
	go run src/consumer/consumer.go

producer:
	@echo "===== RUNNING PRODUCER SERVICES ====="
	go run src/producer/producer.go

grpc-service:
	@echo "===== RUNNING GRPC SERVICES ====="
	go run src/grpc/main.go
