#!/bin/sh

case "$1" in
  "start"|"up")
    docker-compose up -d;;
  "stop"|"down")
    docker-compose down;;
  "restart")
    docker-compose down
    docker-compose up -d
   ;;
  "ps")
    docker-compose ps;;
  "logs")
    docker-compose logs;;
#  "exec")
#    docker-compose exec server $2:@;;
#  "bash")
#    docker-compose exec server bash;;
  "default")
    echo "Hello World!";;
esac 
