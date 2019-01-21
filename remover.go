package replyremover

import (
	"bufio"
	"regexp"
	"strings"
)

var (
	signatureRegex    = regexp.MustCompile(`(?s)(?:^\s*--|^\s*__|^-\w|^-- $)|(?:^(Sent from my|Envoyé de mon) (?:\s*\w+){1,4}$)|(?:^={30,}$)$`)
	quoteRegex        = regexp.MustCompile(`(?s)^>+`)
	quoteHeadersRegex []*regexp.Regexp
)

func init() {
	quoteHeadersRegex = []*regexp.Regexp{
		regexp.MustCompile(`(?ms)^[\s>]*(On\s{1,10}.{1,100}\s+wrote:)$`),
		regexp.MustCompile(`(?ms)^[\s>]*(Le{1,10}.{1,100}\s+écrit :)$`),
		regexp.MustCompile(`(?ms)^[\s>]*(El{1,10}.{1,100}\s+escribió:)$`),
		regexp.MustCompile(`(?ms)^[\s>]*(Il{1,10}.{1,100}\s+scritto:)$`),
		regexp.MustCompile(`(?ms)^[\s>]*(El{1,10}.{1,100}\s+escriure:)$`),
		regexp.MustCompile(`(?m)^.+\s+(написа(л|ла|в)+)+:$`),
		regexp.MustCompile(`(?ms)^[\s>]*(Op\s.+?schreef.+:)$`),
		regexp.MustCompile(`(?ms)^[\s>]*((W\sdniu|Dnia)\s.+?(pisze|napisał(\(a\))?):)$`),
		regexp.MustCompile(`(?m)^[\s>]*(Den\s.+\sskrev\s.+:)$`),
		regexp.MustCompile(`(?m)^[\s>]*(Am\s.+\sum\s.+\sschrieb\s.+:)$`),
		regexp.MustCompile(`(?ms)^(在.+写道：)$`),
		regexp.MustCompile(`(?m)^(20[0-9]{2}\..+\s작성:)$`),
		regexp.MustCompile(`(?m)^(20[0-9]{2}\/.+のメッセージ:)$`),
		regexp.MustCompile(`(?m)^(.+\s<.+>\sschrieb:)$`),
		regexp.MustCompile(`(?m)^[\s>]*(From\s?:.+\s?(\[|<).+(\]|>))`),
		regexp.MustCompile(`(?m)^[\s>]*(发件人\s?:.+\s?(\[|<).+(\]|>))`),
		regexp.MustCompile(`(?m)^[\s>]*(De\s?:.+\s?(\[|<).+(\]|>))`),
		regexp.MustCompile(`(?m)^[\s>]*(Van\s?:.+\s?(\[|<).+(\]|>))`),
		regexp.MustCompile(`(?m)^[\s>]*(Da\s?:.+\s?(\[|<).+(\]|>))`),
		regexp.MustCompile(`(?ms)^(20[0-9]{2}\-(?:0?[1-9]|1[012])\-(?:0?[0-9]|[1-2][0-9]|3[01]|[1-9])\s[0-2]?[0-9]:\d{2}\s.+?:)$`),
		regexp.MustCompile(`(?ms)^[\s>]*([a-z]{3,4}\.\s.+\sskrev\s.+:)$`),
		regexp.MustCompile(`(?ms)^(\[image\:\s.+\](.+){0,100}\*.+\*)$`),
		regexp.MustCompile(`(?ms)^\*(.{1,100})\*.{1,50}\*.{1,100}liberkeys\.com.*\*$`),
		regexp.MustCompile(`(?ms)^\s{3,}\*.*\*.{1,100}liberkeys\.com.*$`),
		regexp.MustCompile(`(?ms)^.*(<div dir="ltr">).*$`),
	}
}

func isQuoteHeader(line string) bool {
	for _, regex := range quoteHeadersRegex {
		if regex.MatchString(line) {
			return true
		}
	}
	return false
}

// RemoveReplies removes the email replies from an email text body
func RemoveReplies(in string) string {
	in = strings.Replace(in, "\r\n", "\n", -1)

	for _, regex := range quoteHeadersRegex {
		for _, match := range regex.FindAllString(in, -1) {
			withoutLinefeed := strings.Replace(match, "\n", " ", -1)
			in = strings.Replace(in, match, withoutLinefeed, -1)
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(in))
	out := ""

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), " \n")

		if isQuoteHeader(line) || signatureRegex.MatchString(line) {
			return strings.TrimSpace(out)
		}

		out += line + "\n"
	}

	return strings.TrimSpace(out)
}
