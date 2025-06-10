package lexer

type Lexer struct {
	input	string
	position	int
	readPosition	int
	char	byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	return lexer
}

func (lexer *Lexer) ReadChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition ++
}
