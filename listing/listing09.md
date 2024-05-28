### Сколько существует способов задать переменную типа slice или map?

Ответ:
```
var some_slice []int
some_slice := []int{1, 2, 3}
some_slice := make([]int)

var some_map map[int]int
some_map := map[int]int{1: 1, 2: 2, 3: 3}
some_map := make(map[int]int)
```