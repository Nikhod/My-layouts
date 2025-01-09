package other_practices

import "fmt"

type Human struct {
	name string // неэкспортированное поле
	Age  int    // экспортированное поле
}

// Экспортированная функция для создания нового объекта Human
func NewHuman(name string, age int) *Human {
	return &Human{name: name, Age: age}
}

// Экспортированная функция
// для получения имени - GETTER
func (p *Human) GetName() string {
	return p.name
}

// Экспортированная функция
// для установки имени - SETTER
func (p *Human) SetName(name string) {
	p.name = name
}

/*
	 предположим что эта функция была
	 ининциализирована в другом пакете......

		Тонкости и детали:
		---Скрытие данных: В примере выше поле name является неэкспортированным, то есть оно скрыто от других пакетов. Поле Age

экспортируется и доступно извне.

	-Интерфейсы и методы: Методы GetName и SetName позволяют управлять доступом к скрытым данным.
	-Пакеты: В Go инкапсуляция достигается за счет разделения кода на пакеты и четкого определения того, что
	экспортируется из каждого пакета.

-Абсолютно верно! Если поле "name" не экспортируется (начинается с маленькой буквы), то оно будет недоступно для прямого
доступа из других пакетов. Однако, вы можете создать экспортированные методы (начинающиеся с заглавной буквы), такие
как GetName и SetName, чтобы косвенно получить доступ к этому полю и изменить его.

	***Название методов должны быть с большой буквы, если надо использовать в других пакетах
*/
func FirstEncapsulationPractice() {
	person := NewHuman("Alice", 30)
	age := person.Age
	name := person.GetName()
	fmt.Printf("old data about person: %v\t%v\n", age, name)

	person.SetName("Bob")
	person.Age = 44
	name = person.GetName()

	fmt.Printf("the new data about person: %v\t%v\n", age, name)
}
