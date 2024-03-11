import { Component } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { AppService } from 'src/app/app.service';

@Component({
  selector: 'app-dailyentry',
  templateUrl: './dailyentry.component.html',
  styleUrls: ['./dailyentry.component.css']
})
export class DailyentryComponent {

  flockData: any = {
    
  };
  constructor(private router: Router, private toastr: ToastrService, private appservice: AppService) {}

  //Mortality
//Eggs
//Feed
//BirdsSold
//CountErr
//Trays

  onSubmit() {
    let obj = {
      "id": this.flockData.id,
      "date": this.flockData.date.toString(),
      "mortality": parseInt(this.flockData.mortality) ,
      "extraEggs":  parseInt(this.flockData.extraEggs),
      "feed": parseInt(this.flockData.feed) ,
      "birdsSold": parseInt(this.flockData.birdsSold) ,
      "countErr":  parseInt(this.flockData.countErr) ,
      "remarks": this.flockData.remarks,
      "trays":parseInt(this.flockData.trays) 
    };

    console.log(obj);
    
    this.appservice.postRequest("dailyentries", this.flockData).subscribe(
      (result: any) => {
        this.showSuccess('dailyentries created successfully!');
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
