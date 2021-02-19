package parse

import (
    "bufio"
    "bytes"
    "io"
    "regexp"
)

var (
    boldItalicReg = regexp.MustCompile(`\*\*\*(.*?)\*\*\*`)
    boldReg       = regexp.MustCompile(`\*\*(.*?)\*\*`)
    italicReg     = regexp.MustCompile(`\*(.*?)\*`)
    strikeReg     = regexp.MustCompile(`\~\~(.*?)\~\~`)
    underscoreReg = regexp.MustCompile(`__(.*?)__`)
    anchorReg     = regexp.MustCompile(`\[(.*?)\]\((.*?)\)[^\)]`)
    escapeReg     = regexp.MustCompile(`^\>(\s|)`)
    blockquoteReg = regexp.MustCompile(`\&gt\;(.*?)$`)
    backtipReg    = regexp.MustCompile("`(.*?)`")

    h1Reg = regexp.MustCompile(`^#(\s|)(.*?)$`)
    h2Reg = regexp.MustCompile(`^##(\s|)(.*?)$`)
    h3Reg = regexp.MustCompile(`^###(\s|)(.*?)$`)
    h4Reg = regexp.MustCompile(`^####(\s|)(.*?)$`)
    h5Reg = regexp.MustCompile(`^#####(\s|)(.*?)$`)
    h6Reg = regexp.MustCompile(`^######(\s|)(.*?)$`)
)

func NewMarkdown(input io.Reader) string {

    buf := bytes.NewBuffer(nil)

    scanner := bufio.NewScanner(input)
    for scanner.Scan() {

        line := bytes.TrimSpace(scanner.Bytes())
        if len(line) == 0 {
            buf.WriteByte('\n')
            continue
        }

        // wrap bold and italic text in "<b>" and "<i>" elements
        line = boldItalicReg.ReplaceAll(line, []byte(`<b><i>$1</i></b>`))
        line = boldReg.ReplaceAll(line, []byte(`<b>$1</b>`))
        line = italicReg.ReplaceAll(line, []byte(`<i>$1</i>`))
        // wrap strikethrough text in "<s>" tags
        line = strikeReg.ReplaceAll(line, []byte(`<s>$1</s>`))
        // wrap underscored text in "<u>" tags
        line = underscoreReg.ReplaceAll(line, []byte(`<u>$1</u>`))
        // convert links to anchor tags
        line = anchorReg.ReplaceAll(line, []byte(`<a href="$2">$1</a>`))
        // escape and wrap blockquotes in "<blockquote>" tags
        line = escapeReg.ReplaceAll(line, []byte(`&gt;`))
        line = blockquoteReg.ReplaceAll(line, []byte(`<blockquote>$1</blockquote>`))
        // wrap the content of backticks inside of "<code>" tags
        line = backtipReg.ReplaceAll(line, []byte(`<code>$1</code>`))
        // convert headings
        if line[0] == '#' {

            count := bytes.Count(line, []byte(`#`))
            switch count {
            case 1:
                line = h1Reg.ReplaceAll(line, []byte(`<h1>$2</h1>`))
            case 2:
                line = h2Reg.ReplaceAll(line, []byte(`<h2>$2</h2>`))
            case 3:
                line = h3Reg.ReplaceAll(line, []byte(`<h3>$2</h3>`))
            case 4:
                line = h4Reg.ReplaceAll(line, []byte(`<h4>$2</h4>`))
            case 5:
                line = h5Reg.ReplaceAll(line, []byte(`<h5>$2</h5>`))
            case 6:
                line = h6Reg.ReplaceAll(line, []byte(`<h6>$2</h6>`))
            }
        }
        buf.Write(line)
        buf.WriteByte('\n')
    }
    return buf.String()
}
