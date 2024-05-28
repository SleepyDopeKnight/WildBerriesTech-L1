### Что выведет данная программа и почему?
```
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(wg sync.WaitGroup, i int) {
      fmt.Println(i)
      wg.Done()
    }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```

Ответ:
```
Выводит: fatal, deadlock. Потому что передаем копию, а не указатель, по этому основной счетчик не уменьшается
```