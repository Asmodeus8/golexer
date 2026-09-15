package main
import("fmt";"os";"github.com/Asmodeus8/golexer/lexer")
func main(){if len(os.Args)<2{fmt.Println("usage: golexer <file>");return};b,err:=os.ReadFile(os.Args[1]);if err!=nil{panic(err)};l:=lexer.New(string(b));for{t:=l.Next();fmt.Printf("%-8s %q @ %d:%d\n",t.Kind,t.Literal,t.Line,t.Column);if t.Kind==lexer.EOF{break}}}
