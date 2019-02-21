package config

import (
	coreConfig "core/config"
)

type Configuration struct {
	App     coreConfig.App     // Информация о приложении
	Logging coreConfig.Logging // Настройки логгирования
	Micro   coreConfig.Micro   // Данные для инициализации сервиса

	Subscrption Subscrption // Настройки подписки
}
