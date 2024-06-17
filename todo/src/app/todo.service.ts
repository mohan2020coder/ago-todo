// todo.service.ts

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TodoService {

  private apiUrl = 'http://localhost:8080/todos'; // Replace with your Golang API URL

  constructor(private http: HttpClient) { }

  getTodos(): Observable<any[]> {
    return this.http.get<any[]>(this.apiUrl);
  }

  addTodo(todo: any): Observable<any> {
    return this.http.post<any>(this.apiUrl, todo);
  }

  deleteTodo(id: number): Observable<any> {
    const url = `${this.apiUrl}/${id}`;
    return this.http.delete<any>(url);
  }
}

