import { Component } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { AppService } from 'src/app/app.service';

@Component({
  selector: 'app-flock-entry',
  templateUrl: './flock-entry.component.html',
  styleUrls: ['./flock-entry.component.css']
})
export class FlockEntryComponent {
  flockData: any = {};
  
 
  constructor(private router: Router, private toastr: ToastrService, private appservice: AppService) {}

  onSubmit() {
    let obj = {
      "flockName": this.flockData.flockName,
      "breedName": this.flockData.breedName,
      "startDate": this.flockData.startDate,
      "startAge": parseInt(this.flockData.startAge),
      "openingBirds": parseInt(this.flockData.openingBirds),
      "shedNumber": this.flockData.shedNumber.toString(),
    };

    console.log(obj);
    
    this.appservice.postRequest("addflock", obj).subscribe(
      (result: any) => {
        this.showSuccess('Flock created successfully!');
        this.resetForm();
        this.router.navigate(['listflock']);
      },
      (error: any) => {
        if (error.error && error.error.errors) {
         
          for (const field in error.error.errors) {
            this.showError(`Error in ${field}: ${error.error.errors[field]}`);
          }
        } else {
          
          this.showError('Error occurred while creating flock!');
        }
      }
    );
  }

  resetForm() {
    this.flockData = {};
  }

  showSuccess(msg: any) {
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
