# Pinger

Для получения информации о контейнерах я использую Docker API, но чтобы он сам работал изнутри докера, необходимо прокидывать docker.socket внутрь. Это несколько понижает безопасность решения, но ничего лучше я не придумал.

Пинг происходит через 5 секунд после завершения предыдущего, таймаут самого пинга также составляет 5 секунд, пингуются все доступные контейнеры сразу (через го-рутины). В очень крупных масштабах может возникнуть проблема с потреблением памяти - в таком случае стоит перейти на очередь.

Клиент для Backend API сгенерирован через openapi-generator.
