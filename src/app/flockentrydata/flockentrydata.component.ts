import { Component, OnInit } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';

interface FlockData {
  mortality?: number;
  extraeggs?: number;
  feedbags?: number;
  extrafeed?: number;
  birdssold?: number;
  errors?: number;
  remarks?: string;
}

@Component({
  selector: 'app-flockentrydata',
  templateUrl: './flockentrydata.component.html',
  styleUrls: ['./flockentrydata.component.css']
})
export class FlockentrydataComponent implements OnInit {
  flockData: FlockData = {};
  id: string = "";

  constructor(private appService: AppService, private route: ActivatedRoute, private router: Router) {}

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.id = params['id']; 
      console.log("fetch call=", this.id); // Access the 'id' route parameter
      this.fetchFlockData(this.id);
    });
  }

  fetchFlockData(id: string): void {
    this.appService.postRequest("listbyflock", { id: id }).subscribe(
      (result: any) => {
        this.flockData = result;
      },
      (error: any) => {
        console.error("Error fetching flock data:", error);
        // Handle error, e.g., show error message
      }
    );
  }

  onSubmit(): void {
    // Handle form submission
    console.log("Form submitted:", this.flockData);
  }
}
