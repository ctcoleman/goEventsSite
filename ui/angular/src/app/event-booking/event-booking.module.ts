import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  ButtonModule,
  ModalModule,
  ModalService,
  NumberModule,
  PlaceholderModule,
  SelectModule,
} from 'carbon-components-angular';
import { FormsModule } from '@angular/forms';
import { EventBookingModalComponent } from './event-booking-modal/event-booking-modal.component';
import { EventBookingFormComponent } from './event-booking-form/event-booking-form.component';
import { EventBookingModalContainerComponent } from './event-booking-modal-container/event-booking-modal-container.component';

@NgModule({
  declarations: [
    EventBookingModalComponent,
    EventBookingFormComponent,
    EventBookingModalContainerComponent,
  ],
  imports: [
    CommonModule,
    SelectModule,
    NumberModule,
    FormsModule,
    ButtonModule,
    ModalModule,
    PlaceholderModule,
  ],
  exports: [ EventBookingModalContainerComponent ],
  providers: [ModalService],
})
export class EventBookingModule {}
