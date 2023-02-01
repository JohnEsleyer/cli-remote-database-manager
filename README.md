# CLI Remote Database Manager
A command-line program that connects to a remote MySQL database (developed for AWS RDS).
It allows you to create, delete, and display tables. As well as, adding and deleting table instances or rows. The program uses the Go SQL driver for MySQL and the database/sql package to handle the database connection and querying.

### How to build the binary?
Type the following in your terminal window (for Linux) or cmd (for Windows)
```
go build
```

### Usage
```
./cli-remote-database <dbUser> <dbPass> <dbName> <dbHost> <dbPort>
```

##### Choose one of the options given
![image](https://user-images.githubusercontent.com/66754038/215781216-d08a41fe-d6d5-478a-824a-3ead9c9b0c39.png)
