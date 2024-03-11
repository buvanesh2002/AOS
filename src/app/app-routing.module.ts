import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { DataComponent } from './data/data.component';
import { FlockEntryComponent } from './flock-entry/flock-entry.component';
import { ListFlockComponent } from './list-flock/list-flock.component';
import { UpadteFlockComponent } from './upadte-flock/upadte-flock.component';
import { FlockentrydataComponent } from './flockentrydata/flockentrydata.component';
import {DailyentryComponent} from './dailyentry/dailyentry.component';
import { ListflockentryComponent } from './listflockentry/listflockentry.component';


const routes: Routes = [
  { path: '', redirectTo: '/addflock', pathMatch: 'full' },
 //{ path: '',  component: FlockEntryComponent },
  { path: 'login', component: LoginComponent },
  {path: 'data', component: DataComponent},
  { path: 'addflock', component: FlockEntryComponent},
  {path  : 'listflock', component:ListFlockComponent},
  {path  : 'updateflock/:id', component:UpadteFlockComponent},
  {path : 'flockentrydata/:id', component: FlockentrydataComponent},
  {path:'dailyentry', component:DailyentryComponent },
  {path :'listflockentry' , component:ListflockentryComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
