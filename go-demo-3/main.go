package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap{}
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()

		switch variant {
		case 4:
			break Menu
		case 3:
			DeleteBookmark(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 1:
			printBookmarks(bookmarks)
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладку")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string

	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkKey)

	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)

	bookmarks[newBookmarkKey] = newBookmarkValue

}

func DeleteBookmark(bookmarks bookmarkMap) {
	var BookmarkKeyToDelete string

	fmt.Print("Введите название: ")
	fmt.Scan(&BookmarkKeyToDelete)

	delete(bookmarks, BookmarkKeyToDelete)

}
