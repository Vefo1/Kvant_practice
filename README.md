Go Прокси API для предсказания показателей здоровья
Это Go API, которое служит прокси для внешнего API предсказания различных показателей здоровья (HBA1C, LDLL, FERR, LDL, TG, HDL). Оно принимает данные в GET-параметрах URL, преобразует их в JSON-формат и пересылает POST-запросом во внешний API, а затем возвращает ответ внешнего API клиенту. Проект разработан с использованием структурированного подхода, включающего конфигурационный файл, логирование и промежуточный обработчик для авторизации.

Тестирование
Прокси-API требует авторизации через заголовок, указанный в config.yml (app_auth.header_name, по умолчанию X-App-Auth), со значением, указанным в config.yml (app_auth.token).

Общие параметры для всех моделей (если применимо):

uid (string, фиксировано "web-client")

age (int)

gender (int)

1. HBA1C (/predict/hba1c)
Параметры: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/hba1c?age=30&gender=1&rdw=12.5&wbc=7.2&rbc=4.5&hgb=14.0&hct=42.0&mcv=90.0&mch=30.0&mchc=33.0&plt=250.0&neu=60.0&eos=2.0&bas=0.5&lym=30.0&mon=7.0&soe=10.0&chol=200.0&glu=90.0"




2. LDLL (/predict/ldll)
Параметры: chol, hdl, tg

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/ldll?age=20&gender=1&chol=1&hdl=2&tg=3"




3. FERR (/predict/ferr)
Параметры: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, crp

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/ferr?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&crp=7"




4. LDL (/predict/ldl)
Параметры: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/ldl?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=7&glu=8"




5. TG (/predict/tg)
Параметры: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/tg?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=7&glu=8"




6. HDL (/predict/hdl)
Параметры: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/hdl?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=6&hct=5&mcv=4&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=77&glu=8"




Как реализовано
Данное API реализовано на языке программирования Go с использованием стандартной библиотеки net/http и следует стандартным рекомендациям по структуре Go-проектов.

Назначение основных компонентов и их директорий:
config/:

Содержит файлы и код, отвечающие за загрузку и управление конфигурацией приложения (например, порты, URL внешних API, токены авторизации, уровни логирования). Это позволяет легко настраивать приложение без изменения исходного кода.

internal/models/:

Определяет структуры данных (Go struct), которые используются для представления сущностей в приложении, в частности, для форматирования данных, отправляемых во внешние API, и их получения. Обеспечивает строгую типизацию и упрощает работу с JSON.

internal/interfaces/:

Содержит определения интерфейсов Go. Эти интерфейсы описывают контракты для бизнес-логики и других компонентов, позволяя отделить абстракцию от конкретной реализации. Это способствует модульности, тестируемости и гибкости архитектуры.

internal/services/:

Реализует бизнес-логику приложения. Здесь находятся службы, которые выполняют основные операции, такие как взаимодействие с внешними API, обработка данных и применение бизнес-правил. Сервисы реализуют интерфейсы, определенные в internal/interfaces/.

internal/handler/:

Содержит HTTP-обработчики (хендлеры), которые принимают входящие HTTP-запросы, парсят параметры, вызывают соответствующие методы сервисов для выполнения бизнес-логики и формируют HTTP-ответы клиенту.

cmd/app/:

Является точкой входа в приложение. Этот пакет отвечает за инициализацию всех компонентов (конфигурации, логгера, сервисов, обработчиков), применение промежуточных обработчиков (middleware) и запуск HTTP-сервера.

Поток выполнения запроса:
Клиент отправляет GET запрос на http://localhost:8080/predict/<model_name> (например, /predict/ldll) с параметрами в URL и заголовком X-App-Auth.

Запрос попадает в main.go, который передает его AuthMiddleware.

AuthMiddleware проверяет X-App-Auth токен. Если он неверный, запрос отклоняется с 401 Unauthorized.

Если токен действителен, AuthMiddleware передает запрос соответствующему хендлеру (например, LdllPredictHandler).

Хендлер (через handlePredictionRequest) парсит параметры из URL, валидирует их и формирует соответствующую структуру models.XxxPredictRequest.

Хендлер вызывает соответствующий метод (PredictXxx) у PredictServiceImpl.

Метод PredictXxx (через makePredictionRequest) формирует POST-запрос к внешнему API (например, https://apiml.labhub.online/api/v1/predict/ldll), добавляет заголовки (включая токен для внешнего API) и отправляет его.

Внешний API обрабатывает запрос и возвращает ответ.

PredictServiceImpl получает ответ и возвращает его в хендлер.

Хендлер передает ответ (статус-код и тело) обратно клиенту.

Все этапы логируются с помощью pkg/logger.
