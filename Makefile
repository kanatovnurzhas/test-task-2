compose:
	docker-compose up -d
topic:
	docker exec -it kafka_container kafka-topics.sh --create --topic my-topic --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092
