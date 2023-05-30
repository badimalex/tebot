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
		names, err := m.chatStore.SelectSubscribes(chatID)
		if err != nil {
			return err.Error()
		}
		if len(names) > 0 {
			res := "Вы подписаны на следующие товары:\n"
			for _, item := range names {
				res += item + "\n"
			}
			return res
		}
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
