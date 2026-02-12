package repl

import(
	"bufio"
	"fmt"
	"io"
	"Hill/lexer"
	"Hill/token"
)

const PROMPT string = "--"

func Start(in io.Reader, out io.Writer) {
	var scanner *bufio.Scanner = bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		
		var scanned bool = scanner.Scan()
		if !scanned {
			return
		}

		var line string = scanner.Text()
		var l *lexer.Lexer = lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.TokenType(token.EOF); tok = l.NextToken() {
			fmt.Fprintf(out, "% + v \n", tok)
		}
	}
}
