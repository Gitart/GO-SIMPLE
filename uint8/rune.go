// Use rune
// https://github.com/chzyer/readline/tree/master/runes

// translate Esc[X
func escapeExKey(key *escapeKeyPair) rune {
	var r rune
	switch key.typ {
	case 'D':
		r = CharBackward
	case 'C':
		r = CharForward
	case 'A':
		r = CharPrev
	case 'B':
		r = CharNext
	case 'H':
		r = CharLineStart
	case 'F':
		r = CharLineEnd
	case '~':
		if key.attr == "3" {
			r = CharDelete
		}
	default:
	}
	return r
}
