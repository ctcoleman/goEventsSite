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
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { EventBookingModalComponent } from './event-booking-modal/event-booking-modal.component';
import { EventBookingPageComponent } from './event-booking-page/event-booking-page.component';
import { EventBookingModalBtnComponent } from './event-booking-modal-btn/event-booking-modal-btn.component';
import { EventBookingFormComponent } from './event-booking-form/event-booking-form.component';
import { EventBookingModalContainerComponent } from './event-booking-modal-container/event-booking-modal-container.component';

@NgModule({
  declarations: [
    EventBookingModalComponent,
    EventBookingPageComponent,
    EventBookingModalBtnComponent,
    EventBookingFormComponent,
    EventBookingModalContainerComponent,
  ],
  imports: [
    CommonModule,
    SelectModule,
    NumberModule,
    BrowserModule,
    FormsModule,
    ButtonModule,
    ModalModule,
    PlaceholderModule,
  ],
  providers: [ModalService],
})
export class EventBookingModule {}
