**Prerequisites:**
Gator is a CLI blog aggregator that requires PostgreSQL database and Go programming language pre-installed.

**Installation:**
1. go install github.com/Ceramik9/blog_agregator@latest
2. Create ".gatorconfig.json" file in your home directory
3. Paste below text into the file:
{"db_url":"postgres://username@localhost:5432/dbname?sslmode=disable","current_user_name":""}
4. Update the "db_url" to point it at your database
5. You can use "goose" and sql files to migrate the database to version 005 using sql files in the "sql/schema" directory

**Commands**
1. register "username" - register a user. It takes one argument, the username.
2. login "username" - log in as existing user. It takes one argument, the username.
3. users - print a list of registered users. It doesn't take any arguments.
4. addfeed "feed title" "feed url" - add feed to database. It takes two arguments, the title and the url.
5. follow "feed url" - follow feed. It takes one argument, the feed url.
6. unfollow "feed url" - unfollow feed. It takes one argument, the feed url.
7. feeds - print list of all feeds. It doesn't take any arguments.
8. following - print a list of all feeds currently logged in user follows. It doesn't take any arguments.
9. agg "number in seconds" - save posts from all feeds current user is following. It takes one argument, number in seconds that determines how often the program will download the posts.
10. browse "num of posts" - prints all posts for the feeds current user is following. It takes one argument, the limit, which determines the limit of the posts that will be printed.
11. reset - it erases all data in the database. It doesn't take any arguments. Use wisely.


