### Что выведет данная программа и почему?
```
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```

Ответ:
```
Выводит: 100 2 3 4 5. Так как слайс v указывает на тот же участок памяти, что слайс a, по этому у нас меняется 0 элемент с 1 на 100.
После добавления нового элемента с помощью append() слайс v уже не указывает на слайс a, 
а создает новый, по этому в оригинальном слайсе не добавляется элемент.
```