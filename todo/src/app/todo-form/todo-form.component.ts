// import { Component } from '@angular/core';

// @Component({
//   selector: 'app-todo-form',
//   standalone: true,
//   imports: [],
//   templateUrl: './todo-form.component.html',
//   styleUrl: './todo-form.component.scss'
// })
// export class TodoFormComponent {

// }
// todo-form.component.ts

import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { TodoService } from '../todo.service';
import {ReactiveFormsModule} from '@angular/forms';

@Component({
  selector: 'app-todo-form',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './todo-form.component.html',
  styleUrls: ['./todo-form.component.scss']
})
export class TodoFormComponent implements OnInit {

  todoForm: FormGroup;

  constructor(private fb: FormBuilder, private todoService: TodoService) {
    this.todoForm = this.fb.group({
      title: ['']
    });
  }

  ngOnInit(): void {
  }

  onSubmit(): void {
    const todoData = this.todoForm.value;
    this.todoService.addTodo(todoData).subscribe(
      () => {
        console.log('Todo added successfully');
        // Optionally, reset form or update list
        this.todoForm.reset();
      },
      error => console.log('Error adding todo', error)
    );
  }

}
