package enums

type AccessType string

const (
	READ_ACCESS  AccessType = "READ"
	WRITE_ACCESS AccessType = "WRITE"
)

type MessageType string

const (
	TEXT       MessageType = "TEXT"
	BOOTUP     MessageType = "BOOTUP"
	DISCONNECT MessageType = "DISCONNECT"
)
