#! /bin/bash

if [[ -z ${ONWARD_BASE_PATH} ]]
then 
   echo "Error: ONWARD_BASE_PATH not defined."
   exit 1
fi 

if [[ -f $ONWARD_BASE_PATH/db/onward.db ]]
then
    echo "Error: Onward DB already exists."
else 
    sqlite3 $ONWARD_BASE_PATH/db/onward.db "create table tasks(id int, description string);"
fi
