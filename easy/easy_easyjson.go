// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package easy

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_decode_github_com_centrifugal_easyerror_easy_clientCommand(in *jlexer.Lexer, out *clientCommand) {
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "uid":
			out.UID = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Params).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_encode_github_com_centrifugal_easyerror_easy_clientCommand(out *jwriter.Writer, in *clientCommand) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"uid\":")
	out.String(string(in.UID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"method\":")
	out.String(string(in.Method))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"params\":")
	out.Raw((in.Params).MarshalJSON())
	out.RawByte('}')
}
func (v *clientCommand) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_encode_github_com_centrifugal_easyerror_easy_clientCommand(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v *clientCommand) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_encode_github_com_centrifugal_easyerror_easy_clientCommand(w, v)
}
func (v *clientCommand) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_decode_github_com_centrifugal_easyerror_easy_clientCommand(&r, v)
	return r.Error()
}
func (v *clientCommand) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_decode_github_com_centrifugal_easyerror_easy_clientCommand(l, v)
}
