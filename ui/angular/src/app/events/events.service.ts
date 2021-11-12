import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Event } from '../../models/event';

@Injectable({
  providedIn: 'root',
})

export class EventsService {
  eventsURL = environment.EVENTSURL;

  getEvents() {
    return this.http.get<Event[]>(this.eventsURL + 'events');
  }

  constructor(private http: HttpClient) {}
}
