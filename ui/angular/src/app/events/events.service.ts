import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Event } from '../../models/event';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class EventsService {
  private eventsURL: string = environment.EVENTSURL;

  getEvents(): Observable<Event[]> {
    return this.http.get<Event[]>(this.eventsURL + 'events');
  }

  constructor(private http: HttpClient) {}
}
