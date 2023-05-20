# Коля и дата-центры
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
>
> **CodeRun** 325. Коля и датацентры Средняя

| Язык                   | Ограничение времени | Ограничение памяти | Ввод                           | Вывод                            |
| ---------------------- | ------------------- | ------------------ | ------------------------------ | -------------------------------- |
| Все языки              | 2.5 секунд          | 512Mb              |                                |                                  |
| OpenJDK 17 + json      | 4 секунды           | 512Mb              | стандартный ввод или input.txt | стандартный вывод или output.txt |
| C# (MS .NET 6.0 + ASP) | 4 секунды           | 512Mb              |                                |                                  |
| Python 3.11.2          | 4 секунды           | 512Mb              |                                |                                  |

Рано или поздно все крупные IT-компании создают свои дата-центры. Коля только устроился в такую компанию и еще не успел во всем разобраться. В его компании есть $N$ дата-цетров,в каждом дата-центре установлено $M$ серверов.

Из-за большой нагрузки серверы могут  выключаться. Из-за спешки при постройке дата-центров включить только один сервер не получается, поэтому приходится перезагружать весь дата-центр. У каждого дата-центра есть два неотрицательных целочисленных параметра: $R_i$ — число перезапусков $i$-го дата-центра и $A_i$ — число рабочих (не выключенных) серверов на текущий момент в $i$-м дата-центре.

Коля получил задачу по сбору некоторых  метрик, которые в будущем позволят улучшить работу дата-центов. Для этого Коля собрал данные о $Q$ событиях, произошедших за текущий день. Коля справился с этой задачей, но просит помочь и проверить свои результаты.

## Формат ввода

В первой строке входных данных записано 3 положительных целых числа $n$, $m$, $q$ ($1 \leq q \leq 105$, $1 \leq n \cdot m \leq 106$) — число дата-центров, число серверов в каждом из дата-центров и число событий соответственно.

В последующих $q$ строках записаны события, которые могут иметь один из следующих видов:

`RESET` $i$ — был перезагружен $i$-й дата-центр ($1 \leq i \leq n$)

`DISABLE` $i$ $j$ — в $i$-м дата-центре был выключен $j$-й сервер ($1 \leq i \leq n$, $1 \leq j \leq m$)

`GETMAX` — получить номер дата-центра с наибольшим произведением $R_i \cdot A_i$

`GETMIN` — получить номер дата-центра с наименьшим произведением $R_i \cdot A_i$

## Формат вывода

На каждый запрос вида `GETMIN` или `GETMAX` выведите единственное положительное целое число — номер дата-центра, подходящий под условие. В случае неоднозначности ответа  выведите номер наименьшего из дата-центров.

### Пример 1

**Ввод**

```
3 3 12
DISABLE 1 2
DISABLE 2 1
DISABLE 3 3
GETMAX
RESET 1
RESET 2
DISABLE 1 2
DISABLE 1 3
DISABLE 2 2
GETMAX
RESET 3
GETMIN
```

**Вывод**

```
1
2
1
```

### Пример 2

**Ввод**

```
2 3 9
DISABLE 1 1
DISABLE 2 2
RESET 2
DISABLE 2 1
DISABLE 2 3
RESET 1
GETMAX
DISABLE 2 1
GETMIN
```

**Вывод**

```
1
2
```



## Примечания

Обратите внимание на 2 пример. DISABLE приходится для уже выключенного сервера. В данном случае сервер по-прежнему остаётся выключенным.

> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> <details>
> На го ты умираешь писать CUMпоратор
> <summary>Python</summary>
>
> ```py
> import sys
> import heapq
> 
> input = sys.stdin.readline
> n, m, q = map(int, input().split())
> db = [0] * n  # битовая маска включенных машин
> da = [m] * n  # число включённыых машин
> dr = [0] * n  # число перезапусков
> dw = [0] * n  # метрика r*a
> 
> min_heap = [(0, i) for i in range(n)]
> max_heap = [(0, i) for i in range(n)]
> 
> for _ in range(q):
>     cmd, *a, = input().split()
>     *a, = map(int, a)
>     if cmd == "RESET":
>         i = a[0] - 1
>         db[i] = 0
>         da[i] = m
>         dr[i] += 1
>         dw[i] = dr[i] * da[i]
> 
>         heapq.heappush(min_heap, (dw[i], i))
>         heapq.heappush(max_heap, (-dw[i], i))
>     elif cmd == "DISABLE":
>         i, j = a
>         i -= 1
>         j -= 1
>         t = 1 << j
>         
>         if db[i] & t: 
>             continue
>         db[i] |= t
>         da[i] -= 1
>         dw[i] -= dr[i]
>         
>         heapq.heappush(min_heap, (dw[i], i))
>         heapq.heappush(max_heap, (-dw[i], i))
>     elif cmd == "GETMAX":
>         while -max_heap[0][0] != dw[max_heap[0][1]]:
>             heapq.heappop(max_heap)
>         print(max_heap[0][1] + 1)
>     elif cmd == "GETMIN":
>         while min_heap[0][0] != dw[min_heap[0][1]]:
>             heapq.heappop(min_heap)
>         print(min_heap[0][1] + 1)
> ```

# Средняя сетевая задержка
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
>
> **CodeRun** 331. Средняя сетевая задержкаы Средняя

Рассмотрим древовидную сеть хостов для обработки запросов. Запрос будет обрабатываться следующим образом:

1. Запрос приходит на корневой хост (пусть он называется `balancer.test.yandex.ru`);
2. хост формирует подзапросы на хосты-потомки (не более одного запроса в потомок);
3. хост ждет получения всех ответов;
4. хост формирует свой ответ и возвращает его;

Все потомки обрабатывают подзапросы по такой же схеме. Все хосты логируют события ввиде следующих записей:

- `datetime` – время наступления события;
- `request_id` – id запроса;
- `parent_request_id` – id родительского запроса (для корневого бэкенда `NULL`);
- `host` – имя бэкенда, на котором возникло событие;
- `type` – тип события (список указан ниже);
- `data` – описание события.

Допустимые типы событий:

- `RequestReceived` – на бэкенд поступил новый запрос (поле `data` пустое);
- `RequestSent` – на бэкенд-потомок был отправлен подзапрос (в поле `data` записывается имя бэкенда-потомка);
- `ResponseSent` – бэкенд отправил ответ родителю (`data` пустое);
- `ResponseReceived` – бэкенд получил ответ от потомка (в поле `data` записываются имя бэкенда-потомка и статус – `OK` или `ERROR` –, разделенные символом табуляции).

Все записи собираются в единую базу данных.

Хосты не идеальны, поэтому им требуется время на пересылку запросов и получение ответов. Определим время выполение запроса к корневому хосту, как сумму разниц `datetime` между всеми соответствующими парами событий `RequestSent`/`RequestReceived` и `ResponseSent`/`ResponseReceived`, которые возникли при обработке запроса. Найдите эту величину, усредненную по запросам на корневой хост.

## Формат ввода

Используется БД postgresql 10.6.1 x64.

Система перед проверкой создает таблицу с событиями следующим запросом:

```sql
CREATE TABLE requests (  
    datetime TIMESTAMP,  
    request_id UUID,  
    parent_request_id UUID,  
    host TEXT,  
    type TEXT,  
    data TEXT  
);
```

После таблица заполняется тестовыми данными.

## Формат вывода

Напишите `SELECT` выражение, которое вернет таблицу из одной строки с колонкой `avg_network_time_ms` типа `numeric`, в которую будет записана средняя сетевая задержка в миллисекундах.

**Внимание!** Текст выражения подставится в систему как подзапрос, поэтому завершать выражение точкой с запятой не надо (в противном случае вы получите ошибку `Presentation Error`).

## Примечания

Для таблицы `requests` с таким содержимым (здесь для компактности пишем числа вместо `UUID`’а и миллисекунды в `datetime`, в проверочной таблице будут `UUID`’ы и `timestamp`’ы):

| **datetime** | **request_id** | **parent_request_id** | **host**                | **type**         | **data**               |
| ------------ | -------------- | --------------------- | ----------------------- | ---------------- | ---------------------- |
| .000         | 0              | NULL                  | balancer.test.yandex.ru | RequestReceived  |                        |
| .100         | 0              | NULL                  | balancer.test.yandex.ru | RequestSent      | backend1.ru            |
| .101         | 0              | NULL                  | balancer.test.yandex.ru | RequestSent      | backend2.ru            |
| .150         | 1              | 0                     | backend1.ru             | RequestReceived  |                        |
| .200         | 2              | 0                     | backend2.ru             | RequestReceived  |                        |
| .155         | 1              | 0                     | backend1.ru             | RequestSent      | backend3.ru            |
| .210         | 2              | 0                     | backend2.ru             | ResponseSent     |                        |
| .200         | 3              | 1                     | backend3.ru             | RequestReceived  |                        |
| .220         | 3              | 1                     | backend3.ru             | ResponseSent     |                        |
| .260         | 1              | 0                     | backend1.ru             | ResponseReceived | backend3.ru      OK    |
| .300         | 1              | 0                     | backend1.ru             | ResponseSent     |                        |
| .310         | 0              | NULL                  | balancer.test.yandex.ru | ResponseReceived | backend1.ru      OK    |
| .250         | 0              | NULL                  | balancer.test.yandex.ru | ResponseReceived | backend2.ru      OK    |
| .400         | 0              | NULL                  | balancer.test.yandex.ru | ResponseSent     |                        |
| .500         | 4              | NULL                  | balancer.test.yandex.ru | RequestReceived  |                        |
| .505         | 4              | NULL                  | balancer.test.yandex.ru | RequestSent      | backend1.ru            |
| .510         | 5              | 4                     | backend1.ru             | RequestReceived  |                        |
| .700         | 5              | 4                     | backend1.ru             | ResponseSent     |                        |
| .710         | 4              | NULL                  | balancer.test.yandex.ru | ResponseReceived | backend1.ru      ERROR |
| .715         | 4              | NULL                  | balancer.test.yandex.ru | ResponseSent     |                        |

запрос участника должен возвращать следующий результат:

| **avg_network_time_ms** |
| ----------------------- |
| 149.5                   |

Тут два корневых запроса. Выпишем времена, которые прошли между отправкой запроса/ответа и его получением.

Запрос с id 0:

- `balancer.test.yandex.ru` -> `backend1.ru` – $50$ мс (от $.100$ до $.150$)
- `balancer.test.yandex.ru` -> `backend2.ru` – $99$ мс (от $.101$ до $.200$)
- `backend1.ru` -> `backend3.ru` – $45$ мс (от $.155$ до $.200$)
- `backend2.ru` -> `balancer.test.yandex.ru` – $40$ мс (от $.210$ до $.250$)
- `backend3.ru` -> `backend1.ru` – $40$ мс (от $.220$ до $.260$)
- `backend1.ru` -> `balancer.test.yandex.ru` – $10$ мс (от $.300$ до $.310$)

Суммарно это $50+99+45+40+40+10=284$ мс.

Запрос с id 4:

- balancer.test.yandex.ru -> backend1.ru – $5$ мс (от $.505$ до $.510$)
- backend1.ru -> balancer.test.yandex.ru – $10$ мс (от $.700$ до $.710$)

Суммарно это $5+10=15$ мс.

Итого, ответ $(284+15) / 2=149.5$
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> <details>
> <summary>SQL</summary>
>
> ```sql
> SELECT CAST(
>    (
>        (SELECT sum(
>                    1000 * EXTRACT (EPOCH FROM rr.datetime) - 1000 * EXTRACT (EPOCH FROM rs.datetime)
>                )   /   (SELECT count(DISTINCT request_id)
>                        FROM requests
>                        WHERE parent_request_id IS NULL)
>                    FROM requests rs
>                    INNER JOIN requests rr ON rs.request_id=rr.parent_request_id
>                    AND rr.host=rs.data
>                    AND rs.type='RequestSent'
>                    AND rr.type='RequestReceived'
>        ) +
>            (SELECT sum(1000 * EXTRACT (EPOCH FROM rr.datetime) - 1000 * EXTRACT (EPOCH FROM rs.datetime))/
>            (SELECT count(DISTINCT request_id)
>                 FROM requests
>                 WHERE parent_request_id IS NULL
>             )
>             FROM requests rs
>             INNER JOIN requests rr ON rr.request_id=rs.parent_request_id
>             AND rr.data like concat(rs.host, '%')
>             AND rs.type='ResponseSent'
>             AND rr.type='ResponseReceived'
>        )
>    ) AS NUMERIC
>) AS avg_network_time_ms
>```

## Сервис подписки
Необъяснимая аномалия! На серверах Яндекс Маркета отказывает оборудование: ломаются жесткие диски, плавится оперативная память, выходит из строя система охлаждения. Системные администраторы локализовали проблему — причиной поломок оказалась используемая база данных. Руководители приняли решение срочно вывести из эксплуатации упомянутую базу данных и заменить ее самописной. Вам нужно как можно скорее предоставить MVP, который поддерживает:

частичное обновление товарных предложений в базе данных
уведомление сервисов-подписчиков при обновлении данных
Товарное предложение в базе описывается следующей JSON схемой:
- частичное обновление товарных предложений в базе данных
- уведомление сервисов-подписчиков при обновлении данных

Товарное предложение в базе описывается следующей JSON схемой:
```json
{  
  "$id": "offer.schema.json",  
  "type": "object",  
  "properties": {  
    "id": {  
      "type": "string",  
      "description": "Offer identifier, only numerical symbols are allowed"  
    },  
    "price": {  
      "type": "integer",  
      "description": "Offer price, value in range from 0 to 10̂9"  
    },  
    "stock_count": {  
      "type": "integer",  
      "description": "Items left on stocks, value in range from 0 to 10̂9"  
    },  
    "partner_content": {  
      "type": "object",  
      "properties": {  
        "title": {  
          "type": "string",  
          "description": "Offer title filled in by the partner"  
        },  
        "description": {  
          "type": "string",  
          "description": "Offer description filled in by the partner"  
        }  
      }  
    }  
  },  
  "required": [  
    "id"  
  ]  
}
```
При межсервисном взаимодействии к товарному предложению добавляется контекст, который содержит идентификатор для сквозной трассировки, его схема:

```json
{  
  "$id": "message.schema.json",  
  "type": "object",  
  "properties": {  
    "trace_id": {  
      "type": "string"  
    },  
    "offer": {  
      "$ref": "offer.schema.json"  
    }  
  },  
  "required": [  
    "trace_id",  
    "offer"  
  ]  
}
```
Сервис, который отправляет запрос на обновление товарного предложения, обязательно заполняет идентификатор оффера (поле offer.id) и идентификатор для трассировки (поле trace_id). Все остальные поля в запросе опциональны. В таком случае при применении обновления будет происходить слияние полей. Например, в базе у оффера заполнены поля price=9990, и приходит обновление stock_count=100, тогда в базе будут сохранены оба поля. Гарантируется, что все входящие запросы валидны и соответствуют схеме. Так как это прототип, удаление товаров из базы и очищение полей было решено не поддерживать. Помимо хранения товарных предложений в базе, в сервисе необходима функция доставки обновлений в сервисы-подписчиков. Одна подписка включает в себя два набора полей: trigger и shipment, не обязательно листовых. Когда изменяется любое trigger поле или поле, вложенное в trigger поле, подписчику отправляется сообщение. В сообщении находятся все shipment и trigger поля этого подписчика, а также идентификаторы оффера и трассировки из запроса, который привел к этому сообщению. Инициализация поля также считается за его изменении и создает сообщение об обновлении.

Формат ввода
Первая строка входных данных содержит два целых числа n и m ( 1 ≤ n ≤ 5 0 , 1 ≤ m ≤ 1 0 , 0 0 0 ) — количество сервисов подписчиков и количество запросов на обновления. Следующие n строк содержат описания сервисов подписчиков: i -я строка содержит описание i -го подписчика. В начале строки задается a i и b i — количество trigger и shipment полей соответственно. Далее a i trigger полей, и b i shipment полей. Следующие m строк содержат запросы на обновление, каждая строка — это валидный json, удовлетворяющий схеме m e s s a g e . s c h e m a . j s o n .

Формат вывода
На каждое событие обновления выведите k j сообщений в формате m e s s a g e . s c h e m a . j s o n , где k j — это количество сервисов-подписчиков, которым данное событие интересно. Сообщения должны идти в том же порядке, что и обновления, которые привели к ним. Сообщения в рамках одного обновления должны быть отсортированы по порядковому номеру подписчика.





```py
import json


class User:

    def __init__(self, triggers: set, shipments: set):
        self.triggers = triggers
        self.shipments = shipments


class Offer:
    _offer: dict = {'price': None, 'stock_count': None, 'partner_content': {}}
    _isChange: set

    # helpers
    def updateOffer(self, upgrade_offer: dict):
        self._isChange = set()

        price = upgrade_offer.get('price')  # int | none
        stock_count = upgrade_offer.get('stock_count')  # int | none
        partner_content = upgrade_offer.get('partner_content')  # dict | none

        if price is not None and self._offer['price'] != price:
            self._offer['price'] = price
            self._isChange.add('price')

        if stock_count is not None and self._offer['stock_count'] != stock_count:
            self._offer['stock_count'] = stock_count
            self._isChange.add('stock_count')

        if partner_content is not None:
            for key in ['title', 'description']:
                value = partner_content.get(key)
                if value is not None and self._offer['partner_content'].get(key) != value:
                    self._offer['partner_content'][key] = value
                    self._isChange.add('partner_content')

    def getOffer(self, trace_id: str, offer_id: str, users):
        result = {'id': offer_id}
        for user in users:
            if self._isChange.intersection(user.triggers):
                for field in user.triggers.union(user.shipments):
                    value = self._offer.get(field)
                    if value is not None:
                        result[field] = value
                print(json.dumps({'trace_id': trace_id, 'offer': result}))


def main():
    n, m = [int(i) for i in input().split()]

    users = []  # все пользователи их может быть от 1 до 50
    offers = {}  # все предложения их может быть от 1 до 10000 

    for i in range(n):
        # входные данные a - количество trigger полей, b количество shipment полей, services список полей 
        #  конкретного пользователя
        a, b, *services = input().split() 
        
        # создаю экземпляр класса User и сую его в список users, класс User имеет 2 поля triggers и shipments
        users.append(User(set(services[0:int(a)]), set(services[int(a):-1]))) 

    for i in range(m):
        # входные данные, которые преобразуются в json
        data = json.loads(input())
        
        # если оффер не найден в offers, то мы создаем новый экземпляр класса Offers и добавляем его в offers
        offer_id = data['offer']['id']
        if offers.get(offer_id) is None:
            offers[offer_id] = Offer()
        
        # получаем оффер из offers
        offer = offers[offer_id]
        
        # обновляем поля
        offer.updateOffer(data['offer'])
        
        # уведомляем всем пользователей чью тригер поля были обновлены
        offer.getOffer(data['trace_id'], offer_id, users)


if __name__ == '__main__':
    main()
```