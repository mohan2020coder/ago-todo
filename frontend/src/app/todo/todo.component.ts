import { Component, OnInit } from '@angular/core';
import { TodoService } from '../todo.service';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.scss']
})
export class TodoComponent implements OnInit {
  todos: any[] = [];
  newTodo: string = '';

  constructor(private todoService: TodoService) { }

  ngOnInit(): void {
    this.getTodos();
  }

  getTodos(): void {
    this.todoService.getTodos()
      .subscribe(todos => this.todos = todos);
  }

  addTodo(): void {
    if (this.newTodo.trim()) {
      this.todoService.addTodo({ title: this.newTodo, completed: false })
        .subscribe(todo => {
          this.todos.push(todo);
          this.newTodo = '';
        });
    }
  }

  updateTodo(todo: any): void {
    this.todoService.updateTodo(todo.id, todo)
      .subscribe(updatedTodo => {
        const index = this.todos.findIndex(t => t.id === updatedTodo.id);
        if (index !== -1) {
          this.todos[index] = updatedTodo;
        }
      });
  }

  deleteTodo(id: number): void {
    this.todoService.deleteTodo(id)
      .subscribe(() => {
        this.todos = this.todos.filter(todo => todo.id !== id);
      });
  }
}
