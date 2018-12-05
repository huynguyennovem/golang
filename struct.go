package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}
func printBook( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book inheritance : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}
func main() {
	var Book1 Books    /* Declare Book1 of type Book */

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* print Book1 info */
	printBook(Book1)

}
