package lexer

import "unicode"

type Kind string
const( Identifier Kind="IDENT"; Number Kind="NUMBER"; String Kind="STRING"; Operator Kind="OP"; EOF Kind="EOF"; Illegal Kind="ILLEGAL")
type Token struct{ Kind Kind; Literal string; Line, Column int }
type Lexer struct{ src []rune; pos,line,col int }
func New(s string)*Lexer{return &Lexer{src:[]rune(s),line:1,col:1}}
func(l *Lexer) Next()Token{ for l.pos<len(l.src)&&unicode.IsSpace(l.src[l.pos]){l.advance()}; if l.pos>=len(l.src){return Token{Kind:EOF,Line:l.line,Column:l.col}}; line,col:=l.line,l.col; ch:=l.src[l.pos]; if unicode.IsLetter(ch)||ch=='_' {start:=l.pos; for l.pos<len(l.src)&&(unicode.IsLetter(l.src[l.pos])||unicode.IsDigit(l.src[l.pos])||l.src[l.pos]=='_'){l.advance()}; return Token{Identifier,string(l.src[start:l.pos]),line,col} }; if unicode.IsDigit(ch){start:=l.pos; for l.pos<len(l.src)&&unicode.IsDigit(l.src[l.pos]){l.advance()}; return Token{Number,string(l.src[start:l.pos]),line,col} }; if ch=='"'{l.advance();start:=l.pos;for l.pos<len(l.src)&&l.src[l.pos]!='"'{l.advance()};lit:=string(l.src[start:l.pos]);if l.pos>=len(l.src){return Token{Illegal,lit,line,col}};l.advance();return Token{String,lit,line,col}}; if contains("+-*/%=<>!(){}[],;",ch){l.advance();return Token{Operator,string(ch),line,col}};l.advance();return Token{Illegal,string(ch),line,col} }
func(l *Lexer)advance(){if l.src[l.pos]=='\n'{l.line++;l.col=1}else{l.col++};l.pos++}
func contains(s string,r rune)bool{for _,x:=range s{if x==r{return true}};return false}
