### Чем отличаются буферизированные и не буферизированные каналы?

Ответ:
```
Не буффиризированные - для синхронной работы рутин, после записи блокируется, пока с него не произойдет считывание.
Буффиризированные - для асинхронной работы рутин, храня данные в буфере, 
блокируется только если он будет заполнен полностью или пуст
```