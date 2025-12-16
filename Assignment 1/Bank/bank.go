package Bank

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BankAccount struct {
	Name               string
	Surname            string
	balance            float64
	CardNumber         string
	transactionHistory []string
}

func NewBankAccount(name, surname, CardNumber string) *BankAccount {

	return &BankAccount{
		Name:               name,
		Surname:            surname,
		balance:            0,
		CardNumber:         CardNumber,
		transactionHistory: []string{},
	}

}

func (bankAccount *BankAccount) updateTransactionHistory(message string) {

	bankAccount.transactionHistory = append(bankAccount.transactionHistory, message)
}

func (bankAccount *BankAccount) deposit(depositMoney float64) {
	bankAccount.balance += depositMoney
	message := fmt.Sprintf("Deposit : %.2f", depositMoney)
	bankAccount.updateTransactionHistory(message)
}

func (bankAccount *BankAccount) withdraw(withdrawMoney float64) (status int) {
	if withdrawMoney <= bankAccount.balance {
		bankAccount.balance -= withdrawMoney
		message := fmt.Sprintf("Withdraw : %.2f, Balance left: %.2f", withdrawMoney, bankAccount.balance)
		bankAccount.updateTransactionHistory(message)
		return 0
	} else {
		message := fmt.Sprintf("Withdraw declined! insufficient balance. Withdraw : %.2f, Current Balance: %.2f", withdrawMoney, bankAccount.balance)
		bankAccount.updateTransactionHistory(message)
		return 1
	}
}

func (bankAccount *BankAccount) getBalance() float64 {

	return bankAccount.balance

}

func (bankAccount *BankAccount) getTransactionHistory() []string {

	return bankAccount.transactionHistory

}

func (bankAccount *BankAccount) CLI() {
	var scanner = bufio.NewScanner(os.Stdin)
	var cursor string
	var money float64
	for {
		fmt.Printf("\nWelcome to bank system!\nCommands:\n1. Deposit money\n2. Withdraw money\n3. Get balance\n4. Get transaction history\n5. Exit\n")
		scanner.Scan()
		cursor = scanner.Text()
		switch cursor {

		case "1":
			fmt.Println("\nHow much money you want to deposit?\n")
			scanner.Scan()
			money, _ = strconv.ParseFloat(scanner.Text(), 64)
			bankAccount.deposit(money)
			fmt.Printf("You successfully deposited %.2f!", money)
		case "2":
			fmt.Println("\nHow much money you want to withdraw?\n")
			scanner.Scan()
			money, _ = strconv.ParseFloat(scanner.Text(), 64)
			if bankAccount.withdraw(money) == 0 {
				fmt.Printf("You successfully withdrawed %.2f!", money)
			} else {
				fmt.Printf("You don't have enough money on your deposit!")
			}
		case "3":
			fmt.Printf("Currently you have: %.2f on your deposit.", bankAccount.getBalance())
		case "4":
			var transactions []string = bankAccount.getTransactionHistory()
			if transactions != nil {
				fmt.Println("\nTRANSACTION HISTORY\n")
				for i, transaction := range transactions {
					fmt.Printf("%d. %s\n", i+1, transaction)
				}
			} else {
				fmt.Println("Transaction list is empty.\n")
			}
		case "5":
			fmt.Println("Thank you for using my banking system app. Goodbye! \n")
			return
		}
	}
}
