### Что выведет данная программа и почему?
```
func main() {
  n := 0
  if true {
    n := 1
    n++
  }
  fmt.Println(n)
}
```

Ответ:
```
sad
```