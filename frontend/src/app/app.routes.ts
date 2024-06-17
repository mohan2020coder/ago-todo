import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { TodoComponent } from './todo/todo.component'; // Import your component(s) for routing

const routes: Routes = [
  { path: 'todos', component: TodoComponent },
  { path: '', redirectTo: '/todos', pathMatch: 'full' },
  // Add more routes as needed
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
