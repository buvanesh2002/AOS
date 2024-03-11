// 

import { Component } from '@angular/core';

import { HttpClient } from '@angular/common/http';


@Component({
  selector: 'app-data',
  templateUrl: './data.component.html',
  styleUrls: ['./data.component.css']
})
export class DataComponent {
  
  selectedFile: File | null = null;
  fileName: any;
  fileText: any;
  length: any;

  constructor(private http: HttpClient) { }

  onFileSelected(event: any) {
    this.selectedFile = event.target.files[0];
  }
  readURL(event: any): void {
    console.log("event ", event)
    var file = event.target.files[0]
    var reader = new FileReader();
    reader.readAsDataURL (file);
    this.fileName = file.name
    reader.onload = (e: any) => {
      this.fileText = e.target.result
      this.length = this.fileText.length
    }
  }

  onSubmit() {
   

   let obj={'file':this.fileText}

    this.uploadFileContent(obj);
  }

  uploadFileContent(formData: any) {
    this.http.post<any>('fileupload', formData).subscribe(
      (response) => {
        console.log('File upload successful:', response);
        // Handle success response here
      },
      (error) => {
        console.error('File upload failed:', error);
        // Handle error response here
      }
    );
  }

}

  // onSubmit(event: Event) {
  //   event.preventDefault(); 

  //   if (this.selectedFile) {
  //     this.uploadFile(this.selectedFile);
  //   } else {
  //     console.log('No file selected.');
  //   }
  // }

  // uploadFile(file: File) {
  //   const formData = new FormData();
  //   formData.append('excelfile', file);
    
  //   // Convert file content to string
  //   const jsonString = JSON.stringify(this.fileContent);
  //   formData.append('fileContent', jsonString);
  
  //   // Send FormData to backend
  //   this.appservice.postRequest("fileupload", formData).subscribe(
  //     (result: any) => {
  //         console.log(result);
  //         if (result && result['message']) {  
  //            // Handle success
  //         } else {
  //             console.log("else")
  //            // Handle failure
  //         }
  //     },
  //     (error: any) => {
  //         console.error("HTTP error:", error);
  //         // Handle error
  //     }
  //   );
  // }



// import { Component } from '@angular/core';
// import { HttpClient } from '@angular/common/http';

// @Component({
//   selector: 'app-data',
//   templateUrl: './data.component.html',
//   styleUrls: ['./data.component.css']
// })
// export class DataComponent {
  
//   selectedFile: File | null = null;
//   fileName: string = '';

//   constructor(private http: HttpClient) { }

//   onFileSelected(event: any) {
//     this.selectedFile = event?.target?.files?.[0] || null; // Use optional chaining
//     this.fileName = this.selectedFile ? this.selectedFile.name : '';
//   }

//   onSubmit() {
//     if (!this.selectedFile) {
//       console.error('No file selected');
//       return;
//     }

//     const formData: FormData = new FormData();
//     formData.append('file', this.selectedFile);

//     this.uploadFile(formData);
//   }

//   uploadFile(formData: FormData) {
//     this.http.post<any>('fileupload', formData).subscribe(
//       (response) => {
//         console.log('File upload successful:', response);
//         // Handle success response here
//       },
//       (error) => {
//         console.error('File upload failed:', error);
//         // Handle error response here
//       }
//     );
//   }
// }
