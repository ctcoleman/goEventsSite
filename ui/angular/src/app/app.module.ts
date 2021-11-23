import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { UIShellModule } from 'carbon-components-angular';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { EventBookingModule } from './event-booking/event-booking.module';
import { EventsModule } from './events/events.module';
import { HeaderComponent } from './header/header.component';

@NgModule({
  declarations: [AppComponent, HeaderComponent],
  imports: [BrowserModule, AppRoutingModule, EventsModule, UIShellModule, EventBookingModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
