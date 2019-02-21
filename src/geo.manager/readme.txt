Пинг-проверка работы приложения:
	http://localhost:8002/api/ping/

Регистрация Города (/api/city/add_city/:name):
	http://localhost:8002/api/city/add_city/Moscow

---

Разворачивание:
	В новую диру копируем:
		geo.manager - файл приложения
		conf.yml - файл конфига приложения

	Переходим в диру:
		cd /root/geo
	
	Правим файл конфига приложения для раоты с базйо и прочим:
		conf.yml
	
	Устанавливаем приложение как системный сервис:
		./geo.manager --service install

	Запускаем приложением как системный сервис:
		service geo.manager start

	Смотрим состояние приложения как системного сервиса:
		service geo.manager status

	Если сервис нормально запустился, то смотрим его работу:
		curl "http://localhost:8002/ms/geo.manager/v1/ping/" -v

	Смотрим в браузере адрес:
		http://localhost:8002/ms/geo.manager/v1/ping/

---

Работа с сервисом:
	Запускаем приложением как системный сервис:
		service geo.manager start

	Смотрим состояние приложения как системного сервиса:
		service geo.manager status

	Остановка приложениея как системный сервис:
		service geo.manager stop

	Приложение в режиме системного сервиса пишет логи:
		/var/log/geo.manager.log - системный лог
		/var/log/geo.manager.err - лог ошибок
