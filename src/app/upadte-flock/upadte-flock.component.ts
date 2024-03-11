// import { Component, OnInit } from '@angular/core';
// import { AppService } from 'src/app/app.service';
// import { ActivatedRoute } from '@angular/router';




// @Component({
//   selector: 'app-upadte-flock',
//   templateUrl: 'upadte-flock.component.html',
//   styleUrls: ['upadte-flock.component.css']
// })
// export class UpadteFlockComponent implements OnInit {
//   responseData: any[] = []; // Assuming this is where you store the fetched flock data
//   flockData: any = {}; // Object to hold flock data for form
  
//   constructor(private appService: AppService,public route:ActivatedRoute) {}

//   // ngOnInit(): void {
//   //   this.fetchFlockData();
//   // }

//   ngOnInit() {
//     this.route.params.subscribe(params => {
//       const  id = params['id']; // Access the 'name' route parameter
//       console.log(id); // Here you have the received name
// //   const name ="c983990258"
//       this.fetchFlockData(id);
//    });
//   }
//   fetchFlockData(name :string): void {
//     let obj = name
//     this.appService.postRequest("listbyflock",obj).subscribe((result: any) => {
//       this.responseData = result;
//     });
//   }

//   onSubmit(): void {
//     // Implement logic to submit form data
//     // For example, you can send a POST request to update or create a new flock
//   }

//   editFlock(item: any): void {
//     // Populate the form fields with the selected flock's data
//     this.flockData = {
//       flockName: item.flockName,
//       breedName: item.breedName,
//       startDate: item.startDate,
//       startAge: item.startAge,
//       openingBirds: item.openingBirds,
//       shedNumber: item.shedNumber,
//       // Add more properties if needed
//     };
//   }

//   toggleFileDetails(event: Event): void {
//     const container = event.currentTarget as HTMLElement;
//     container.classList.toggle('expanded');
//   }
// }


import { Component, OnInit } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';



@Component({
  selector: 'app-upadte-flock',
  templateUrl: 'upadte-flock.component.html',
  styleUrls: ['upadte-flock.component.css']
})
export class UpadteFlockComponent implements OnInit {
  responseData: any[] = []; // Assuming this is where you store the fetched flock data
  flockData: any = {}; // Object to hold flock data for form
  id : string ="";
  constructor(private appService: AppService, private route: ActivatedRoute, private router: Router, private toastr: ToastrService) {}

  ngOnInit() {
    this.route.params.subscribe(params => {
       this.id = params['id']; 
      console.log("fetch call=",this.id)// Access the 'id' route parameter
      this.fetchFlockData(this.id);
    });
  }

  fetchFlockData(id: string): void {
    this.appService.postRequest("listbyflock", { id: id }).subscribe((result: any) => {
      console.log("result", result);
      this.responseData = result;
      // Assign fetched data to flockData object for binding to form fields
      this.flockData = {
        active:result.active,
        id:result.id,
        flockName: result.flockName,
        breedName: result.breedName,
        startDate: result.startDate,
        startAge: result.startAge,
        openingBirds: result.openingBirds,
        shedNumber: result.shedNumber,
      };
    });
  }



 

  onSubmit() {
    
    let obj = {
      "id":this.id,
      "flockName": this.flockData.flockName,
      "breedName": this.flockData.breedName,
      "startDate": this.flockData.startDate,
      "Active":this.flockData.active,
      "startAge": parseInt(this.flockData.startAge),
      "openingBirds": parseInt(this.flockData.openingBirds),
      "shedNumber": this.flockData.shedNumber.toString(),
    };

    console.log(obj);
    
    this.appService.postRequest("updateflock", obj).subscribe(
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

  toggleFileDetails(event: Event): void {
    const container = event.currentTarget as HTMLElement;
    container.classList.toggle('expanded');
  }
}
