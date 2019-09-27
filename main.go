
  
package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
)

type tokenType int

const (
	tokenWord tokenType = iota
	tokenWhitespace
	tokenPunctuation
	tokenEof
	tokenRelop
	tokenSpecial
	tokenOperator
	tokenNumber
	tokenNewline
	tokenPrimitiveTypes
)

var names = map[tokenType]string{
	tokenWord:        "WORD",
	tokenWhitespace:  "SPACE",
	tokenPunctuation: "PUNCTUATION",
	tokenEof:         "EOF",
	tokenRelop:	 "REL_OP",
	tokenSpecial:	 "SPECIAL",
	tokenOperator:	 "OPERATOR",
	tokenNumber:	 "NUMBER_CONST",
	tokenNewline:	 "NEWLINE",
 tokenPrimitiveTypes:"PRIMITIVE",
}
type token_data struct {
    values [2]interface{}
}
//Must Compile is used to define the regex structure. It is different from compile in the sense that it panics when the match is violated
var wordRegexp = regexp.MustCompile("[A-Za-z]+")	
var whitespaceRegexp = regexp.MustCompile("[\\s]+")
var specialRegexp = regexp.MustCompile("[\\(\\)\\[\\]\\{\\}]+")
var punctuationRegexp = regexp.MustCompile("[\\:\\;\\,]+")
var reloperatorRegexp = regexp.MustCompile("[\\=\\<\\>\\!]+")
var operatorRegexp = regexp.MustCompile("[\\+\\-\\*\\/\\%]+")
var numberRegexp = regexp.MustCompile("[0-9]+")
var newlineRegexp = regexp.MustCompile("[\n]+")
var primitiveRegexp = regexp.MustCompile("int")
type token struct {
	value     string
	pos       int
	tokenType tokenType
}

func (tok token) String() string {
	return fmt.Sprintf("{%s '%s' %d}", names[tok.tokenType], tok.value, tok.pos)
}

type stateFn func(*lexer) stateFn

type lexer struct {
	start  int 
	pos    int
	input  string
	tokens chan token
	state  stateFn
}
// go one position forward
func (l *lexer) next() (val string) {
	if l.pos >= len(l.input) {
		l.pos++
		return ""
	}

	val = l.input[l.pos : l.pos+1]

	l.pos++

	return
}
// go one position back while scanning 
func (l *lexer) backup() {
	l.pos--
}

func (l *lexer) peek() (val string) {
	val = l.next()

	l.backup()

	return
}
// emit is used to send the token to l.tokenbs channel
func (l *lexer) emit(t tokenType) {
	val := l.input[l.start:l.pos]
	tok := token{val, l.start, t}
	l.tokens <- tok
	l.start = l.pos
}

func (l *lexer) tokenize() {
	for l.state = lexData; l.state != nil; {
		l.state = l.state(l)
	}
}

func lexData(l *lexer) stateFn {
	v := l.peek()
	
	switch {
	case v == "":
		l.emit(tokenEof)
		return nil
	// case primitiveRegexp.MatchString(v):
	// 	return lexPrimitive
	case punctuationRegexp.MatchString(v):
		return lexPunctuation

	case whitespaceRegexp.MatchString(v):
		return lexWhitespace
	case reloperatorRegexp.MatchString(v):
		return lexRelop
	case operatorRegexp.MatchString(v):
		return lexOperator
	case specialRegexp.MatchString(v):
		return lexSpecial

	case numberRegexp.MatchString(v):
		return lexNumber
	case newlineRegexp.MatchString(v):
		return lexNewline

	}

	return lexWord
}
// for matching primitive token regex so as to label as primitive token
func lexPrimitive(l *lexer) stateFn {
	matched := primitiveRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenPrimitiveTypes)

	return lexData
}
// for matching punctuation token regex so as to label as punctuation token
func lexPunctuation(l *lexer) stateFn {
	matched := punctuationRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenPunctuation)

	return lexData
}
// for matching relop token regex so as to label as relop token
func lexRelop(l *lexer) stateFn {
	matched := reloperatorRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenRelop)

	return lexData
}
// for matching whitespace token regex so as to label as whitespace token
func lexWhitespace(l *lexer) stateFn {
	matched := whitespaceRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenWhitespace)

	return lexData
}
// for matching word token regex so as to label as word(variable) token
func lexWord(l *lexer) stateFn {
	matched := wordRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenWord)

	return lexData
}
// for matching special token regex so as to label as special token
func lexSpecial(l *lexer) stateFn {
	matched := specialRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenSpecial)

	return lexData
}
// for matching operator token regex so as to label as operator token
func lexOperator(l *lexer) stateFn {
	matched := operatorRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenOperator)

	return lexData
}
// for matching Number Constant regex so as to label as Number constant token
func lexNumber(l *lexer) stateFn {
	matched := numberRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenNumber)

	return lexData
}
func lexNewline(l *lexer) stateFn {
	matched := numberRegexp.FindString(l.input[l.pos:])
	l.pos += len(matched)
	l.emit(tokenNewline)

	return lexData
}
func newLexer(input string) *lexer {
	return &lexer{0, 0, input, make(chan token), nil}
}
var flag int =0

func recursivedescentParser(m map[int]token_data) {
for i:=0;i<len(m);{
	if m[i].values[0]==tokenWhitespace {
		
		i++
	}
	if m[i].values[1]=="int" {
		
		i++
	}
	E(&i,m)
	r:=0
	if flag==1 {
		break;
	}
	if m[i].values[1]=="," {
		r=1

	}
	//fmt.Println(i)
	
	
	//fmt.Println(flag)
	//fmt.Println(m[i].values[1])
	i++
	if m[i].values[0]==tokenWhitespace {
		
		i++
	}
	
	if m[i].values[0]==tokenEof {
		
		if r==1 {
		
		//fmt.Println("wow")
		flag=1
	}

		break;
	}

}
}
func E(index *int,m map[int]token_data ){
	 if m[*index].values[1]=="," || m[*index].values[1]==";" || flag==1{
	 	//fmt.Println("E() Flag",flag)
	 	flag=1
        return
            }
     if m[*index].values[0]==tokenEof {
     	flag=1;
     	//fmt.Println("E() Flag",flag)
     	return
     }
     if m[*index].values[0]==tokenWhitespace{
	*index++} 
     //  fmt.Println("In E()",*index)    
       check(index,m,0);
      Tprime(index,m,0);
}
 
func Tprime(index *int,m map[int]token_data,val int){
 if m[*index].values[1]=="," || m[*index].values[1]==";" || flag==1{
 	//fmt.Println("TPrime() Flag",flag)
	 	
        return
            }
if m[*index].values[0]==tokenEof {
     	flag=1
     	return
     }
	if m[*index].values[0]==tokenWhitespace {
		(*index)++
            
      }
    //  fmt.Println("In TPrime()",*index)  
      if m[*index].values[0] == tokenOperator || m[*index].values[0] == tokenRelop {
      	*index++
      	check(index,m,val+1)
      	
        Tprime(index,m,val+1)
     }
      // }
}
 
func check(index *int,m map[int]token_data,val int ){
	//fmt.Println("In check()",*index,m[*index].values[1])  
if m[*index].values[1]=="," || m[*index].values[1]==";" || flag==1{
	//fmt.Println("Check() Flag()",flag)
	 	flag=1
        return
            }
     if m[*index].values[0]==tokenEof {
     	flag=1;
     	//fmt.Println("E() Flag",flag)
     	return
     }
	if m[*index].values[0]==tokenWhitespace {
		*index++
      }
      if m[*index].values[0]==tokenEof {
      flag=1
      return}
      if m[*index].values[1]=="," {
      return}
      if m[*index].values[1]==";"{
      return}
      if m[*index].values[0]==tokenWord {
            *index++
      }else if m[*index].values[0]==tokenNumber && val>0 {
            *index++
            
      }else if m[*index].values[1]== "("{
      	//fmt.Println("In check()",*index,m[*index].values[1]) 
            *index++
            E(index,m)
           
            if m[*index].values[1] == ")" {
                  *index++
            }else{
                  flag = 1;
                  return  
            }
      }else{
            flag = 1;
      }
}
 
func main() {
	var m=make(map[int]token_data)
	
	reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter String to Parsed: ")
text, _ := reader.ReadString('\n')
	//text:="int i=2,y=3   ;  "
lex := newLexer(text)
fmt.Println(text)
	go lex.tokenize()
    index:=0
	for {
		tok := <-lex.tokens
		fmt.Println(tok)
		m[index]=token_data{[2]interface{}{tok.tokenType,tok.value}}
		index++;
		if tok.tokenType == tokenEof {
			break
		}
		
	}

	recursivedescentParser(m)
	if flag == 1{
		fmt.Println("Failure. The above statement is not correct.")
	}else{
		fmt.Println("Success. The above statement is correct.")
	}
	
}
