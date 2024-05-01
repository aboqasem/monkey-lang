package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-lang/lexer"
	"monkey-lang/token"
)

const (
	PROMPT = ">> "
	QUIT   = "\\q"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == QUIT {
			break
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
