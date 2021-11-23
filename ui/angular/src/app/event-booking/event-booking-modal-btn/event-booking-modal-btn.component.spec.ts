import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EventBookingModalBtnComponent } from './event-booking-modal-btn.component';

describe('EventBookingModalBtnComponent', () => {
  let component: EventBookingModalBtnComponent;
  let fixture: ComponentFixture<EventBookingModalBtnComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EventBookingModalBtnComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EventBookingModalBtnComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
