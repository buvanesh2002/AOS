import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpadteFlockComponent } from './upadte-flock.component';

describe('UpadteFlockComponent', () => {
  let component: UpadteFlockComponent;
  let fixture: ComponentFixture<UpadteFlockComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [UpadteFlockComponent]
    });
    fixture = TestBed.createComponent(UpadteFlockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
