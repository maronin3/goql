#!/bin/bash
 
if [[ $EUID -ne 0 ]]
then
  echo "You must be root to run this."
  exit 1
fi
 
docker-compose down --remove-orphans &&
./remove_dbdata.sh &&
./run_compose.sh