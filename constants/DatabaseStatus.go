package constants

type DBOperationStatus string

const (
	CREATED DBOperationStatus = "CREATED RECORD TO DB"
	UPDATED DBOperationStatus = "UPDATED RECORD FROM DB"
	QUERIED DBOperationStatus = "QUERIED RECORD FROM DB"
	DELETED DBOperationStatus = "DELETED RECORD FROM DB"
	ERROR   DBOperationStatus = "ERROR WHILE SAVING IN DB"
)