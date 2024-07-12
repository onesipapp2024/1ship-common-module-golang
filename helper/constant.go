package helper

type ClientType uint
type ChoiceString uint

const (
	ClientTypeUndefined ClientType = iota
	ClientTypeIOS
	ClientTypeAndroid
	ClientTypeBrowser
	ClientTypeBot
	ClientTypeOther
)

const (
	RandomStringSource      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandomUpperStringSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandomLowerStringSource = "0123456789abcdefghijklmnopqrstuvwxyz"
)

const (
	ChoiceStringUndefined ChoiceString = iota
	ChoiceStringUpper
	ChoiceStringLower
)

func (t ClientType) String() string {
	switch t {
	case ClientTypeIOS:
		return "ios"
	case ClientTypeAndroid:
		return "android"
	case ClientTypeBrowser:
		return "browser"
	case ClientTypeBot:
		return "bot"
	case ClientTypeOther:
		return "other"
	}
	return "unknown"
}
