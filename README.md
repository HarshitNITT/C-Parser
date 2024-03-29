# C-Parser
This is a combination of a **lexer** and a **recursive descent parser** which is a **top-down parser** where the lexer converts the input string to the stream of tokens and recursive descent parser is used to determine the correctness of the statements.

# Installation
First clone the project in your local machine using:
~~~
git clone https://github.com/HarshitNITT/C-Parser
cd C-Parser
cd code
go run main.go
~~~
The open the go.main file and compile/run it in the terminal to get started to using the project.

# Usage 
Running the code gives the following result:

<img src="https://github.com/HarshitNITT/C-Parser/blob/master/images/demo.png" />

The First Part of the Output is the stream of tokens generated.
The Second Part of the Output is the **Recursive Descent Parser** Output. 
## Note
It can deal with arithmetic expressions and declaration(int) expressions only.
### Arithmetic Expressions:
~~~
a=b+c;b=(c+d);f=(y+10)/2;
~~~
### Declaration Expressions:
~~~
int a=b+c,b=(c+d); int f=(y+10)/2;
~~~
# Languages
Golang
# License
This project is licensed under the MIT License.
