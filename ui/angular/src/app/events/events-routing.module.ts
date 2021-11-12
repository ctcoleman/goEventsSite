import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EventsListContainerComponent } from './events-list-container/events-list-container.component';

const routes: Routes = [
    {
        path: '',
        component: EventsListContainerComponent,
    },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class EventsRoutingModule {}
