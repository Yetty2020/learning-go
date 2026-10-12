package main


//Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. Instead of using if/else chains, we just return early from the function at the end of each conditional block.
// Nobody has to remember “I’m inside the passport check, which is inside the ticket check…”. You go through the checks one at a time, top to bottom. That’s a guard clause.


//Nested Style

func canWithdrawNest(balance, amount int, accountLocked bool) string {
	if !accountLocked {
		if amount > 0 {
			if balance >= amount {
				return "approved"
			} else {
				return "insufficient funds"
			}
		} else {
			return "invalid amount"
		}
	} else {
		return "account locked"
	}
}

//Guard clause

func canWithdrawGuard(balance, amount int, accountLocked bool) string{
	if accountLocked{
		return "account locked"

	}
	if amount <= 0 {
		return "invalid amount"
	}
	if balance < amount{
		return "insufficient funds"
	} 
	return "approved"
}