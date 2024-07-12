package validator

const (
	PasswordRegexString = "^[A-Za-z0-9\\`\\~\\!\\@\\#\\$\\%\\^\\&\\*\\(\\)\\-\\_\\=\\+\\[\\{\\]\\}\\\\\\|\\;\\:'\\\"\\,\\<\\.\\>\\/\\?]+$"
	DateTimeRegexString = "^([0-9]{4})\\-(1[0-2]|0[1-9])\\-(3[0-1]|0[1-9]|[1-2][0-9]) (2[0-3]|[0-1][0-9])\\:([0-5][0-9])\\:([0-5][0-9])$"
)

const (
	ContainFourByteCharacterRegexString = "[\\x{10000}-\\x{10FFFF}]"
)
