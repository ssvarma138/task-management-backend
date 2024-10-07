# Task Management Backend

Hey there! 👋 This is the backend part of my task management app for my college project. It's built with Go and uses the Gin framework.

## What it does

- Handles all the behind-the-scenes stuff for our task app
- Stores and manages tasks in a SQLite database
- Provides API endpoints so the frontend can talk to it

## How to get it running

1. Clone this repo to your computer
2. Make sure you've got Go installed (version 1.16 or later should work)
3. Open a terminal, go to the project folder, and run:  go mod tidy
This will get all the packages we need
4. To start the server, run: go run cmd/server/main.go
5. The server will run on port 8080, so you can access the API at http://localhost:8080

## API Stuff

Here's what you can do:
- GET /api/tasks: Get all tasks
- POST /api/tasks: Add a new task
- PUT /api/tasks/:id: Update a task
- DELETE /api/tasks/:id: Delete a task

## What I want to add later

- User login (so everyone can have their own tasks)
- Maybe add categories for tasks
- Make it look fancier in the terminal with some colors

## Notes

This is just the backend part. You'll need the frontend too to see anything in the browser.