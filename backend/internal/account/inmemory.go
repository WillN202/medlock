package account

func NewInMemoryStore() Store {
	return &inmemory{
		accounts: make(map[string]Account),
	}
}

type inmemory struct {
	accounts map[string]Account
}

func (i *inmemory) AccountExists(id string) bool {
	for _, acc := range i.accounts {
		if acc.ID() != id {
			continue
		}
		return true
	}
	return false
}

type CodeConflictError struct {
	code string
}

type CodeDoesNotExistError struct {
	code string
}

func (e *CodeDoesNotExistError) Error() string {
	return "Account Does not exist " + e.code
}

func (e CodeConflictError) Error() string {
	return "account already exists with code " + e.code
}

func (i *inmemory) SaveAccount(account Account) error {
	if _, ok := i.accounts[account.Code()]; ok {
		return CodeConflictError{code: account.Code()}
	}
	i.accounts[account.Code()] = account
	return nil
}

func (i *inmemory) Login(code string) (Account, error) {
	stud, ok := i.accounts[code]
	if !ok {
		return nil, &CodeDoesNotExistError{code: code}
	}
	return stud, nil
}
