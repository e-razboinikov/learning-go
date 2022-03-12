# Section 2: Overview of the Go languge

## Variables and functions:
* Два вида объявления переменной: 

```
var myVar string - объявление переменной с явным указанием типа
```
```
myVar := "John" - type inference, компилятор сам определяет тип переменной
 ```

* Функции могут возвращать несколько значений:

```
func saySomething(something string) (string, string) {
	return something, "World"
}
```