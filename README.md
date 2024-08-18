# GoPass Manager

GoPass Manager is a lightweight, command-line tool for securely managing your passwords. This tool allows you to store, retrieve, and manage your passwords with strong encryption, making it easy to keep your sensitive information safe.

## Features

Secure Storage: Passwords are encrypted using AES-256 encryption.
Command-Line Interface: Manage your passwords directly from the terminal.
SQLite Database: All passwords are stored in a local SQLite database.
Master Key: Protects your encrypted data with a master key, which is generated automatically during the first run.

## Installation
Clone the repository:
```bash
git clone https://github.com/vshulcz/gopass-manager.git
cd gopass-manager
```
Install the application:
```bash
make install
```
This command will build the project and install the password-manager command globally, making it accessible from any directory in your terminal.

## Usage
    
### Add a new password

To add a new password to the manager, use the add command:

```bash
gopass add [service]
```

You'll be prompted to enter the service name, username, and password.

### Retrieve a password

To retrieve a password for a specific service, use the get command:

```bash
gopass get [service]
```

You'll be prompted to enter the service name. The username and decrypted password will be displayed.

### List all stored passwords

To list all stored services with their usernames, use the list command:

```bash
gopass list
```

### Delete a password

To delete a password from the manager, use the delete command:

```bash
gopass delete [service]
```

You'll be prompted to enter the service name, and the associated entry will be removed from the database.

## Uninstalling

If you ever need to remove the application, simply move to the project repository and run:

```bash
make uninstall
```

## Master Key Management

The master key is automatically generated on the first run and stored in a hidden file named .gopass_key in your home directory (~/.gopass_key). This key is essential for accessing your encrypted data. Make sure to back it up securely. If you lose this key, you will not be able to access your stored passwords.

## Contributing

Contributions are welcome! If you have any suggestions or improvements, feel free to fork the repository, make your changes, and submit a pull request.

## Disclaimer

This project is intended for educational purposes.