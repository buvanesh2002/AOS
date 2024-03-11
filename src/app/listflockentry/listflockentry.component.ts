import { Component, OnInit } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { Router } from '@angular/router';

interface Item {
  id: string;
  flockName: string;
  breedName: string;
  startDate: string;
  age: string;
  openingBirds: string;
  shedNumber: string;
  active: string;
  listentry: {
    entrydate: string;
    age: number; // Assuming it should be a number
    openingbirds: number;
    mortality: number;
    birdssold: number;
    closingbirds: number;
    cummortality: number;
    mortalityPercent: number; // Assuming it should be a percentage
    eggsperDay: number;
    eggproducion: number;
    productionPercent: number; // Assuming it should be a percentage
    feed: number;
    feedperBird: number;
    feedperEgg: number;
    cumFPE: number;
  }[]; // Array of entry objects
}

@Component({
  selector: 'app-listflockentry',
  templateUrl: './listflockentry.component.html',
  styleUrls: ['./listflockentry.component.css']
})

export class ListflockentryComponent implements OnInit {

  responseData: Item[] = [];

  constructor(private appservice: AppService, public router: Router) { }

  ngOnInit() {
    this.listData();
  }

  listData() {
    let obj = {};
    this.responseData = [];
    this.appservice.postRequest("listflock", obj).subscribe(
      (response) => {
        console.log('List data received:', response);
        if (response && response.length) {
          this.responseData.push(...response);
        }
      },
      (error) => {
        console.error('Failed to fetch list data:', error);
        // Handle error response here
      }
    );
  }

  getname(id: string) {
    console.log(id);
    this.router.navigate(['updateflock', id]);
  }
}
