import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-events-list-container',
  templateUrl: './events-list-container.component.html',
  styleUrls: ['./events-list-container.component.scss'],
})
export class EventsListContainerComponent implements OnInit {
  constructor() {}

  ngOnInit(): void {
    console.log('hello from Events')
  }
}
