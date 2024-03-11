import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FlockentrydataComponent } from './flockentrydata.component';

describe('FlockentrydataComponent', () => {
  let component: FlockentrydataComponent;
  let fixture: ComponentFixture<FlockentrydataComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [FlockentrydataComponent]
    });
    fixture = TestBed.createComponent(FlockentrydataComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
