#!/bin/bash
 
if [[ $EUID -ne 0 ]]
then
  echo "You must be root to run this."
  exit 1
fi
 
echo "removing ./db-data"
rm -rf db-data
echo "./db-data removed, done"
