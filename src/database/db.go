package db

func Connect() {
	_ = map[string]string{
		"tx_isolation": "'READ-COMMITED'",
	}
}
