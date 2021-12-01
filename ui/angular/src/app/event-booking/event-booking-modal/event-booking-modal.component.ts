import { Component, Inject } from '@angular/core';
import { BaseModal } from 'carbon-components-angular';

@Component({
  selector: 'app-event-booking-modal',
  templateUrl: './event-booking-modal.component.html',
  styleUrls: ['./event-booking-modal.component.scss'],
})
export class EventBookingModalComponent extends BaseModal {
  constructor(@Inject('eventID') public eventID: string) {
    super();
  }
  
  bookEvent(): void {
  }
}
