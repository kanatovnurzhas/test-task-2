compose:
	docker-compose up -d
topic:
	docker exec -it kafka_container kafka-topics.sh --create --topic stud-to-course --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092; \
	docker exec -it kafka_container kafka-topics.sh --create --topic course-to-stud --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092; \
	docker exec -it kafka_container kafka-topics.sh --create --topic answer-for-stud --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092; \
	docker exec -it kafka_container kafka-topics.sh --create --topic answer-for-course --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092



