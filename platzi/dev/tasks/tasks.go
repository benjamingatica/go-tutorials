package main

import "fmt"

type Task struct {
	name, description string
	done              bool
}

type TaskList struct {
	tasks []Task
}

func (t *TaskList) addToList(tl Task) {
	t.tasks = append(t.tasks, tl)
}

func (t *TaskList) removeFromList(i int) {
	t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
}

func (t Task) printTaskFormat() {
	fmt.Println("Name:", t.name)
	fmt.Println("Description:", t.description)
	fmt.Println("Done:", t.done)
	fmt.Printf("\n")
}

func (t *TaskList) printTastks() {
	for _, v := range t.tasks {
		v.printTaskFormat()
	}
}

func (t *TaskList) printDoneTastks() {
	for _, v := range t.tasks {
		if !v.done {
			continue
		}

		v.printTaskFormat()
	}
}

func (t *Task) markDone() {
	t.done = true
}

func (t *Task) updateDescription(description string) {
	t.description = description
}

func (t *Task) updateName(name string) {
	t.name = name
}

func main() {
	t1 := Task{
		name:        "complete my go course",
		description: "complete my platzi go course this week",
	}

	t2 := Task{
		name:        "complete my python course",
		description: "complete my platzi python course this week",
	}

	tList1 := TaskList{
		tasks: []Task{t1, t2},
	}

	fmt.Printf("%+v\n", tList1)

	t3 := Task{
		name:        "complete my Node JS course",
		description: "complete my platzi Node JS course this week",
		done:        true,
	}

	tList1.addToList(t3)

	fmt.Printf("%+v\n", tList1)

	tList1.removeFromList(0)
	fmt.Printf("%+v\n", tList1)

	t4 := Task{
		name:        "complete my Docker course",
		description: "complete my platzi Docker course this week",
	}

	tList1.addToList(t4)

	for i := 0; i < len(tList1.tasks); i++ {
		fmt.Printf("Index: %v\nname: %v\n\n", i, tList1.tasks[i].name)
	}

	for i, v := range tList1.tasks {
		fmt.Printf("Index: %v\nname: %v\n\n", i, v.name)
	}

	fmt.Println("Tasks:")
	tList1.printTastks()

	tList1.tasks[2].markDone()

	fmt.Println("Done tasks:")
	tList1.printDoneTastks()

	t5 := Task{
		name:        "complete my Java course",
		description: "complete my platzi Java course this week",
	}

	t6 := Task{
		name:        "complete my C# course",
		description: "complete my platzi C# course this week",
	}

	tList2 := TaskList{
		tasks: []Task{t5, t6},
	}

	tasksMap := make(map[string]*TaskList)

	tasksMap["benjamin"] = &tList1
	tasksMap["amparo"] = &tList2

	fmt.Println("benjamin's tasks:")
	tasksMap["benjamin"].printTastks()

	fmt.Println("amparo's tasks:")
	tasksMap["amparo"].printTastks()
}
