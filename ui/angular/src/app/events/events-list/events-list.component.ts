import { Component, ElementRef, OnInit, TemplateRef, ViewChild } from '@angular/core';
import {
  Table,
  TableItem,
  TableHeaderItem,
  TableModel,
} from 'carbon-components-angular';
import { EventBookingModalContainerComponent } from 'src/app/event-booking/event-booking-modal-container/event-booking-modal-container.component';
import { Event } from '../../../models/event';
import { EventsService } from '../events.service';

@Component({
  selector: 'app-events-list',
  templateUrl: './events-list.component.html',
  styleUrls: ['./events-list.component.scss'],
})
export class EventsListComponent implements OnInit {
  @ViewChild('customItemTemplate')
  protected customItemTemplate?: TemplateRef<any>;

  events: Event[] = [];
  model: TableModel = {} as TableModel;
  skeleton: boolean = true;
  skeletonModel = Table.skeletonModel(this.events.length, 5);

  constructor(private eventsService: EventsService) {}

  ngOnInit(): void {
    this.model = new TableModel();

    this.model.header = [
      new TableHeaderItem({ data: 'Event' }),
      new TableHeaderItem({ data: 'Where' }),
      new TableHeaderItem({ data: 'Start Date' }),
      new TableHeaderItem({ data: 'End Date' }),
      new TableHeaderItem({ data: 'Actions' }),
    ];

    this.eventsService.getEvents().subscribe((resp: any) => {
      if (resp.error) {
        const errorData = [];
        errorData.push([new TableItem({ data: 'error!' })]);
        this.model.data = errorData;
      } else if (resp.loading) {
        this.skeleton = true;
      } else {
        this.events = resp;
        this.skeleton = false;
        this.model.pageLength = 10;
        this.model.totalDataLength = this.events.length;
        this.selectPage(1);
      }
    });
  }

  prepareEvents(events: Event[]) {
    this.skeleton = false;
    const eventsTableArray = [];

    for (const event of events) {
      eventsTableArray.push([
        new TableItem({ data: event.Name }),
        new TableItem({ data: event.Location.Name }),
        new TableItem({
          data: new Date(event.StartDate).toLocaleDateString(undefined, {
            year: '2-digit',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
          }),
        }),
        new TableItem({
          data: new Date(event.EndDate).toLocaleDateString(undefined, {
            year: '2-digit',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
          }),
        }),
        new TableItem({ data: { eventID: event.ID }, template: this.customItemTemplate }),
      ]);
    }
    return eventsTableArray;
  }

  selectPage(page: number) {
    const offset = this.model.pageLength * (page - 1);
    const pageRawData = this.events.slice(
      offset,
      offset + this.model.pageLength
    );
    this.model.data = this.prepareEvents(pageRawData);
    this.model.currentPage = page;
  }
}
