import { AfterContentInit, Component, Input } from '@angular/core';
import { ModalService } from 'carbon-components-angular';
import { Observable, Subject } from 'rxjs';
import { Event } from '../../../models/event';
import { BookingService } from '../booking.service';
import { EventBookingModalComponent } from '../event-booking-modal/event-booking-modal.component';

@Component({
  selector: 'app-event-booking-modal-container',
  templateUrl: './event-booking-modal-container.component.html',
  styleUrls: ['./event-booking-modal-container.component.scss'],
})
export class EventBookingModalContainerComponent {
  @Input() eventID: string = '';
  @Input() seats: number = 0;

  public data: Observable<number> = new Subject<number>();
  public eventName: string = 'Event Booking';
  public state: string = 'loading';

  constructor(
    protected modalService: ModalService,
    protected bookingService: BookingService
  ) {}

  openModal(): void {
    this.modalService.create({
      component: EventBookingModalComponent,
      inputs: {
        seats: this.seats,
        data: this.data,
        eventName: this.eventName,
        eventID: this.eventID,
        state: this.state,
      },
    });
  }
}
