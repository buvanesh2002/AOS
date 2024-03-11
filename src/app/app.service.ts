import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  })
};

@Injectable({
  providedIn: 'root'
})


export class AppService {
  getDevice(emailId: any): any {
    throw new Error('Method not implemented.');
  }
  setToLocalStorage(arg0: string, emailId: any) {
    throw new Error('Method not implemented.');
  }
  constructor(private http: HttpClient, private route: Router) { }

  postRequest(method: any, obj: any) {
    return this.http.post<any>(method, obj, httpOptions);
  }
}
