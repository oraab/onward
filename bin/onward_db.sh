#! /bin/bash 

sqlite3 $1 'CREATE TABLE LISTS(ID int64, NAME varchar, TYPE varchar)';
sqlite3 $1 'CREATE TABLE ITEMS(ID int64, LIST int64, ORDER int, STATUS varchar, DONE_FOR_TODAY bool)';

