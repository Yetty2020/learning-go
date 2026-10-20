package main
// a stuct with no field takes 0 bytes
// The Ts comparison of this is a Set
//Go has no built-in Set. Instead you use a map whose values are empty structs:

/*
seen := map[string]struct{}{}
seen["Fatihah"] = struct{}{}*/
//You only care about the keys, so the value is “nothing,” and it costs 0 bytes. In TS terms it’s like Record<string, true>, except the true takes up no space.
// anonymous empty struct type
// empty := struct{}{}

// named empty struct type
// type emptyStruct struct{}
// namedEmpty := emptyStruct{}