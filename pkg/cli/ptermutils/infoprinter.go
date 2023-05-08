package ptermutils

import "github.com/pterm/pterm"

func NewCustumInfoPrinter(prefix string, prefixFixedLength int) pterm.PrefixPrinter {
	ptext := prefix
	leftOver := prefixFixedLength - len(ptext)
	if leftOver > 0 {
		for i := 0; i < leftOver; i++ {
			ptext = " " + ptext
		}
	}
	return pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  ptext,
			Style: pterm.Info.Prefix.Style,
		},
		Scope:            pterm.Info.Scope,
		MessageStyle:     pterm.Info.MessageStyle,
		Fatal:            false,
		ShowLineNumber:   false,
		LineNumberOffset: 0,
		Writer:           pterm.Info.Writer,
		Debugger:         false,
	}
}
