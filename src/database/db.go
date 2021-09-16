package db

func Connect() {
	_ = map[string]string{
		"transaction_isolation": "'READ-COMMITED'",
	}
}
