# streamprocessing
streamprocessing golang architecture 



go consumer with kafka endpoint 
go producer  with kafka endpoint 
kafka   docker  without zookeeper 
processor 


analytics  + websocket  + 1 webpage loded html 

create topic 

docker exec -it kafka kafka-topics --create --topic streamprocessing --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1


consumde from confluentinc
docker exec -it $(docker ps --filter "ancestor=confluentinc/cp-kafka" -q) \
  kafka-console-consumer --topic streamprocessing --from-beginning --bootstrap-server localhost:9092
