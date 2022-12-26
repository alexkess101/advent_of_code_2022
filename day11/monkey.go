package main

type Monkey struct {
	items     []int
	operation func(old int) int
	condition func(lvl int) (int, int)
	count     int
}

func NewMonkey(items []int, operation func(old int) int, condition func(lvl int) (int, int)) *Monkey {
	return &Monkey{
		items:     items,
		operation: operation,
		condition: condition,
		count:     0,
	}
}

func (this *Monkey) run(monkeys []*Monkey) {
	for _, item := range this.items {
		worryLevel := this.operation(item)
		this.count++
		monkeyToThrowTo, newWorryLevel := this.condition(worryLevel)
		this.throw(newWorryLevel, monkeys[monkeyToThrowTo])
	}
}

func (this *Monkey) throw(item int, monkey *Monkey) {
	this.items = this.items[1:]

	monkey.items = append(monkey.items, item)
}

func Condition(lvl int, div int, trueOption int, falseOption int) (monkey, newValue int) {
	newWorryLevel := lvl / 3
	if newWorryLevel%div == 0 {
		return trueOption, newWorryLevel
	} else {
		return falseOption, newWorryLevel
	}
}

func NoMoreMonkeys(rounds int, monkeys []*Monkey) {
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.run(monkeys)
		}
	}
}
