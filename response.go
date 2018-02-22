package gospeakCommon

type Response interface {
	SetText(string) Response
	SetImageCard(string, string, string) Response
	GetBytes() []byte
}
