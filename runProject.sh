#!/bin/bash

echo "Running project"
go run backend/main.go
cd frontend
npm install
npm run dev
cd ..
