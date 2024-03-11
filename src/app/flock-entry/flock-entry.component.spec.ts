import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FlockEntryComponent } from './flock-entry.component';

describe('FlockEntryComponent', () => {
  let component: FlockEntryComponent;
  let fixture: ComponentFixture<FlockEntryComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [FlockEntryComponent]
    });
    fixture = TestBed.createComponent(FlockEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
