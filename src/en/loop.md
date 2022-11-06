## Using goroutine on loop iterator variables

### Case 1
When interating in Go, one might attempt to use goroutine to process values in parallel. For example, one might write like this:

```go
{
  	for _, val := range values {
		go func() {
        fmt.Println(val)
      }()
    }
}
```

The code above might not do as you expect because the `val` variable is actually a same single variable that takes on the value of each slice element. Because the closures are all only bound to that one variable, there is a very good chance that when you run you code you will see the last element printed for every iteration instead of each value in sequence. Because the goroutine will probably not begin executing until after the loop.

There are several ways to correct the above code:

correct 1:

```go
{
  for _, val := range values {
		val := val
		go func() {
			fmt.Println(val)
		}()
	}
}
```


correct 2:

```go
for _, val := range values {
  go func(val int) {
    fmt.Println(val)
  }(val)
}
```

correct 3:

```go
	for i := range values {
		val := values[i]
		go func() {
			fmt.Println(val)
		}()
	}
```

You can play the example code here: [../code/loop/pf_1.go](../code/loop/pf_1.go)

### Case 2

This case is similar to Case 1, but is more unnoticable.

```go
for _, val := range values {
	go val.MyMethod()
}

func (v *val) MyMethod() {
	fmt.Println(v)
}
```

The above example also will print the last element of values, the reason is the same. `Val` is the same single variable that takes on the value of each iteration. 
One can fix it using the same methods as Case 1.

You can play the example code here: [../code/loop/pf_2.go](../code/loop/pf_2.go)

## Use reference to iterator variable

### Case 1

In Go, the loop variable is a single variable that takes different values in each loop iteration. This is very efficient, but might lead to unintended behavior when used incorrectly.

For example, see the following program:

```go
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
```

It will output unexpected results: The printed values are all 3 and the printed Addresses are all the same.

Why? 

In each iteration, we append the address of `i` to the `out` slice, but since it is the same variable, we append the same address which eventually contains the last value that was assigned to `i`. One of the solution is to copy the loop variable, in this case which is `i`, into a new variable:

```go
	var out []*int
	for i := 0; i < 3; i++ {
		i := i //Copy i into a new variable.
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
```

You can play the example code here: [../code/loop/pf_3.go](../code/loop/pf_3.go)


### Case 2

While the case 1 example might look a bit obvious, the same unexpected behavior could be more hidden in some other cases. For example, the loop variable can be an array and the reference can be a slice:

```go
	var out [][]int
	for _, i := range [][1]int{{1}, {2}, {3}} {
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
```

On each iteration, `i` is the same variable of array type. `i[:]` would return the slice of array, since slice inner pointer is pointing to the address of `i`, which remains the same through the loop body. Though On each iteration, the values of `i` changes, but the appended slice to `out` points to the same array variable `i`, so the printed element of values is always the last value of `i`.

You can play the example code here: [../code/loop/pf_4.go](../code/loop/pf_4.go)