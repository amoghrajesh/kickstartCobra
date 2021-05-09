# kickstartCobra

A simple attempt to build a CLI tool using Cobra coded in Golang.

The binary can be run with various input args and basically does addition of 'any' number of numbers (int or float) based on flags. 
There is also an option to add only the even numbers and odd numbers.

Steps to run - 

1. `git clone https://github.com/amoghrajesh/kickstartCobra` just outside the $GO_PATH/$GOPATH directory so that package matching is simpler.
2. `cd kickstartCobra`
3. `go install mycli` 
4. `cd ../../bin` or `cd $HOME/go/bin`
5. To run the most basic command, do this - 
   
    a. `./mycli` - Just displays a message on the screen
    
    b. `./mycli hello` - Greets with a Hello message
   
    c. `/mycli hello 1 2 3 4` - Returns sum of all numbers provided
   
    d. `/mycli hello even 1 2 3 4 5 6` - Returns sum of all the even numbers present
   
    e. `/mycli hello 1.2 3.4 5.6 --isfloat` - Usage of this flag returns a floating-point sum of the numbers provided as input

###Note
The repo is setup with a post commit hook to build the project and run the main_test.go test cases on the repo. Do not edit the main_test.go file.