Functions

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

Call by Value

- Immutable copied argument into the parameter
- `Advantage` => encapsulation
- `Disadvantage` => copying time

---

Call by Reference (Pointer)

- Mutable pointer
- `Advantage` => copying time
- `Disadvantage` => encapsulation

---

Parameters

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

First-Class Values

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

Variable Number of Arguments

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

Deferred Function Calls

-     func main(){
        defer fmt.Println('Bye')

        fmt.Println('Hello')
      }

- Arguments are evaluated immediately
- Call are deferred

---

OOP

- `Encapsulation` => protecting data exposing only public methods
- ``

---

Classes (alike)

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

- Pointer Receivers
-     func (self *Point) OffsetX(val float64){
        self.x += v
      }
- Best practices
  - all methods have pointer receivers
  - non have pointer receivers
