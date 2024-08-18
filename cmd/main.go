package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vshulcz/gopass/internal"
	"github.com/vshulcz/gopass/internal/crypto"
	"github.com/vshulcz/gopass/internal/db"
	"github.com/vshulcz/gopass/internal/services"
	"github.com/vshulcz/gopass/internal/storage"

	"github.com/spf13/cobra"
)

var (
	passwordService *services.PasswordService
	masterKey       []byte
)

func main() {
	cfg := internal.LoadConfig()

	dbConn, err := db.NewDatabaseConnection(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	store := storage.NewGormStorage(dbConn)

	masterKey, err = crypto.InitializeMasterKey(cfg.MasterKeyPath)
	if err != nil {
		log.Fatalf("Failed to initialize master key: %v", err)
	}

	passwordService = services.NewPasswordService(store, masterKey)

	var rootCmd = &cobra.Command{Use: "gopass"}
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var addCmd = &cobra.Command{
	Use:   "add [service]",
	Short: "Add a new password",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var service, username, password string
		if len(args) > 0 {
			service = args[0]
		} else {
			fmt.Print("Enter service name: ")
			fmt.Scan(&service)
		}
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		err := passwordService.AddPassword(service, username, password)
		if err != nil {
			log.Fatalf("Failed to add password: %v", err)
		}

		fmt.Println("Password added successfully.")
	},
}

var getCmd = &cobra.Command{
	Use:   "get [service]",
	Short: "Get a password",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var service string
		if len(args) > 0 {
			service = args[0]
		} else {
			fmt.Print("Enter service name: ")
			fmt.Scan(&service)
		}

		username, password, err := passwordService.GetPassword(service)
		if err != nil {
			if err.Error() == "record not found" {
				fmt.Printf("No password found for service: %s\n", service)
			} else {
				log.Fatalf("Failed to get entry: %v", err)
			}
			return
		}

		fmt.Printf("Service: %s\nUsername: %s\nPassword: %s\n", service, username, password)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored passwords",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := passwordService.ListPasswords()
		if err != nil {
			log.Fatalf("Failed to list entries: %v", err)
		}

		for _, entry := range entries {
			fmt.Printf("Service: %s, Username: %s\n", entry.Service, entry.Username)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [service]",
	Short: "Delete a password",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var service string
		if len(args) > 0 {
			service = args[0]
		} else {
			fmt.Print("Enter service name: ")
			fmt.Scan(&service)
		}

		err := passwordService.DeletePassword(service)
		if err != nil {
			log.Fatalf("Failed to delete entry: %v", err)
		}

		fmt.Println("Password deleted successfully.")
	},
}
