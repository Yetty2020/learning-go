### Step 1: Decide the shared behaviour

Ask yourself: what should all these types be able to **do**? Write it in plain words first.

> "Anything that can speak and give back a string."

### Step 2: Write the interface

Turn that sentence into the checklist: just the method name, `()` and the return type. There's no `func`, no receiver and no body.

```go
type speaker interface {
	speak() string
}
```

### Step 3: Write the first struct

Ask: what **data** does this type hold?

```go
type dog struct {
	name string
}
```

### Step 4: Give that struct the method

Copy the interface line **exactly**, then add `func`, a receiver in front, and a body:

```
interface line:            speak() string
add func + receiver:  func (d dog) speak() string
add body:             func (d dog) speak() string { return ... }
```

The name, brackets and return type must match the interface character for character.

### Step 5: Repeat steps 3–4 for each other type

`cat`: its own struct, the same method line, its own body.

### Step 6: Write the function that uses the interface

Its parameter type is the **interface name**, not a struct:

```go
func introduce(s speaker) {
	// only call methods listed in the interface
}
```

### Step 7: Use it in `main()`

Create each struct and pass it in:

```go
introduce(dog{name: "Rex"})
```

### Summary to keep beside you

```
1. Behaviour in words
2. Interface     → type X interface { method() T }
3. Struct        → type Y struct { fields }
4. Method        → func (y Y) method() T { body }
5. Repeat 3–4 for each type
6. Function      → func use(v X) { v.method() }
7. main()        → use(Y{...})
```

Pause and drill it now. Ten minutes of writing it yourself now will make the next lessons much easier. You understand the idea; the syntax just needs reps.

### The only 3 pieces of syntax

Every interface example is built from these three, always in this order:

**① The interface (the checklist)**

```go
type NAME interface {
	method() returnType
}
```

**② A struct (the data)**

```go
type TYPE struct {
	field fieldType
}
```

**③ A method on that struct (passing the checklist)**

```go
func (x TYPE) method() returnType {
	return ...
}
```

Plus a function that **uses** the interface:

```go
func doThing(v NAME) {
	v.method()
}
```

Here's the same thing in TS, side by side:

| Go | TS |
|---|---|
| `type Speaker interface { speak() string }` | `interface Speaker { speak(): string }` |
| `type dog struct { name string }` | `type Dog = { name: string }` |
| `func (d dog) speak() string { ... }` | `speak() { ... }` inside the object or class |
| `func greet(s Speaker)` | `function greet(s: Speaker)` |

### Drill: write it yourself, without copying

Make a scratch folder, `ch07-interfaces/practice/`, and write a `main.go` that does this:

1. An interface **`speaker`** with one method, `speak()`, that returns a `string`.
2. A struct **`dog`** with a `name string` field, and a `speak()` method that returns `name + " says woof"`.
3. A struct **`cat`** with a `name string` field, and a `speak()` method that returns `name + " says meow"`.
4. A function **`introduce(s speaker)`** that prints `s.speak()`.
5. In `main()`, call `introduce` once with a `dog` and once with a `cat`.

Run it with `go run .`, and the output should be:

```
Rex says woof
Tom says meow
```

**Rules:**

- Use the 3-piece template above, but **type** everything yourself.
- If you get stuck, write the part you *can*, and paste it with `// stuck here`.
- Let `go build` errors guide you.

Paste your code and the output, then we'll commit lesson 01 and move on.