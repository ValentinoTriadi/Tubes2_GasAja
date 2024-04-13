#!/bin/bash

cd frontend 
npm i
cd ../
docker compose build
docker network create local