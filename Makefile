consumer:
	@echo "===== RUNNING CONSUMER SERVICES ====="
	go run src/consumer/consumer.go

producer:
	@echo "===== RUNNING PRODUCER SERVICES ====="
	go run src/producer/producer.go
