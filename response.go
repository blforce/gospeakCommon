package gospeakCommon

type Response interface {
	AddText(string) Response
	SetImageCard(string, string, string) Response
	GetBytes() []byte
}
