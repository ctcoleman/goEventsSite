import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Event } from 'src/models/event';

@Injectable({
  providedIn: 'root',
})
export class BookingService {
  private eventURL: string = environment.EVENTSURL;
  private bookingURL: string = environment.BOOKINGURL;

  getEvent(eventID: string): Observable<Event> {
    return this.http.get<Event>(this.eventURL + 'events' + eventID);
  }

  postBooking(eventID: string, payload: string): Observable<Event> {
    return this.http.post<Event>(
      this.bookingURL + 'events' + eventID + 'bookings',
      payload
    );
  }

  constructor(private http: HttpClient) {}
}
