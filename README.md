- efficient compiling (not take minutes to compile a code when you want to test)
- efficient execution
- easy to code

- strong, static type system (everything has a type)
- C inspired syntax
- compiled (we need to compile the code to execute)
- multi paradigm (object oriented, struct programming, imperative programming. It all up to you)
- garbage collected (process that deletes data that not using anymore. No one is pointing at memory space, so it's deleted)
- creates a single binary compilation (when compile, creates one file. One executable output for code)
- new versions will be always fully backwards compatible(new version is compatible with the previous versions..).


- every file must be within a package(folder)
- the project must contains an entry point, where the code will init and run (package main and function main)


- it can generate executable binary files for different platforms and operating systems
    - it can compile to WebAssembly (WASM, is ability of the web platform, like browsers, to execute native code that can be written from differents languages)
        Ex: games, criptocurrencies, Artificial inteligence
    - it can transpile to frontend javascript (GopherJS)
- multi platform (creating app for a platform, need to pick both DOS and the architecture)

- Go (Golang) GOOS and GOARCH https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

- Common use cases
    - Web services
    - Web applications
    - DevOps
    - Desktop UI
    - Machine Learning
    ...


Language types

    Interpreted Languages --> Source Code (javascript, python..)

    Intermediate Compiled Languages --> Bytecode (Java, C#)

    Compiled languages --> Machine Code (code execute directly in one CPU)


syntax

    Code blocks in {}
    No styling freedom
    does not need semi-colon (;) unless you need. To separate sentences
        Ex: print("hello"); print("another print in the same line")
    Case sensitive language
    Strongly typed
    No object oriented

Rules

    One file act as the entry point with a main function

    folder is a package (pasta é um pacote)

    within a file
        Variables
        functions
        Type declarations
        Methods declarations

Modules and CLI

    Module is a project! is a group of packages
    it contains go.mod file with configuration and metadata

    CLI manipulates the module
        go mod init
        go build
        go run 
        go test
        go get


Workspaces and CLI

    is a new kind of super module app

    it contains a go.work file with configuration and metadata including which module to use

    CLI manipulates the workspaces

        go work init

    workspaces--
        modules--
            packages--


Data types goes after identifier
Variables have nil by default
Constants can be only bool, string or numbers

Immutable variable x constant

    Immutable variable is a varible, so it is an a space alocate in memory
    immutable was set and no one can not change the value

    constant: The compiler finds a constant reference, it goes to the definition copy the value and paste the value
    "copy and paste"
    it's not in a memoty, it is a fix value

variable with initialization
    var name string
    name = "John"

or

initialization shortcut 
    otherName := "Mike"


Types
    string
    numbers values: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
        uint - unsigned int - only positive numbers
    floating values: float32, float64
    bool
        == != < > <= >= && || !
    pointers
        pointer is we have a variable, and then we have another variable that is pointing to the original one.


    
    numbers

        id := 4
        price := 45.4

        priceAsInt := int(price)
        idAsFloat := float32(id)


    Strings

        str1 := "Just a text"

        str2 := "Every text
            is multiline
            by default"


    Collections

        Arrays: fixed length
            [5]int

        Slices: similar to a dynamic length array,
        but they are actually chunck of arrays
            []int

        Maps: key/value dictionaries
            map[keytype]valuetype

        Generics

Visibility
    variables, constants, functions..

    if it's camelCase = private to the package

        func notExported() {}

    if it's TitleCase = public and exported to others packages

        func Exported() {}

Variables and lambda functions can be:

    - Module Scoped
    - Function Scoped
    - Block Scoped

Functions
    
    - It can receive arguments
    - arguments can have default values
    - in Go, functions can return more than one value
    - can return labeled(rotulado) variables
    - receive arguments always by value
        pass arguments as copy
        when pass a value as an argument, it not pass
        the reference to the variable that has the value
        you are cloning the value, a copy of the value
    
    Functions receiving references

        receive pointer instead of the value

        func increment (x *int) {
            *x++
        }

        func main() {
            i := 1
            increment(&i)
        }

Control structeres

    if - else

        if user != nil {
        } else {
        }

        you can create variables that will be just used inside the if scope
        first expression - variable (message)
        second expression - condition (user != nil)

        if message := "hello"; user != nil {
        } else {
        }

    switch (reloaded)

        switch day {

            case "Monday":
                fmt.Println("It's Monday")
            case "Saturday":
                fmt.Println("It's Saturday")
            case "Sunday":
                fmt.Println("It's Sunday")
            default:
                fmt.Println("It's another day")
        }

        switch with any condition with any variable
        ex: can be used to verify parameters when the backend receive datas
        switch {

            case user != nil:
                ..
            case user.active == false:
                ..
            case user.deleted == true:
                ..
            default:
        }

    for

        for i:=0; i<len(collection); i++ {
        }

        for range, similar to "for in" in JS
        for index := range collection {
        }

        for range, similar to "foreach"
        for key, value := range map {
        }

        endGame := false
        for endGame {
            // process game loop
            // like while
        }

        for count < 10 {
            count += 1;
        }

        for {
            // infinite loop
        }


compiler

    does not matter for the compiler the name of the go files
    what matters is the name of the package

Conversion to string

    use fmt.Sprintf to make a custom string, it will return a string
    like when you want to send a string customized

    ex:
    fmt.Sprintf("%.2f", value) // value will print just 2 decimal after the comma
    fmt.Sprintf("Hello %v how are you doing this %v?", name, day)


Type definition

    // type alias
    type definition = float32

    // new type based on
    type distance float32

    // types have a constructor and conversion functions
    d := distance(34.5)

    METHOD X FUNCTION

    // Function - just receive argument, or not, and does not hava attribute argument
    func ToKm(miles distance) distanceKm {
        return distanceKm(1.60934 * miles)
    }

    // Method - has an attribute argument
    func (miles distance) ToKm() distanceKm {
        return distanceKm(1.60934 * miles)
    }


Complex types for definitions

    Structures
        - kind replace the class idea
        - data type with strongly typed properties
        - have default constructor
        - can be added methods to it

    Interfaces
        - definition of methods
        - emulate polimorsphism
        - implicit implementation
        - can embed interfaces in other interfaces

GoRoutines

    go way of using threads
    invoke any function with a go prefix
        go func whatever()

    can communicate through channels, an special
    type of variable

    A channel contains a value of any kind

    A routine can define a value for a channel
    and other routine can wait for that value

    Channels can buffered or not

  Using Channels:
    
    Ex 1:
        var m1 chan string

        m2 := make(chan string)

        m2 <- "hello"

        messagem := <- m2

    Ex 2 (buffer channel):
        // 2 - how many of those values you wanna send over the channel
        logs := make(chan string, 2)
        // sending two messages
        logs <- "hello"
        logs <- "world"
        fmt.Println(<-messages)
        fmt.Println(<-messages)

    __Channel has two goals__:
        
        1. send and receive data between GoRoutines
        2. use it for one routine to wait for another routine
            it's kind of an event listener between goroutines

    TO AVOID DEADLOCKS:
        close the channels before ending the program with close(chan)