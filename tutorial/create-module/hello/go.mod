module github.com/bert82503/golang-learn/tutorial/create-module/hello

go 1.19

replace github.com/bert82503/golang-learn/tutorial/create-module/greetings => ../greetings

require github.com/bert82503/golang-learn/tutorial/create-module/greetings v0.0.0-00010101000000-000000000000
