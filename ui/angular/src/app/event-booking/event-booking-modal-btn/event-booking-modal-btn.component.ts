import { Component, OnInit } from '@angular/core';
import { ModalService } from 'carbon-components-angular';
import { Observable, Subject } from 'rxjs';
import { EventBookingModalComponent } from '../event-booking-modal/event-booking-modal.component';

@Component({
  selector: 'app-event-booking-modal-btn',
  templateUrl: './event-booking-modal-btn.component.html',
  styleUrls: ['./event-booking-modal-btn.component.scss'],
})
export class EventBookingModalBtnComponent implements OnInit {
  public inputValue = 0;
  public data: Observable<number> = new Subject<number>();

  constructor(protected modalService: ModalService) {}

  openModal(): void {
    this.modalService.create({
      component: EventBookingModalComponent,
      inputs: {
        inputValue: this.inputValue,
        data: this.data,
      },
    });
  }

  ngOnInit(): void {
    this.data.subscribe((value) => (this.inputValue = value));
  }
}
