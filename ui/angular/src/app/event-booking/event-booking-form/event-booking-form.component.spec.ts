import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EventBookingFormComponent } from './event-booking-form.component';

describe('EventBookingFormComponent', () => {
  let component: EventBookingFormComponent;
  let fixture: ComponentFixture<EventBookingFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EventBookingFormComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EventBookingFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
