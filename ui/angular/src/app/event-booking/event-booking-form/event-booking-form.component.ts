import { Component, Inject, Output, EventEmitter } from '@angular/core';
import { NumberChange } from 'carbon-components-angular';

@Component({
  selector: 'app-event-booking-form',
  templateUrl: './event-booking-form.component.html',
  styleUrls: ['./event-booking-form.component.scss'],
})
export class EventBookingFormComponent {
  constructor(
    @Inject('data') public data: any,
    @Inject('seats') public seats: number,
    @Inject('eventName') public eventName: string,
    @Inject('eventID') public eventID: string,
  ) {}

  onChange(event: NumberChange): void {
    this.data.next(event.value);
  }
}
