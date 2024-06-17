// import { Component } from '@angular/core';

// @Component({
//   selector: 'app-todo-list',
//   standalone: true,
//   imports: [],
//   templateUrl: './todo-list.component.html',
//   styleUrl: './todo-list.component.scss'
// })
// export class TodoListComponent {

// }
// todo-list.component.ts

import { Component, OnInit } from '@angular/core';
import { TodoService } from '../todo.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-todo-list',
  standalone: true,
  imports: [CommonModule], // Import CommonModule here
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss']
})
export class TodoListComponent implements OnInit {

  todos: any[] = [];


  constructor(private todoService: TodoService) { }

  ngOnInit(): void {
    this.loadTodos();
  }

  loadTodos(): void {
    this.todoService.getTodos().subscribe(
      todos => this.todos = todos,
      error => console.log('Error fetching todos', error)
    );
  }

  deleteTodo(id: number): void {
    this.todoService.deleteTodo(id).subscribe(
      () => {
        this.todos = this.todos.filter(todo => todo.id !== id);
      },
      error => console.log('Error deleting todo', error)
    );
  }

}
