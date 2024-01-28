package main // Пакет main - стандартная точка входа для исполняемых Go программ.
//import "net/http"
//
//import (
//	"fmt"      // Импорт пакета fmt для форматированного ввода/вывода.
//	"net/http" // Импорт пакета net/http для работы с HTTP.
//)
//
//type User struct {
//	name                 string
//	age                  uint16
//	money                int16
//	avg_grades, hapiness float64
//}
//
//// здесь мы обращаемся к объекту User
//// (u User) - объект и их него вызывается метод getAllInfo()
//func (u User) getAllInfo() string {
//	return fmt.Sprintf("User name is: %s. He is %d and he has money "+
//		"equal: %d", u.name, u.age, u.money)
//}
//func (u User) setNewName(newName string) {
//	u.name = newName
//}
//
//// chelikBob - обработчик
//func chelikBob(w http.ResponseWriter, r *http.Request) {
//	// bob- новый экземпляр структуры User
//	bob := User{"Bob", 25, -50, 4.2, 0.8}
//	//bob.name = "Alex"
//	//тут мы вызываем функцию getAllInfo, куда зашит объект bob
//	bob.setNewName("Alex")
//	fmt.Fprintf(w, bob.getAllInfo())
//}
//
//// homePage - обработчик
//func homePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Go is nice!")
//}
//
//// homeContatcs - обработчик
//func homeContatcs(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Пошел на хуй отсюда, мудила!")
//}
//func handleRequest() {
//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/home-contacts", homeContatcs)
//	http.HandleFunc("/chel", chelikBob)
//	http.ListenAndServe(":6520", nil)
//}
//
//func main() {
//	//var bob User
//	//bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, hapiness: 0.8}
//	handleRequest()
//}
