# gator

This tool was created for the boot.dev course Build a Blog Aggregator in Go.

## Prerequisites

Using this tool requires:

* go - version 1.24.4 or later
* postgresql -  version 17.6 or later

## Getting started

1. Start a postgresql server that is accessible from the machine where you will be running this tool.
2. Check out the code for this repository.
3. Run `go build` and then `go install` in the root of the repository.
4. Create a file called `.gatorconfig.json` in your home directory. The file should contain the following two values:
    * db_url - The url for the postgres database to use.
    * current_user_name - Does not need to have a value to start.
Your file should look something like this:

```json
{"db_url":"postgres://username:password@localhost:5432/gator?sslmode=disable","current_user_name":""}
```

## Usage

The CLI can be run with the following syntax: `gator [command] [optional arguments]`

The available commands are:

* login - Takes an argument _username_ and sets the current user to the user with that username.
* register - Takes an argument _username_, creates a user with that username, and sets the current user to the user with that username.
* reset - Deletes all users.
* users - Lists all users, indicating the current user.
* agg - Takes an argument _time-between-reqs_ and polls all feeds at that interval looking for new posts. This is a long running task, and requires you to cress ctrl-c to break it.
* addfeed - Takes two arguments _name_ and _url_, adds a new feed with that name and url, and follows the field for the current user.
* feeds - Lists all feeds.
* follow - Takes an argument _url_ and follows the feed with that url for the current user.
* following - Lists all feeds followed by the current user.
* unfollow - Takes an argument _url_ and unfollows the feed with that url for the current user.
* browse - Takes an optional argument _number-of-posts_ and lists the most recent posts from the current users feeds, up to that number or 2, if no number is specified.
