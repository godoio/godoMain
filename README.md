# Go do
> A simple to-do app.
In your coding adventures you have seen many to-do apps as basic as displaying the hierarchy of what all you will do in your day. You might have had interacted with the to-do app too. This is also the same.

In the Go do you get to play with your tasks in four ways, create, edit, view and delete, also you save it in your database whichever you like to have. That's it.

## Technologies used and prerequisites
1. Go version go1.17.8 Get it here https://go.dev/dl/
2. Node.js v16.14.2 https://nodejs.org/en/download/
3. MySQL 8.0 Get it here https://dev.mysql.com/downloads/mysql/8.0.html

## Usage
After you have installed your prerequisites, follow the steps:
1. Create the database and table in MySQL
> create database taskdatabase;
> use taskdatabase;
> create table tasks(task_id int PRIMARY_KEY, task_title varchar(255), task_description varchar(255), task_date varchar(255));

2. To run the Go API:
Replace the User: "root", and Passwd: "root1234", in <b>databaseServiceImpl.go</b> to your preferred username and password in MySQL that you have set up.
> cd godoMain
> go run main.go

3. To run the UI:
> cd godoMain/ui
> npm install Vue@latest
> npm install
> npm run dev
