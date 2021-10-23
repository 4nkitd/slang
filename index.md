# slang 🐕‍🦺

goal was to learn how a interpreter works,  in other works who does these programing languages i use on daily basis works behind the seen  how a variable is assigned it's value and other things in the same realm.

## why ![](https://golang.org/lib/godoc/images/go-logo-blue.svg) ?

before this i mainly worked with scripting language other than java. the gofmt and binary file output are the main reasons though.

## can i use this language ?

### NO ☠️

this barely even compare to the stepping stones of a daily driven language like go or js. it exists only so that someone can read and understand how a language works under the hood.

---<!-- markdownlint-capture -->

### Main Components of a language

- Token list
- Tokenizer ( lexer )
- Parser
- Compiler

### How to RUN 🎽

```bash
    #file is a flag , use it to run your file
    slang -file ./sample/sample.so
```
OR

```bash
    slang 
    # running slang in your terminal without any flags gives you a console to work with.
```

### ⛷️ Supported Datatype  

- Boolean
- Integer
- String

### How to define a variable
```
let x = 42;
```

### How to if/else 

    Does not support elseif

```
        if (x == 1) {
            return true;
        } else {
            return false;
        }
```

### How to use a function
```
let fibonacci = fn (x) {
    if (x == 0) {
        return 0;
    } else {
        if (x == 1) {
            return 1;
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

fibonacci(15);
```
