import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EventsListContainerComponent } from './events-list-container/events-list-container.component';
import { EventsListComponent } from './events-list/events-list.component';
import { HttpClientModule } from '@angular/common/http';
import {
  GridModule,
  LoadingModule,
  TableModule,
} from 'carbon-components-angular';
import { EventsRoutingModule } from './events-routing.module';
import { EventBookingModule } from '../event-booking/event-booking.module';

@NgModule({
  declarations: [EventsListContainerComponent, EventsListComponent],
  imports: [
    CommonModule,
    GridModule,
    TableModule,
    LoadingModule,
    HttpClientModule,
    EventsRoutingModule,
    EventBookingModule
  ],
  exports: [EventsListContainerComponent],
})
export class EventsModule {}
