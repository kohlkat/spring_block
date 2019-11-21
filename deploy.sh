#!/usr/bin/env bash

go build
./spring_block --verb True &
cd frontend
npm start &
cd ..

