package transaction

import (
	"errors"
	"fmt"
	"kv-store/util"
)

// map to store global data.
var GlobalDb = make(map[string]string)

// pointer to parent, and its own internal store. from the example README: it is able to READ from its internal store and COMMIT bubbles up storage to its parent (or eventually the global db)
type Transaction struct {
	parent *Transaction
	db     map[string]string
}

// basically just a Stack data structure of Transactions. transactions are pushed/popped from the stack as they are started and commited.
type Transactions struct {
	curr *Transaction
}

// Start a transaction. Pushes a transaction struct to the root stack and changes the pointer to the current transaction to the most recent transaction.
func (t *Transactions) StartTransaction() {
	if t == nil {
		t = &Transactions{}
		var transaction = Transaction{db: map[string]string{}}
		transaction.parent = t.curr
		t.curr = &transaction
	} else {
		var transaction = Transaction{db: map[string]string{}}
		//nested transaction inherits parent DB
		if t.curr == nil {
			util.PopulateDb(GlobalDb, transaction.db)
		} else {
			util.PopulateDb(t.curr.db, transaction.db)
		}
		transaction.parent = t.curr
		t.curr = &transaction
	}

}

// Commit the current transaction. Copies the data to either its parent DB (or the root DB) and then pops the transaction from the root stack.
func (t *Transactions) Commit() error {
	if t == nil || t.curr == nil {
		return errors.New("There is no current transaction to commit.")
	}
	currentTransaction := t.curr
	if currentTransaction.parent != nil {
		util.PopulateDb(currentTransaction.db, currentTransaction.parent.db)
	} else {
		util.PopulateDb(currentTransaction.db, GlobalDb)
	}
	t.curr = t.curr.parent //pops current transaction
	return nil
}

// write to either global DB, or the current transaction.
func (t *Transactions) Write(key string, value string) error {
	if t == nil {
		return errors.New("Can't write data, root transaction store is nil!")
	}
	var curr = t.curr
	if curr == nil {
		GlobalDb[key] = value
	} else {
		curr.db[key] = value
	}
	return nil
}

// read from either global DB, or the current transaction. return any errors.
func (t *Transactions) Read(key string) (string, error) {
	var curr = t.curr
	if curr == nil {
		var val = GlobalDb[key]
		if val == "" {
			return "", fmt.Errorf("Key not found: %s", key)
		} else {
			return GlobalDb[key], nil
		}
	} else {
		var val = curr.db[key]
		if val == "" {
			return "", fmt.Errorf("Key not found: %s", key)
		} else {
			return curr.db[key], nil
		}
	}
}

// delete from either current DB, or the current transactions DB
func (t *Transactions) Delete(key string) error {
	if t == nil {
		return errors.New("Can't delete data, root transaction store is nil!")
	}
	var curr = t.curr
	if curr == nil {
		GlobalDb[key] = ""
	} else {
		curr.db[key] = ""
	}
	return nil
}

// Abort the current transaction
func (t *Transactions) Abort() error {
	if t.curr == nil {
		return errors.New("There is no current transaction to abort.")
	}
	t.curr = t.curr.parent //pop the stack
	return nil
}
