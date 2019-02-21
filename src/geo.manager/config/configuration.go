package config

import (
	coreConfig "core/config"
)

type Configuration struct {
	App     coreConfig.App     // Информация о приложении
	Logging coreConfig.Logging // Настройки логгирования
	Micro   coreConfig.Micro   // Данные для инициализации сервиса

	Server      Server      // Данные для запуска сервера
	Database    Database    // Работа с БД
	Subscrption Subscrption // Настройки подписки
}
