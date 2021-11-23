import { Component } from '@angular/core';
import { BaseModal, ModalService } from 'carbon-components-angular';
import { Observable, Subject } from 'rxjs';

@Component({
  selector: 'app-event-booking-modal',
  templateUrl: './event-booking-modal.component.html',
  styleUrls: ['./event-booking-modal.component.scss'],
})
export class EventBookingModalComponent extends BaseModal {
  constructor(protected modalService: ModalService) {
    super();
  }
}
