import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DailyentryComponent } from './dailyentry.component';

describe('DailyentryComponent', () => {
  let component: DailyentryComponent;
  let fixture: ComponentFixture<DailyentryComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DailyentryComponent]
    });
    fixture = TestBed.createComponent(DailyentryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
