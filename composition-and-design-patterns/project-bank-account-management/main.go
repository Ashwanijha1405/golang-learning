package main

import (
	"fmt"
	"errors"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *Account) Deposit(amount float64) error {

	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	acc.Balance += amount
	fmt.Printf("Deposited $%.2f to %s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil

}

func (acc *Account) Withdraw(amount float64) error {

	if amount <= 0 {
		return errors.New("Withdrawal amount must be positive")
	}

	if acc.Balance < amount {
		return fmt.Errorf("insufficient funds in %s. Balance: $%.2f, Tried to Withdraw: $%.2f",
			acc.AccountNumber, acc.Balance, amount)
	}

	acc.Balance -= amount
	fmt.Printf("Withdraw $%.2f from %s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil

}

func (acc *Account) String() string {
	return fmt.Sprintf("Account [%s] Owner: %s, Balance: $%.2f",
		acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account              // Embed Account struct ( annonymous field )
	InterestRate float64 // e.g., 0.02 for 2%
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate // Accesses promoted Balance field
	fmt.Printf("Adding Interest $%.2f to savings account %s. ", interest, sa.AccountNumber)
	err := sa.Deposit(interest) // Uses promoted Deposit method
	if err != nil {
		fmt.Printf("AddInterest: Error depositing $%.2f to savings account. %v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account       // Embed account struct
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be positive")
	}

	// Allow withdraw upto balance + overdraftLimit
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("Withdrawal of $%.2f exceeds overdraft limit for %s, Available including overdraft: $%.2f", amount, oa.AccountNumber, oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance -= amount // Balance can go negative
	fmt.Printf("Withdrew $%.2f from overdraft account %s, New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {

	fmt.Println("----------------- Bank Account System ----------------------")

	savAcc := SavingsAccount{

		Account: Account{ // initialise the embed account
			AccountNumber: "SAV101",
			Balance:       1000.00,
			OwnerName:     "Alice Saver",
		},
		InterestRate: 0.02, // 2%
	}
	fmt.Printf("\n--------- Savings Account Operations -----------")
	fmt.Printf(savAcc.Account.String())

	err := savAcc.Deposit(200.00) 
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v\n", 200.00, err)
	}
	savAcc.AddInterest()
	err = savAcc.Withdraw(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Saving details:", savAcc.Account.String())

	ovdAcc := OverdraftAccount{
		Account: Account{
			AccountNumber: "0VD002", 
			Balance: 100.00, 
			OwnerName: "Bob Spender", 
		},
		OverdraftLimit: 200.00, 
	}

	fmt.Println("\n---------- Overdraft Account Operations ------------")
	fmt.Println(ovdAcc.Account.String())

	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
