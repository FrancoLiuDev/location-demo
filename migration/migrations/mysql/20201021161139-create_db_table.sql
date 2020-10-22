
-- +migrate Up
CREATE DATABASE MAIN;
USE MAIN;

CREATE TABLE LOCATION ( id VARCHAR(30), location json, dateAdd datetime);




-- +migrate Down
DROP SCHEMA IF EXISTS MAIN;