package pointers

import "fmt"

// Category представляет структуру категории
type Category struct {
	name          string
	subcategories []*Category
}

// AddSubcategory Добавление подкатегории
func (c *Category) AddSubcategory(sub *Category) {
	c.subcategories = append(c.subcategories, sub)
}

// FindCategory Рекурсивный поиск категории по имени
func (c *Category) FindCategory(name string) *Category {
	if c.name == name {
		return c
	}
	for _, sub := range c.subcategories {
		if found := sub.FindCategory(name); found != nil {
			return found
		}
	}
	return nil
}

// ChangeName Изменение имени категории
func (c *Category) ChangeName(newName string) {
	c.name = newName
}

// PrintTree Вывод дерева категорий
func (c *Category) PrintTree(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(c.name)
	for _, sub := range c.subcategories {
		sub.PrintTree(level + 1)
	}
}

func TreeExample() {
	// Создаем корневую категорию
	electronics := &Category{name: "Электроника"}

	// Добавляем подкатегории
	laptops := &Category{name: "Ноутбуки"}
	smartphones := &Category{name: "Смартфоны"}
	smartphones.AddSubcategory(&Category{name: "iPhone"})
	electronics.AddSubcategory(smartphones)
	electronics.AddSubcategory(laptops)

	// Добавляем подкатегорию в ноутбуки
	ultrabooks := &Category{name: "Ультрабуки"}
	laptops.AddSubcategory(ultrabooks)

	// Поиск категории
	found := electronics.FindCategory("Ультрабуки")
	if found != nil {
		fmt.Println("Найдена категория:", found.name)
	}

	// Изменение имени категории
	if found != nil {
		found.ChangeName("Ультратонкие ноутбуки")
	}

	// Вывод структуры дерева
	fmt.Println("Структура категорий:")
	electronics.PrintTree(0)
}
