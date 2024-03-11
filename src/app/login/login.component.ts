import { Component } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { AppService } from 'src/app/app.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  
  constructor(private router: Router,private toastr: ToastrService,private appservice: AppService){}

  forms: any[] = [];
  username: string = ''; // Added username variable
  password: string = ''; // Added password variable

  // fetch() {
  //   let obj = {};
  //   this.appservice.postRequest("list", obj).subscribe((result: any) => {
  //     this.forms = result;
  //   });
  // }

  login() {
    console.log("login");
    let obj = {
        "email": this.username,
        "password": this.password
    };

    console.log(obj);

    this.appservice.postRequest("login", obj).subscribe(
        (result: any) => {
            console.log(result);
            if (result && result['message']) {  
            //     this.showSucess(result['message']);
            //    this.router.navigate(['listflock']);
            //  this.router.navigate(['updateflock']);
            // this.router.navigate(['addflock']);
            //  this.router.navigate(['dailyentry']);
             this.router.navigate(['listflockentry']);
            } else {
                console.log("else")
                this.showError(result['response']);
            }
        },
        (error: any) => {
            console.error("HTTP error:", error);
            this.showError("Ivalid Credential");
        }
    );
}

Agecalc(){
  
}

  showSucess( msg : any)  {
    if (msg) {
      this.toastr.success(msg);
    }
  }
  showError(msg: any) {
    if (msg) {
      this.toastr.error(msg);
    }
  }
}
