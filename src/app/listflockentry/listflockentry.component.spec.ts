import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListflockentryComponent } from './listflockentry.component';

describe('ListflockentryComponent', () => {
  let component: ListflockentryComponent;
  let fixture: ComponentFixture<ListflockentryComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ListflockentryComponent]
    });
    fixture = TestBed.createComponent(ListflockentryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
