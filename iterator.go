package main

import "fmt"

type Aggregate interface {
	iterator() *Iterator
}

type Iterator interface {
	hasNext() bool
	next() interface{}
}

type Book struct {
	name string
}

func (book *Book) getName() string {
	return book.name
}

type BookShelf struct {
	books []*Book
	last  int
}

func (shelf *BookShelf) iterator() *BookShelfIterator {
	bookShelfIterator := &BookShelfIterator{bookShelf: shelf}
	return bookShelfIterator
}

func (shelf *BookShelf) bookShelf(maxsize int) {
	shelf.books = make([]*Book, maxsize)
}

func (shelf *BookShelf) getBookAt(index int) *Book {
	return shelf.books[index]
}

func (shelf *BookShelf) appendBook(book *Book) {
	shelf.books[shelf.last] = book
	shelf.last++
}

func (shelf *BookShelf) getLength() int {
	return shelf.last
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func bookShelfIterator(bookshelf *BookShelf) *BookShelfIterator{
	return &BookShelfIterator{
		bookShelf: bookshelf,
		index:     0,
	}
}

func (i *BookShelfIterator) hasNext() bool {
	if i.index < i.bookShelf.getLength() {
		return true
	}
	return false
}
func (i *BookShelfIterator) next() interface{} {
	book := i.bookShelf.getBookAt(i.index)
	i.index++
	return book
}

func main() {
	var bookshelf = new(BookShelf)
	bookshelf.bookShelf(4)
	bookshelf.appendBook(&Book{name: "１冊目の本"})
	bookshelf.appendBook(&Book{"Number 2"})
	bookshelf.appendBook(&Book{"33333"})
	bookshelf.appendBook(&Book{"Forth"})
	var it Iterator = bookshelf.iterator()
	for it.hasNext() {
		book := it.next().(*Book)
		fmt.Println(book.getName())
	}
}
