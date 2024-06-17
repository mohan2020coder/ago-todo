import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app.routes'; // Import your routing module
import { AppComponent } from './app.component';
import { AppConfig } from './app.config'; // Import your app config

@NgModule({
  declarations: [
    AppComponent
    // Add other components and directives
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule // Include your routing module
    // Other modules imports as needed
  ],
  providers: [
    { provide: 'APP_CONFIG', useValue: AppConfig } // Provide your AppConfig
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
