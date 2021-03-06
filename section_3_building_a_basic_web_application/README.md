# Section 3 - Building a basic Web Application

## 3.1 - Как работают Web-приложения и requestresponse lifecicle:

* Сайт хранится на сервере.
* www. - name server, который возвращает ip адрес сервера, на котором хрантся сайт.
* На этот ip адрес клиентом отправляется запрос (request) на показ домашней страницы сайта (в который может входить дополнительная информация, например - куки, хранящие информацию о текущем пользователе).
* Сервер, на котором есть база данных, кэш, приложения и тд. формирует домашнюю страницу и возвращает её в ответе (response), в который так же могут входить куки для идентификации пользователя.
* В таком режиме (request-response lifecicle) и работают web-приложения.
* Сервера строятся на Service Oriented Architecture (SOA).

## 3.2 - Web-приложение Hello, world:

* Система маршрутизации позволяет сопоставить определенные запросы с определенными ресурсами внутри веб-приложения. Для создания протейшей системы маршуртизации в приложении может применяться функция HandleFunc().
* [Ссылка на статью.](https://metanit.com/go/web/1.2.php)
* http.HandleFunc() - функция из стандартной библеотеки http, позволяющая зарегестрировать функцию, которая будет исполняться при переходе по указанному адресу.
* http.ResponseWriter - поток ответных данных.
* http.Request - информаци о запросе.
* localhost - локальный сервер
* :8080 - порт.

## 3.3 - Функции-обработчики

* Выносить handlers-функции в отдельные функции - отличный способ разгрузить main-функцию.
* const - ещё один способ объявления переменной. Такая переменная не изменяема: 
    ```
    const myVar = "Hello"
    ```


## 3.4 - Обработка ошибок

* Функции в Go могут возвращать несколько значений:
    
    ```
        func Addition(x, y int) (int, err) {
            // Тело функции
        }
    ```
* Это часто используется для обработки ошибок, когда второе возвращаемое значение - error или nil.

## 3.5 - Отображение html-шаблонов

* В Go есть встроенный функционал для отображения html-шаблонов (templates)
    
    ```
    /// Парсинг шаблона
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)

	/// Отображение шаблона
    err = parsedTemplate.Execute(w, nil)
    ```

## 3.7 - Рефакторинг кода и добавление стиля для шаблонов

* Для добавления дополнительного стиля в html шаблоны использовали Bootstrap

## 3.8 - Изменение структуры кода

* Для использования директорий и иморта кода из других файлов нужно инициализировать модуль [ссылка на урок.](../../learning-go/section_2_overview_of_the_go_language/2.11_packages/main.go)
* main.go помещаем в cmd/web/
* Вызываемые пакеты - в pkg/...
* handlers.go в pkg/handlers/
* render.go в pkg/render/
