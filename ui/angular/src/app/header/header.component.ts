import { Component, HostBinding, Input, OnInit } from '@angular/core';
import { NavigationItem } from 'carbon-components-angular';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  @HostBinding('class.bx--header') headerClass=true;

  hasHamburger = false;

  headerItems: NavigationItem[] = [
    {
      type: "item",
      content: "Home",
      title: "Home",
      route: ["/"],
    },
    {
      type: "item",
      content: "Events",
      title: "Events",
      route: ["/events"],
    },
  ]

  constructor() { }

  ngOnInit(): void {
  }

}
