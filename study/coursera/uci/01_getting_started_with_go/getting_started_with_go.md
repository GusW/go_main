## Create new

- $ `go mod init` <path/to/folder>

---

## Workspaces & Packages

- `src` - source code files
- `pkg` - libs
- `bin` - executables

---

## Go Tool

- `go build` => compiler
- `go doc` => docs for a package
- `go fmt` => formatter
- `go get` => package manager (donwloader)
- `go list` => list local packages
- `go run` => compiles and run
- `go test` => testing '\*\_test.go' files

---

## Variables

- `var year, month, day int`
- `var year int = 2021`
- `var month = 1` => int type infered
- `day := 2` => shorthand declartion, type infered, in-function only

---

## Type Declarations

- `type Celsius float64`
- `var temp Celsius`
- `type IDNum int`
- `var pid IDNum`

---

## Pointers

- address to data in memory
- `&<pointer>` => returns the address
- `*<pointer>` => returns the data (dereferencing)

1. `var x int = 1`
2. `var y int`
3. `var ip *int`
4. `ip = &x` => ip references x
5. `y = *ip` => y sets value to be equal ip e.g. 1

---

## new()

- creates variables and returns pointer
- initialised as 0 (zero)

1. `ptr := new(int)`
2. `*ptr = 3`

---

## Blocks

- `{}`
- Universe block => all Go source
- Package block
- File block
- Lexical Scoping
  - How variable references are resolved
  - `bi >= bj` if `bj` is defined inside of `bi`

---

## Stack & Heap

1. `Stack`

   - memory area primariarly dedicated to function calls
   - normally deallocated after function completes

2. `Heap`
   - persistent memory chunk
   - deallocated manually (in most compiled languages)
     - `x=malloc(32);`
     - `free(x);`

---

## Garbage Collection

- Although in the example below the `x := 1` is kept in the Stack memory the pointer reference is returned to the main function thus no deallocation is performed
-     func foo() *int {
          x := 1
          return &x
      }

      func main() {
          var y *int
          y = foo()
          fmt.Printf("%d", *y)
      }

---

## Comments

- `// single line comment`
- `/* block comment */`

---

## Print

- `import "fmt"`
-     x := "yo"
      fmt.Prinft("Hi %s", x)

---

## Integers

- Signed bytes => `int8`, `int16`, `int32`, `int64`
- Unsigned bytes => `uint8`, `uint16`, `uint32`, `uint64`
  - can get bigger due to extra bit not used to sign

---

## Type Conversions

-     var x int8 = 1
      var y int16 = 2
      y = x // throws an error once types must match

      y = int16(x)

---

## Floating Point

- `float32` => ~ 6 digits of precision
- `float64` => ~ 15 digits of precision

---

## Strings

- `ASCII` => 8-bit (American)
- `Unicode` => 32-bit
  - `UTF8` => variable length 8-bit that matches ASCII and which can be expanded to 32-bit
- `Code points` => unicode characters (2^32 code points)
- `Rune` => a code point in Go ('A' => Rune of 0x41)
- Strings are basically arrays of runes
- `Unicode` package
  - `IsDigit(r rune)`
  - `IsSpace(r rune)`
  - `IsLetter(r rune)`
  - `IsLower(r rune)`
  - `IsPunct(r rune)`
  - `ToUpper(r rune)`
  - `ToLower(r rune)`
- `Strings` package
  - `Compare(s1, s2)`
  - `Contains(s, substring)`
  - `HasPrefix(s, prefix)`
  - `Index(s, substring)`
  - `Replace(s, old, new, n)` => first `n` instances
  - `ToUpper(s)`
  - `ToLower(s)`
  - `TrimSpace(s)`
- `Strconv` package

  - `Atoi(s)` => ASCII to Integer
  - `Itoa(i)` => Integer to String
  - `FormatFloat(f, fmt, prec, bitSize)` => Float to string
  - `ParseFloat(s, bitSize)` => String to float

---

## Constants

-     const x = 1.32
      const (
        y = 100
        z = "hej"
      )

- `iota` => enum
  -      type Grades int
         const (
           A Grades = iota
           B
           C
           D
           E
           F
         )

---

## Control Flows

- Conditional

  -      if <condition> {
            <consequent>
         }

- For loops

  -      for <init>; <cond>; <update> {
            <statements>
         }

- Switch/Case

  -      switch x {
          case 1:
            fmt.Printf("1")
          case 2:
            fmt.Printf("2")
          default:
            fmt.Printf("no case")
         }
  - unlike another languages it `break` automatically after case is found
  - tagless:
    -     switch {
          case x > 1:
            fmt.Printf("1")
          case x <= 1:
            fmt.Printf("2")
          default:
            fmt.Printf("no case")
      }

- `break` - ceases control flow
- `continue` - skips iteration

---

## Scan

- Captures manually input data
- Takes a pointer as an arg
  -      var appleNum int
         fmt.Printf("How many apples?")
         num, err := fmt.Scan(&appleNum)
         fmt.Printf(appleNum)

---

## Arrays

- Fixed length
-     var myList [5]int
      myList[0] = 2
      fmt.Print(myList[1])
      // Array literal
      var literal [5]int = [5]{1,2,3,4,5}
      // or
      literal := [...]{1,2,3,4,5}

- Loops
-      x := [...]{1,1,2,3,5,8,13,21}
       for idx, val := range x {
         fmt.Printf("Index %d, Value %d", idx, val)
       }

---

## Slices

- slice of array
- variable size, up to the size of the array
- `Pointer` => slice start point
- `Length len()` => size
- `Capacity cap()` => max size (`end of array - pointer`)
-      array := [...]{10,20,30,40,50,60}
       s1 := array[1:4]
       s2 := array[2:5]

       // or literal
       sli := []int{1,2,3}
       /* in this case an array is crated
          and a slice is pointed to that array
          covering from arr[0] to arr[-1]
        */

        // or without predefined values
        // (type, len/cap)
        sliceFromMake = make([]int, 10)

        // or without predefined values
        // (type, len, cap)
        sliceFromMake = make([]int, 10, 15)

- `append()` => adds els to the end of the slice
  - inserts into underlying array
  - increases size of array if necessary
  -     sli = append(sli, "whatever")

---

## Hash Table

- Unique keys
- `Hashing function` => process the slot for the key
- Advantages
  - constant lookup O(1) x O(N) in arrays
  - arbitrary keys
- Disadvantages
  - may have collisions (two keys into same slot)
- `Maps`
-     var idMap map[string][int]
      idMap = make(map[string]int)

      // or literal
      idMap := map[string]int{"joe": 123}

      // delete key
      delete(idMap, "joe")

      // unpacking
      value, isPresent := idMap["joe"]

      // looping
      for key, value := range idMap {
        fmtPrint(key, value)
      }

---

## Structs

- `class`
-      type Person struct {
         name string
         addr string
         phone string
       }

       var p1 Person
       p1.name = "Jane"

       // initilizing with 0-ish values
       p1 := new(Person)

       // struct literal
       p1 := Person(name: "Jane", addr: "Rosetta Av", phone: "dunno")

---

## Packages (RFCs/protocols)

- `net/http`=> web. `http.Get(www.whatever.com)`
- `net` => tcp/ip `net.Dial("tcp", "www.hoho.com:8081"`
- `encoding/json` => JSON manipulation
  - json `marshalling` => transform go object into json object
  -     p1 := Person(name: "joe", pass: "123")
        // returns JSON representation as []byte
        barr, err := json.Marshal(p1)
  - json `unmarshalling` => transform json object in go struct
  -     var p2 Person
        err := json.Unmarshal(barr, &p2)

---

## Files

- Linear access, not random (mechanical delay, begin -> end)
- `io/ioutil` => basic ops
  - Open => ...
  - Read => bytes into []byte. opens and closes out of the box
  -     dat, err:= ioutil.ReadFile("something.txt")
  - Write => write []byte into file, creating it
  -     dat = "hey joe"
        err := ioutil.WriteFile("output.txt", dat, 0777) // unix-like perms
  - Close => ...
  - Seek => move read/write head
- `os` => better package
  - `os.Open()` => returns a file descriptor(File)
  - `os.Close()` => ...
  - `os.Read()` => reads from a file into a []byte
    - can control the amount read by the size of byte
    -     f, err := os.Open("test.txt")
          barr := make([]byte, 10)
          nb, err := f.Read(barr)
          f.Close()
  - `os.Write()` => writes a []byte into a file, controlled size as well
    -     f, err := os.Create("output.txt")
          barr := []byte{1,2,3}
          nb, err := f.Write(barr)
          // or
          nb, err := f.WriteString("Hi")
