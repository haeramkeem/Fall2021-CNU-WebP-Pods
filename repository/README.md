# PODs DB Settings

## Docker settings
### Inital settings
1. Image: [Postgres docker official image, latest](https://hub.docker.com/_/postgres)   
```
docker pull postgres
```
2. Create volume to sync with host
```
docker volume create pgdata
```
### Run container
1. Run
```
docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=1q2w3e4r -d -v pgdata:/var/lib/postgresql/data postgres
```
2. Entering container
```
docker exec -it postgres /bin/bash
```

## DB & User setting
1. Entering Postgres Shell
```
psql -U postgres
```
2. Create superuser
```
CREATE USER hrk PASSWORD '1q2w3e4r' SUPERUSER;
```
3. Create DB
```
CREATE DATABASE pods OWNER hrk;
```

## Table setting
1. Entering DB Shell
```
\c pods hrk
```
2. Create Table
```
CREATE TABLE problem (
id integer NOT NULL,
date integer NOT NULL,
category character varying(32),
title character varying(1000),
link character varying(2000),
description character varying(10485760),
CONSTRAINT problem_pk PRIMARY KEY (id)
);
```
