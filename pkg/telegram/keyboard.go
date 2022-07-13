package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var Menu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Поиск тура", "https://relax62.ru/poisk-tura/"),
		tgbotapi.NewInlineKeyboardButtonURL("Горящие туры", "https://relax62.ru/goryashchie-tury/"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Группа ВКонтакте", "https://vk.com/travel_relax62"),
		tgbotapi.NewInlineKeyboardButtonURL("Канал Telegram", "https://t.me/travel_relax62"),
	),
)
