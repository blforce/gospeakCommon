package gospeakCommon

type Response interface {
	SetText(string) Response
	GetBytes() []byte
}
