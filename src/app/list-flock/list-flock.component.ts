import { Component, OnInit } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { Router } from '@angular/router';


interface Item {
  id: string;
  flockName:string;
  breedName: string;
  startDate: string;
  age: string;
  openingBirds: string;
  shedNumber: string;
  active: string;
}

@Component({
  selector: 'app-list-flock',
  templateUrl: './list-flock.component.html',
  styleUrls: ['./list-flock.component.css']
})

export class ListFlockComponent implements OnInit {
  
 responseData: Item[] = [];


   constructor(private appservice: AppService,public router:Router) {}
 
  ngOnInit() {
    this.listData()
  }


  listData() {
    let obj = {};
    this.responseData= []
    this.appservice.postRequest("listflock", obj).subscribe(
      (response) => {
        console.log('List data received:', response);
        if (response && response.length) {
          console.log(response)
          console.log(response.id);
          this.responseData.push(...response); 
          
        }
      },
      (error) => {
        console.error('Failed to fetch list data:', error);
        // Handle error response here
      }
    );
  }
  getname( id:string){
    console.log(id);
    this.router.navigate(['updateflock',id])
}
}
