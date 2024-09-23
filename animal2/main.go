package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Интерфейс для всех животных
type Animal interface {
	MakeSound() (string, error) // Возвращаем ошибку для обработки
	Move() string
	GetInfo() string
	GetType() string
}

// Интерфейс для животных, которые умеют плавать
type Swimmer interface {
	Swim() string
	CanSwim() bool
}

// Собака
type Dog struct {
	Name string
}

func (d Dog) MakeSound() (string, error) {
	return "Гав-гав!", nil
}

func (d Dog) Move() string {
	return "Собака бегает"
}

func (d Dog) GetInfo() string {
	return fmt.Sprintf("Собака %s - верный друг!", d.Name)
}

func (d Dog) GetType() string {
	return "Млекопитающее"
}

func (d Dog) Bark() string {
	return "Гав-гав-гав!"
}

func (d Dog) CanSwim() bool {
	return true
}

func (d Dog) Swim() string {
	return "Собака плавает лапами"
}

// Кот
type Cat struct {
	Name string
}

func (c Cat) MakeSound() (string, error) {
	return "Мяу!", nil
}

func (c Cat) Move() string {
	return "Кот ходит мягко"
}

func (c Cat) GetInfo() string {
	return fmt.Sprintf("Кот %s - пушистый и игривый!", c.Name)
}

func (c Cat) GetType() string {
	return "Млекопитающее"
}

func (c Cat) Meow() string {
	return "Мяу-мяу!"
}

func (c Cat) CanSwim() bool {
	return false
}

// Лошадь
type Horse struct {
	Name string
}

func (h Horse) MakeSound() (string, error) {
	return "И-го-го!", nil
}

func (h Horse) Move() string {
	return "Лошадь бежит быстро"
}

func (h Horse) GetInfo() string {
	return fmt.Sprintf("Лошадь %s - грациозное животное!", h.Name)
}

func (h Horse) GetType() string {
	return "Млекопитающее"
}

func (h Horse) Neigh() string {
	return "И-го-го-го!"
}

func (h Horse) CanSwim() bool {
	return true
}

func (h Horse) Swim() string {
	return "Лошадь плавает мощными движениями"
}

// Курица
type Chicken struct {
	Name string
}

func (c Chicken) MakeSound() (string, error) {
	return "Куд-кудах!", nil
}

func (c Chicken) Move() string {
	return "Курица ходит"
}

func (c Chicken) GetInfo() string {
	return fmt.Sprintf("Курица %s - несушка!", c.Name)
}

func (c Chicken) GetType() string {
	return "Птица"
}

func (c Chicken) Cluck() string {
	return "Куд-кудах-кудах!"
}

func (c Chicken) CanSwim() bool {
	return false
}

// Верблюд
type Camel struct {
	Name string
}

func (c Camel) MakeSound() (string, error) {
	return "Р-р-р!", nil
}

func (c Camel) Move() string {
	return "Верблюд ходит медленно"
}

func (c Camel) GetInfo() string {
	return fmt.Sprintf("Верблюд %s - корабль пустыни!", c.Name)
}
func (c Camel) GetType() string {
	return "Млекопитающее"
}

func (c Camel) Grunt() string {
	return "Р-р-р!"
}

func (c Camel) CanSwim() bool {
	return false
}

// Неизвестное животное -  генерирует ошибку при вызове MakeSound
type UnknownAnimal struct {
	Name string
}

func (u UnknownAnimal) MakeSound() (string, error) {
	return "", fmt.Errorf("Неизвестное животное не может издавать звуки")
}

func (u UnknownAnimal) Move() string {
	return "Движется неизвестным образом"
}

func (u UnknownAnimal) GetInfo() string {
	return fmt.Sprintf("Неизвестное животное %s", u.Name)
}

func (u UnknownAnimal) GetType() string {
	return "Неизвестный тип"
}

func (u UnknownAnimal) CanSwim() bool {
	return false
}

// Запись ошибок в файл логов
func logError(err error) {
	f, err := os.OpenFile("errors.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия лог-файла:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%sn", err))
	if err != nil {
		fmt.Println("Ошибка записи в лог-файл:", err)
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите тип животного (собака, кот, лошадь, курица, верблюд, неизвестное) или 'exit' для выхода:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		var animal Animal

		switch input {
		case "собака":
			fmt.Println("Введите имя собаки:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и новую строку
			animal = Dog{Name: name}
		case "кот":
			fmt.Println("Введите имя кота:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и новую строку
			animal = Cat{Name: name}
		case "лошадь":
			fmt.Println("Введите имя лошади:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и новую строку
			animal = Horse{Name: name}
		case "курица":
			fmt.Println("Введите имя курицы:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и новую строку
			animal = Chicken{Name: name}
		case "верблюд":
			fmt.Println("Введите имя верблюда:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и новую строку
			animal = Camel{Name: name}
		case "неизвестное":
			fmt.Println("Введите имя неизвестного животного:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name) // Удаляем пробелы и новую строку
			animal = UnknownAnimal{Name: name}
		default:
			fmt.Println("Неверный тип животного. Попробуйте снова.")
			continue
		}

		// Вывод информации о животном
		fmt.Println(animal.GetInfo())

		// Вывод звука
		sound, err := animal.MakeSound()
		if err != nil { // Обработка ошибок при вызове MakeSound
			fmt.Println("Ошибка:", err)
			logError(err)
			continue
		}
		fmt.Println("Звук:", sound)

		// Вывод движения
		fmt.Println("Движение:", animal.Move())

		// Вывод типа
		fmt.Println("Тип:", animal.GetType())

		// Проверка умения плавать
		if swimmer, ok := animal.(Swimmer); ok {
			fmt.Println("Умеет плавать:", swimmer.CanSwim())
			fmt.Println("Способ плавания:", swimmer.Swim())
		}

		fmt.Println("----")
	}
}
