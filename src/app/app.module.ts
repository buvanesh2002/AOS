import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { FormsModule } from '@angular/forms';
import { ToastrModule } from 'ngx-toastr';
import { HttpClientModule } from '@angular/common/http';
import{ BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DataComponent } from './data/data.component';
import { FlockEntryComponent } from './flock-entry/flock-entry.component';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ListFlockComponent } from './list-flock/list-flock.component';
import { MatTableModule } from '@angular/material/table';
import { UpadteFlockComponent } from './upadte-flock/upadte-flock.component';
import { FlockentrydataComponent } from './flockentrydata/flockentrydata.component';
import { ReminderComponent } from './reminder/reminder.component';
import { DailyentryComponent } from './dailyentry/dailyentry.component';
import { ListflockentryComponent } from './listflockentry/listflockentry.component';
// import { AppService } from './app.service'; 

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DataComponent,
    FlockEntryComponent,
    ListFlockComponent,
    UpadteFlockComponent,
    FlockentrydataComponent,
    ReminderComponent,
    DailyentryComponent,
    ListflockentryComponent
    
    
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot(),
    MatInputModule,
    MatFormFieldModule,
    MatTableModule
  ],
  providers: [AppComponent],
  bootstrap: [AppComponent]
})
export class AppModule { }
