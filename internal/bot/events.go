package bot

type Events struct {
	chatStore *Storage
}

func NewEvents(chatStore *Storage) *Events {
	return &Events{chatStore: chatStore}
}

func (m *Events) HandleMessage(messageText string, chatID int64) string {
	switch messageText {
	case "/start":
		return "Привет, напишите название продукта который вы хотите искать:"
	case "Стоп":
		// сделать выборку если есть searches с текущим chatID то дописать к До новых встреч
		// Вы будете подписаны на следующие товары
		// 1. ываыва
		// 2.
		return "До новых встреч!"
	case "Нет":
		return "Вы не подписались на уведомления. \nВведите название товара для поиска либо напишите Стоп чтобы выйти"
	case "Да":
		m.chatStore.Subscribe(chatID)
		return "Вы успешно подписались на уведомления. \nВведите название товара для поиска либо напишите Стоп чтобы выйти"
	default:
		m.chatStore.AddItemSearchInTable(messageText, chatID)
		return "Хотите получать уведомления о новых товарах? Да/Нет"
	}
}
