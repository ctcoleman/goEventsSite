import { ComponentFixture, TestBed } from '@angular/core/testing';
import { TableModule } from 'carbon-components-angular';
import { Event } from '../../../models/event';

import { EventsListComponent } from './events-list.component';

describe('EventsListComponent', () => {
  let component: EventsListComponent;
  let fixture: ComponentFixture<EventsListComponent>;
  let eventsServiceStub: Partial<EventsService>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EventsListComponent ],
      imports: [ TableModule ],
      providers: [ { provide: EventsService, useValue: eventsServiceStub } ],
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EventsListComponent);
    component = fixture.componentInstance;
    eventsServiceStub = {
      getEvents() {
        return {
          "ID": "616d658f3f7d8346e19f1fbd",
          "Name": "Chris's awesome party",
          "Duration": 0,
          "StartDate": 0,
          "EndDate": 0,
          "Location": {
            "ID": "616d658f3f7d8346e19f1fbe",
            "Name": "The Pit",
            "City": "Buffalo",
            "State": "NY",
            "Country": "USA",
            "OpenTime": 0,
            "CloseTime": 0,
            "Venues": []
          }
        }
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
