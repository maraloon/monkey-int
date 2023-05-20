package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    illegal = "ILLEGAL"
    EOF = "EOF"
    
    IDENT = "IDENT"
    INT = "INT"

    ASSIGN = "="
    PLUS = "+"

    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    
    FUNCTION = "FUNCTION"
    let = "let"
)
