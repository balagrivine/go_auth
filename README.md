# GOAuth

GOAuth is a Go project for implementing a robust and secure user authentication system.

## Prerequisities
Ensure you have Goland installed on your system. If not, you can download and install it from [Golang official site](https://go.dev/doc/install)

## Clone the repository

Clone the go_auth repository using the command below to get started:

```bash
git clone https://github.com/balagrivine/go_auth/git
```

## Database Setup

The project uses PostgreSQL. While Windows is technically supported, using a Unix-like system
(macOS or Linux) is recommended for better long-term compatibility.
* ForWindows users: [Guide to install PostgreSQL on Windows](https://www.geeksforgeeks.org/install-postgresql-on-windows/)
* For Linux users: [Guide to install PostgreSQL on Linux](https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql-macos/)
* For MacOS: [Guide to install PostgreSQL on MacOS](https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql-linux/)

After installation, update the database URL in the .env file with your local PostgreSQL URL.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
