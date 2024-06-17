import { Component } from '@angular/core';
import { TodoFormComponent } from './todo-form/todo-form.component';
import { TodoListComponent } from './todo-list/todo-list.component';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http'; // Import this
import { TodoService } from './todo.service';

@Component({
  selector: 'app-root',
  standalone: true,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'], // Corrected this
  imports: [
    TodoFormComponent,
    TodoListComponent,
    ReactiveFormsModule,
    HttpClientModule // Add this
  ],
  providers: [TodoService],
})
export class AppComponent {
  title = 'todo';
}
