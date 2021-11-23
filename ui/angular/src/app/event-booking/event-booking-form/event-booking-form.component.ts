import { Component, Inject, OnInit } from '@angular/core';
import { NumberChange } from 'carbon-components-angular';

@Component({
  selector: 'app-event-booking-form',
  templateUrl: './event-booking-form.component.html',
  styleUrls: ['./event-booking-form.component.scss'],
})
export class EventBookingFormComponent implements OnInit {
  constructor(
    @Inject('data') public data: any,
    @Inject('inputValue') public inputValue: number
  ) {}

  onChange(event: NumberChange): void {
    this.data.next(event.value);
  }

  ngOnInit(): void {}
}
