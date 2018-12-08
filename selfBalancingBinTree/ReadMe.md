# ━━(￣ー￣*|||━━

Slice has a trick to make them like vector.

https://github.com/golang/go/wiki/SliceTricks

But there is a problem.

If I pass my slice to another function during a function call, and change my slice size in that function i.e. 
```go
slice = append(slice,make(type,1)...)
```

So the append function likely will create a new underlying array and return it to `slice`. 

Which means, `slice` are not the `slice` anymore, and you just **cannot** have this new `slice` in original function unless  you return it. what the heck? ? ! !

This is same as https://stackoverflow.com/questions/10740153/problems-about-slice-and-append-in-go

Do not have workaround currently. I only can append it first then pass it.