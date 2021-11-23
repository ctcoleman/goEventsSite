import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EventBookingPageComponent } from './event-booking/event-booking-page/event-booking-page.component';

const routes: Routes = [
  {
    path: '',
    loadChildren: () => import('./home/home.module').then((m) => m.HomeModule),
  },
  {
    path: 'events',
    loadChildren: () => import('./events/events.module').then((m) => m.EventsModule),
  },
  {
    path: 'booking',
    component: EventBookingPageComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
