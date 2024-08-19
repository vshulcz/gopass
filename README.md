# GoPass Manager

GoPass Manager is a lightweight, command-line tool for securely managing your passwords.

## Features

Secure Storage: Passwords are encrypted using AES-256 encryption.
Command-Line Interface: Manage your passwords directly from the terminal.
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

## Usage
    
### Add a new password

To add a new password to the manager, use the add command:

```bash
gopass add [service]
```

### Retrieve a password

To retrieve a password for a specific service, use the get command:

```bash
gopass get [service]
```

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

## Uninstalling

If you ever need to remove the application, simply move to the project repository and run:

```bash
make uninstall
```

## Contributing

Contributions are welcome! If you have any suggestions or improvements, feel free to fork the repository, make your changes, and submit a pull request.

## Disclaimer

This project is intended for educational purposes.