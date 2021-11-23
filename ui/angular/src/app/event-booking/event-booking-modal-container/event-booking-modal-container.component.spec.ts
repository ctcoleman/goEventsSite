import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EventBookingModalContainerComponent } from './event-booking-modal-container.component';

describe('EventBookingModalContainerComponent', () => {
  let component: EventBookingModalContainerComponent;
  let fixture: ComponentFixture<EventBookingModalContainerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EventBookingModalContainerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EventBookingModalContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
