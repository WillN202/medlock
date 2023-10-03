# Medlock Primary School Project
This repository contains all the project files relating to the Medlock Primary School project. 

## Contents

  * [Requirements](#Requirements)
  * [Frontend](#Frontend)
  * [Backend](#Backend)
  * [Installation](#Installation)
    * [Go](#Go)
    * [NodeJS](#NodeJS)
    * [Config the repo](#Config-the-repo)
  * [Run the Project](#Running-The-Project)
  * [Contributing & Organisation](#Contributing)

## Requirements 

An web application where teachers can set achievements for students. These achievements can then be picked up by students and progress updated.

* Teacher Achievment Creation
* View/Edit Status of a student's achievement
* Student update status of achievement
* Teacher can see all students progress
* Actor can log in with code to teacher or student account
* Teacher Editing an achievement (?)

* System can track rewards/points/whatever
* Teacher can set due date for achievement

## Frontend

This will comprise many of the parts the user will actually see. It's ran using NextJS (basically React + server-side rendering). The main folders are:

  * /Pages contains the actual pages the user will see
  * /Public contains any assets (e.g. images) used by the page

## Backend

The backend folder is where the backend API will live, coded in Go using GorrilaMux for routing. This API will be queried by the NextJS frontend and returns data formatted in JSON.

## Installation

### **Go**

**If you have installed GoLand or Go 1.17.5, you can ignore these steps.**

This can be done automatically or by using Goland.

### Using Goland (Recommended)

If you're using GoLand, it will configure the Go SDK and GoPath automatically. To do this:

1. Click on the file option, then settings
2. In the options list, click on Go, then GOPATH
3. Choose version 1.17.5 (second in the list) GoLand will automatically download and install the SDK.
4. Add the Go SDK to the ```$PATH``` by adding the following to your .profile file.

```
export PATH=$PATH:<sdk install path>/go1.17.5/bin 
```

### Manually

This is the option you should use if you're using another IDE (e.g. VSCode).

1. Visit the [Go installation page](https://go.dev/dl/go1.17.5.linux-amd64.tar.gz) and save the file in your home directory.
2. In terminal, type: 
```
tar -xf $Downloded_Go_Archive -C /usr/local
```
3. In your home directory type: 
```
sudo nano ~/.profile
``` 
4. Add the following to your .profile in your home directory: 
```
export PATH=$PATH:/usr/local/go/bin
```
5. Reload your profile by typing the following, or by closing and reopening your terminal window.
```
. ~/.profile
```
5. Test that the installation worked by typing `go version` into the terminal. You should see something like: 
```
go1.17.5 linux/amd64
```

### NodeJS

**If you have already installed NodeJS 16.13.1 you can ignore these steps.**

These steps will also install Tailwind and React.

1. [Download the NodeJS binaries](https://nodejs.org/dist/v16.13.1/node-v16.13.1-linux-x64.tar.xz) to your home directory.
2. Type the following into the console:
```
tar -xf node-v16.13.1-linux-x64.tar.xz -C /usr/local/lib/
```
3. Open .profile (`sudo nano ~/.profile`) and add the following text:
```
export PATH=/usr/local/lib/node-16.13.1-linux-x64/bin:$PATH
```
4. Reload profile and test that $PATH has been set correctly. You should expect to see the following when running the last three commands.
```
. ~/.profile
node -v
npm version
npx -v
```
```v14.16.0
{
  npm: '6.14.11',
  ares: '1.16.1',
  brotli: '1.0.9',
  cldr: '37.0',
  icu: '67.1',
  llhttp: '2.1.3',
  modules: '83',
  napi: '7',
  nghttp2: '1.41.0',
  node: '14.16.0',
  openssl: '1.1.1j',
  tz: '2020a',
  unicode: '13.0',
  uv: '1.40.0',
  v8: '8.4.371.19-node.18',
  zlib: '1.2.11'
}
8.1.2
```
Note the packages may change depending on your Node installation.

## Config the Repo

1. Clone the repo to your device (using the command line or Git tools in GoLand)

2. Run the following commands

```
cd frontend
npm install
```

## Running the Project

 In the Medlock folder, run the runProject bash script. This script will build and run both the Go API and NextJS frontend
```
./runProject.sh
```

 Open your web browser and go to ```localhost:3000```, you should see a page saying "Welcome to hell." To close the server, open the terminal and press ctrl+c. 

To save time you can create a run config within GoLand to launch the program from within the IDE.


## Contributing

> N.B. This is usually in a CONTRIBUTING.md file but is included here for ease of use. 

This project will be using a semi agile methodology. Standup meetings are held on Monday and sometimes also on Friday where possible. This is a chance for everybody to discuss what they've achieved so far and what problems they have faced. Using this information everybody can plan what they can achieve in the week ahead so the project remains on schedule. 

Tasks in this project will be organised in the issue tracker, accessible from the issue tab. Issues will also be discussed in the weekly standup meetings. 

To complete an issue, click on the issue in the tracker. Remember to:
* Assign yourself to the issue
* Assign the issue a tag that describes it: bug, documentation, or enhancement (feature)
* Associate the issue with the Medlock Primary School project

When working on the project, add your changes to a new branch. A new branch can be created in the Git tab in GoLand or using the command:  `git checkout -b <branch name>`

Make sure to commit as you go along, and push your commits so other team members can see how you are progressing. 

When you're finished with the issue, push all your changes. It would also make sense to run all the unit tests at this point to ensure you haven't introduced any bugs into the code. Now create a pull request (also called merge requests outside of Github) and associate it with the branch you created earlier. **Do not be tempted to merge your work directly into master.**

Github will automatically run the test suite when a new pull request is created. Karl will review your code and provided everything is correct it will be merged into the master branch. This should automatically close the issue you created earlier.

