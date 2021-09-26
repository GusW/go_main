## Functions

- `Reusability`
- `Abstraction`
- `Single responsability`
  - cohesion
  - only one operation
- `Understandability`
  - find logic quickly
  - trace through data
- `Low Complexity`
  - Function length
  - Control flow
    - little conditionals
    - `Partional Conditionals` => a conditional calls another function

---

## Call by Value

- Immutable copied argument into the parameter
- `Advantage` => encapsulation
- `Disadvantage` => copying time

---

## Call by Reference (Pointer)

- Mutable pointer
- `Advantage` => copying time
- `Disadvantage` => encapsulation

---

## Parameters

- Passing `Array Pointers`

  -     func foo(x *[3]int) int{
          (*x)[0] ++
        }

        func main() {
          a := [3]int{1,2,3}
          foo(&a)
          fmt.Print(a) // {2,2,3}
        }

- Passing `Slice Values`

  - contains a pointer to the original array
  - can pass by value and mutate the original array instead of original reference
  -     func foo(sli []int) int{
          (*x)[0] ++
        }

        func main() {
          a := []int{1,2,3}
          foo(a)
          fmt.Print(a) // {2,2,3}
        }

---

## First-Class Values

- Functions can be treated as other types

  - can be created dinamically
  - can be passed as arguments
  - can be returned by other functions
  -     var funcVar func(int) int

        func incFn(x int) int {
          return x++
        }

        func main(){
          funcVar = incFn
          fmt.Print(funcVar(1))
        }

  -     func applyFn(someFn func(int) int,
            val int) int {
              return someFn(val)
        }

---

## Variable Number of Arguments

- `...` => ellipsis
-     func maxVal(vals ...int) int {
        for idx, val := range vals {
          //...
        }
      }
- `Variadic Slice Argument`
-     valSlice := []int{1,2,3,4,5,7}
      fmt.Print(maxVal(valSlice...))

---

## Deferred Function Calls

-     func main(){
        defer fmt.Println('Bye')

        fmt.Println('Hello')
      }

- Arguments are evaluated immediately
- Call are deferred

---

## OOP

- `Encapsulation` => protecting data exposing only public methods

---

## Classes (alike)

- Associating methods with data
- Call by value
- Implicit method argument `self`
-     type Point struct { x, y float64 }

      func (self Point) DistToOrig() float64 {
        t := math.Pow(self.x, 2) + math.Pow(self.y, 2)
        return math.Sqrt(t)
      }

      func main(){
        p := Point(3, 4)
        fmt.Print(p.DistToOrig())
      }

- Controlling access

  - `Public access` => Capital letter
  -     package data

        x := 1

        func PrintX() { fmt.Print(x) }

  -     package main
        import "data"

        func main() {
          data.PrintX()
        }

---

## Pointer Receivers

-     func (self *Point) OffsetX(val float64){
        self.x += v
      }
- Best practices
  - all methods have pointer receivers
  - non have pointer receivers

---

## Interfaces

- Set of `method signatures`
- Express conceptual similarities between types
-     type Shape2D interface {
        Area() float64
        Perimeter() float64
      }

      type Triangle { ... // whatever comes in here }

      func (t Triangle) Area() float64 { ... }
      func (t Triangle) Perimeter() float64 { ... }

---

## Concrete vs Interfaces Types

- Concrete

  - Exact representation of data and methods
  - Complete method implementation

- Interface

  - Specify some method signatures
  - Implementations abstracted

- Interface Values

  - can be assigned to variables
  - passed, returned
  - have 2 components:
    - dynamic type => concrete type which it is assigned to
    - dynamic value => value of the dynamic type
  - `Interface value` = (dynamic type, dynamic value)
  -     type Speaker interface {Speak() string}

        type Dog struct {name: string}

        func (d Dog) Speak() string {fmt.Print(d.name)}

        func main() {
          var s1 Speaker
          d1 := Dog('Toby') // dynamic type => Dog
          s1 = d1           // dynamic value => d1
          s1.Speak()
        }

  - May have `nil` dynamic value if equals type Pointers
  -     var s1 Speaker
        var d1 *Dog
        s1 = d1 // has dynamic type (Dog) but no dynamic value (is a pointer...)

---

## Interface uses

- Need a function which takes multiple types of parameters
-     type Shape2D interface {
        Area() float64
        Perimeter() float64
      }

      type Triangle { ... // whatever comes in here }

      func (t Triangle) Area() float64 { ... }
      func (t Triangle) Perimeter() float64 { ... }

      type Rectangle { ... }

      func (r Rectangle) Area() float64 { ... }
      func (r Rectangle) Perimeter() float64 { ... }

      func FitInYard(s Shape2D) bool {
        if s.Area() < 100 && s.Perimeter < 100 { // random numbers...
          return true
        }
        return false
      }

- `Empty Interface` => specifies no methods. All types satisfy it
-     func PrintMe(val interface{}){
        fmt.Println(val)
      }

---

## Type Assertions

- Exposing type differences => concrete type, not dynamic
- Used to disambiguation
-     func DrawShape(s Shape2D) {
        rect, ok := s.(Rectangle)
        if ok {
          DrawRect(rect)
        }

        tri, ok := s.(Triangle)
        if ok {
          DrawTri(tri)
        }
      }

- `Type Switch`
  -     func DrawShape(s Shape2D) {
           switch := sh := s.(type){
             case Rectangle:
               DrawRect(sh)
             case Triangle:
               DrawTri(sh)
           }
        }

---

## Error Interface

-     type error interface{
        Error() string
      }
