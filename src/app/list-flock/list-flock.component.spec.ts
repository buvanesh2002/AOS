import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListFlockComponent } from './list-flock.component';

describe('ListFlockComponent', () => {
  let component: ListFlockComponent;
  let fixture: ComponentFixture<ListFlockComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ListFlockComponent]
    });
    fixture = TestBed.createComponent(ListFlockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
