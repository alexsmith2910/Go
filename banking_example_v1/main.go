package main

import (
	"banking_example_v1/account"
	"banking_example_v1/transaction"
	"fmt"
	"time"
)

func main() {
	user1 := account.Account{
		UserID:    "alexsmith2910",
		FirstName: "Alexander",
		LastName:  "Smith",
		Balance:   0.0,
	}

	user2 := account.Account{
		UserID:    "user2",
		FirstName: "Jacob",
		LastName:  "Santier",
		Balance:   32.0,
	}

	test_transaction := transaction.Transaction{
		PayeeUserID: &user1,
		PayerUserID: &user2,
		Title:       "Test Transaction",
		Amount:      25.0,
		Description: "This is a test transaction",
		Date:        time.Now(),
	}

	test_transaction2 := transaction.Transaction{
		PayeeUserID: &user2,
		PayerUserID: &user1,
		Title:       "Test Transaction 2",
		Amount:      12.49,
		Description: "This is a test transaction - decimal value",
		Date:        time.Now(),
	}

	transaction.Transfer(test_transaction)

	transaction.Transfer(test_transaction2)

	fmt.Printf("User1 Details:\n\tUserID: %v\n\tFirstName: %v\n\tLastName: %v\n\tBalance: %v\n",
		user1.UserID, user1.FirstName, user1.LastName, user1.Balance)

	fmt.Printf("User2 Details:\n\tUserID: %v\n\tFirstName: %v\n\tLastName: %v\n\tBalance: %v\n",
		user2.UserID, user2.FirstName, user2.LastName, user2.Balance)

	fmt.Printf("Transaction Details:\n\tTitle: %v\n\tAmount: %v\n\tDescription: %v\n\tDate: %v\n",
		test_transaction.Title, test_transaction.Amount,
		test_transaction.Description, test_transaction.Date)

	fmt.Printf("Transaction 2 Details:\n\tTitle: %v\n\tAmount: %v\n\tDescription: %v\n\tDate: %v\n",
		test_transaction2.Title, test_transaction2.Amount,
		test_transaction2.Description, test_transaction2.Date)
}
