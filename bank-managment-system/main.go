package main

import "fmt"

type BankAccount struct {
    accountNo   string 
    holderName  string   
    balance     float64  
    accountType string   
    isActive    bool     
}

// Constructor function
func NewBankAccount(accountNo, holderName, accountType string) *BankAccount {
    return &BankAccount{
        accountNo:   accountNo,
        holderName:  holderName,
        balance:     0.0,
        accountType: accountType,
        isActive:    true,
    }
}

// Value receiver - only returns information
func (ba BankAccount) GetAccountInfo() string {
    return fmt.Sprintf("Account: %s, Holder: %s, Balance: %.2f", 
                      ba.accountNo, ba.holderName, ba.balance)
}

func (ba BankAccount) GetBalance() float64 {
    return ba.balance
}

func (ba BankAccount) IsActive() bool {
    return ba.isActive
}

// Pointer receiver - modifies balance
func (ba *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("Deposit amount must be greater than 0")
    }
    
    if !ba.isActive {
        return fmt.Errorf("Cannot transact on inactive account")
    }
    
    ba.balance += amount
    fmt.Printf("%.2f deposited. Current balance: %.2f\n", amount, ba.balance)
    return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("Withdrawal amount must be greater than 0")
    }
    
    if !ba.isActive {
        return fmt.Errorf("Cannot transact on inactive account")
    }
    
    if amount > ba.balance {
        return fmt.Errorf("Insufficient balance. Current balance: %.2f", ba.balance)
    }
    
    ba.balance -= amount
    fmt.Printf("%.2f withdrawn. Current balance: %.2f\n", amount, ba.balance)
    return nil
}

func (ba *BankAccount) Transfer(toAccount *BankAccount, amount float64) error {
    if err := ba.Withdraw(amount); err != nil {
        return fmt.Errorf("Transfer failed: %v", err)
    }
    
    if err := toAccount.Deposit(amount); err != nil {
        // If deposit fails in destination account, refund the money
        ba.balance += amount
        return fmt.Errorf("Transfer failed: %v", err)
    }
    
    fmt.Printf("%.2f transferred to %s\n", amount, toAccount.holderName)
    return nil
}

// Method to deactivate account
func (ba *BankAccount) DeactivateAccount() {
    ba.isActive = false
    fmt.Printf("Account %s deactivated\n", ba.accountNo)
}

func main() {
    // Create new accounts
    account1 := NewBankAccount("123456", "Mohammad Shahriar", "Savings")
    account2 := NewBankAccount("789012", "Tuli", "Current")
    
    fmt.Println("=== Account Information ===")
    fmt.Println(account1.GetAccountInfo())
    fmt.Println(account2.GetAccountInfo())
    
    fmt.Println("\n=== Transactions ===")
    // Deposit money
    account1.Deposit(5000)
    account2.Deposit(3000)
    
    // Withdraw money  
    account1.Withdraw(1500)
    
    // Transfer money
    account1.Transfer(account2, 1000)
    
    fmt.Println("\n=== Final Balances ===")
    fmt.Printf("Account1 Balance: %.2f\n", account1.GetBalance())
    fmt.Printf("Account2 Balance: %.2f\n", account2.GetBalance())
}


