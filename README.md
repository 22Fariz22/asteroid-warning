# asteroid-warning

GET запрос возвращает общее количество астероидов на основе даты их ближайшего сближения с
Землей (neo)

Источник API - NASA (https://api.nasa.gov/) во вкладке "Asteroids NeoWs: Near Earth Obiect Web
Service", запрос "Neo - Feed"

На вход приходит 5 - 10 разных дат в формате YYYY-MM-DD

Пример запроса к сервису /neo/count?dates-2020-01-20&dates-2020-03-20&dates-2020
05-24&dates=2020-08-22&dates=2020-08-29&dates=2020-10-20&dates=2020-10-29

POST запрос записывает (или обновляет если такая дата уже есть) в БД (postgres или mysq|) в таблицу nео_count количество астероидов (пео) за определенную дату.

На вход приходит массив из дат и количества астероидов в json формате. 
Пример запроса к готовому сервису:
URN - /neo/count
Тело запроса:
{
"neo_counts": [
{
"date". "2020-01-20"
"count": 12
  },
  {
"date": "2020-02-26",
"count": 9
    }
  ]
}
